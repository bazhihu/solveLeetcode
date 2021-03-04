package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

/*
	多线程下载的原理
	可以突破百度网盘下载限制
	https://mojotv.cn/go/go-range-download
	关键信息
	http - Accept-Ranges: bytes
			Content-Length: 13131


	流程：
	1、启动文件下载器 初始化信息
	2、开始下载任务
	3、获取文件head信息, 文件大小、文件名称、是否支持 多线程下载等
	4、计算分片信息
	5、开启多线程，开始下载
	6、下载完成、将分片信息🔂依次写入文件，并校验sha
*/

func main() {
	startTime := time.Now()
	var url string // 下载文件的地址

	url = "https://download.jetbrains.com/go/goland-2020.2.2.dmg"
	downloader := NewFileDownLoader(url, "", "", 10)
	if err := downloader.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n 文件下载完成耗时：%f second\n", time.Now().Sub(startTime).Seconds())
}

// FileDownLoader 文件下载器
type FileDownLoader struct {
	fileSize       int
	url            string
	outputFileName string
	totalPart      int // 下载线程
	outputDir      string
	doneFilePart   []filePart
}

// filePart 文件分片
type filePart struct {
	Index int    // 文件分片的序号
	From  int    // 开始byte
	To    int    // 结束byte
	Data  []byte // http
}

func NewFileDownLoader(url, outputFileName, outputDir string, totalPart int) *FileDownLoader {
	if outputDir == "" {
		wd, err := os.Getwd() // 获取当前工作目录
		if err != nil {
			log.Println(err)
		}
		outputDir = wd
	}
	return &FileDownLoader{
		fileSize:       0,
		url:            url,
		outputFileName: outputFileName,
		totalPart:      totalPart,
		outputDir:      outputDir,
		doneFilePart:   make([]filePart, totalPart),
	}
}

// Run 开始下载任务
func (d *FileDownLoader) Run() error {
	fileTotalSize, err := d.head()
	if err != nil {
		return err
	}
	d.fileSize = fileTotalSize
	jobs := make([]filePart, d.totalPart)
	eachSize := fileTotalSize / d.totalPart

	// 分片划分大小
	for i := range jobs {
		jobs[i].Index = i
		if i == 0 {
			jobs[i].From = 0
		} else {
			jobs[i].From = jobs[i-1].To + 1
		}
		if i < d.totalPart-1 {
			jobs[i].To = jobs[i].From + eachSize
		} else {
			// the last filePart
			jobs[i].To = fileTotalSize - 1
		}
	}

	// 启动多线程下载
	var wg sync.WaitGroup
	for _, j := range jobs {
		wg.Add(1)
		go func(job filePart) {
			defer wg.Done()
			err := d.downLoadPart(job)
			if err != nil {
				log.Println("下载文件失败：", err, job)
			}
		}(j)
	}
	wg.Wait()
	return d.mergeFileParts()
}

// 下载分片
func (d FileDownLoader) downLoadPart(c filePart) error {
	r, err := d.getNewRequest("GET")
	if err != nil {
		return err
	}
	log.Printf("开始[%d]下载from: %d to: %d\n", c.Index, c.From, c.To)
	r.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", c.From, c.To))
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	if resp.StatusCode > 299 {
		return errors.New(fmt.Sprintf("服务错误状态码： %v", resp.StatusCode))
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(bs) != (c.To - c.From + 1) {
		return errors.New("下载文件分片长度错误")
	}
	c.Data = bs
	d.doneFilePart[c.Index] = c
	return nil
}

// head 获取要下载的文件的基本信息（header）使用HTTP Method Head
func (d *FileDownLoader) head() (int, error) {
	r, err := d.getNewRequest("HEAD")
	if err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode > 299 {
		return 0, errors.New(fmt.Sprintf("Can't process, response is %v", resp.StatusCode))
	}

	// 检查是否支持， 断点续传
	//https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Ranges
	if resp.Header.Get("Accept-Ranges") != "bytes" {
		return 0, errors.New("服务器不支持文件断点续传")
	}
	d.outputFileName = parseFileInfoFrom(resp)
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Length
	return strconv.Atoi(resp.Header.Get("Content-Length"))
}

// 解析文件信息
func parseFileInfoFrom(resp *http.Response) string {
	contentDisposition := resp.Header.Get("Content-Disposition")
	if contentDisposition != "" {
		_, params, err := mime.ParseMediaType(contentDisposition)

		if err != nil {
			panic(err)
		}
		return params["filename"]
	}
	filename := filepath.Base(resp.Request.URL.Path)
	return filename
}

// getNewRequest 创建一个request
func (d FileDownLoader) getNewRequest(method string) (*http.Request, error) {
	r, err := http.NewRequest(method, d.url, nil)
	if err != nil {
		return nil, err
	}
	r.Header.Set("User-Agent", "herman")
	return r, nil
}

// mergeFileParts 合并下载的文件
func (d FileDownLoader) mergeFileParts() error {
	log.Println("开始合并文件")
	path := filepath.Join(d.outputDir, d.outputFileName)
	mergedFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer mergedFile.Close()

	hash := sha256.New()
	totalSize := 0
	for _, s := range d.doneFilePart {
		mergedFile.Write(s.Data)
		hash.Write(s.Data)
		totalSize += len(s.Data)
	}

	if totalSize != d.fileSize {
		return errors.New("文件不完整")
	}

	// 校检 md5
	if hex.EncodeToString(hash.Sum(nil)) != "3af4660ef22f805008e6773ac25f9edbc17c2014af18019b7374afbed63d4744" {
		return errors.New("文件损坏")
	} else {
		log.Println("文件SHA-256校验成功")
	}
	return nil
}
