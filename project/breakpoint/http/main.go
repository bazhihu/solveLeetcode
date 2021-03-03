package main

import (
	"fmt"
	"log"
	"time"
)

/*
	多线程下载的原理
	可以突破百度网盘下载限制

	关键信息
	http - Accept-Ranges: bytes
			Content-Length: 13131
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

func (d *FileDownloader) head() (int, error) {

}
