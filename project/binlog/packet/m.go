package packet

//import (
//	"bufio"
//	"bytes"
//	"crypto/sha1"
//	"fmt"
//	"io"
//	"net"
//	"os"
//	"os/signal"
//	"syscall"
//	"time"
//	"unsafe"
//)
//
//var (
//	QUERY = uint8(3)
//	SequenceId uint8 // 串包码
//	salt []uint8
//	password string = "123456"
//)
//
//func main() {
//	var errChan = make(chan error)
//	conn, err := net.DialTimeout("tcp", "127.0.0.1:3306", 10*time.Second)
//	if err != nil {
//		fmt.Println("err", err)
//		os.Exit(-1)
//	}
//
//	fmt.Println("conn-success", conn)
//
//	go read(conn)
//
//	send(conn, `SHOW GLOBAL VARIABLES LIKE 'BINLOG_CHECKSUM'`)
//
//	go func() {
//		c := make(chan os.Signal, 1)
//		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
//		errChan <- fmt.Errorf("%s", <-c)
//	}()
//	error := <-errChan
//
//	fmt.Println("err", error)
//}
//
//func getMysqlVersion(data []byte) {
//	pos := 1 + bytes.IndexByte(data[1:], 0x00) + 1
//	fmt.Println("pos", pos)
//}
//
//
//// read mysql data
//func read(conn net.Conn) {
//	go processBinlog(conn)
//}
//
//func processBinlog(conn net.Conn) {
//	defer conn.Close()
//
//	for {
//		reader := bufio.NewReader(conn)
//		//rec := new(*bytes.Buffer)
//		resHead := make([]byte, 4)
//		// 读取数据
//		_, err := io.ReadFull(reader, resHead)
//		if err != nil {
//			fmt.Println("read head failed, err:", err)
//			break
//		}
//
//		length := int(uint32(resHead[0]) | uint32(resHead[1])<<8 | uint32(resHead[2])<<16)
//		v := uint8(resHead[3])
//		if v != 0 {
//			fmt.Println("read version failed, err:", err)
//			break
//		}
//
//		//(*rec).Grow(length)
//
//		bodyBuf := make([]byte, length)
//
//		_, err = io.ReadAtLeast(reader, bodyBuf, length)
//		if err != nil {
//			fmt.Println("read body failed, err:", err)
//			break
//		}
//
//		//(*rec).Write(bodyBuf)
//
//		fmt.Println("收到client端发来的数据：", string(bodyBuf))
//
//		getMysqlVersion(bodyBuf)
//
//		handshake(conn)
//
//
//	}
//}
//
//func send(conn net.Conn, command string) {
//	arg := makeHead(getCommandBuf([]byte(command)))
//
//	if n, err := conn.Write(arg); err != nil || n != len(arg) {
//		fmt.Println("err", err)
//		return
//	}
//	return
//}
//
//func makeHead(arg []byte) []byte {
//	length := len(arg) - 4
//	arg[0] = byte(length)
//	arg[1] = byte(length >> 8)
//	arg[2] = byte(length >> 16)
//	arg[3] = uint8(0)
//	return arg
//}
//
//// header has 4 bytes
//func getCommandBuf(arg []byte) []byte {
//	length := len(arg) + 1 + 4
//	// new array byte
//	data := make([]byte, length)
//
//	data[4] = QUERY
//	copy(data[5:], arg)
//	return data
//}
//
//func StringToByteSlice(s string) []byte {
//	return *(*[]byte)(unsafe.Pointer(&s))
//}
//
//func handleAuthResult() {
//
//}
//
//func handshake(conn net.Conn) {
//	capability := uint32(565765)
//
//	auth := CalcPassword(salt, []byte(password))
//
//	data := make([]byte, 84)
//	data[4] = byte(capability)
//	data[5] = byte(capability >> 8)
//	data[6] = byte(capability >> 16)
//	data[7] = byte(capability >> 24)
//
//	data[8] = 0x00
//	data[9] = 0x00
//	data[10] = 0x00
//	data[11] = 0x00
//
//	data[12] = uint8(33)
//
//	pos := 13
//	for ; pos < 13+23; pos++ {
//		data[pos] = 0
//	}
//
//	// User
//	pos += copy(data[pos:], `root`)
//	data[pos] = 0x00
//	pos++
//
//	var authRespLEIBuf [9]byte
//	authRespLEI := append(authRespLEIBuf[:0], byte(uint64(len(auth))))
//	// auth [length encoded integer]
//	pos += copy(data[pos:], authRespLEI)
//	pos += copy(data[pos:], auth)
//
//	// db [null terminated string]
//
//	// Assume native client during response
//	pos += copy(data[pos:], "mysql_native_password")
//	data[pos] = 0x00
//
//	WritePacket(conn, data)
//}
//
//func WritePacket(conn net.Conn, data []byte) error {
//	length := len(data) - 4
//
//	data[0] = byte(length)
//	data[1] = byte(length >> 8)
//	data[2] = byte(length >> 16)
//	data[3] = SequenceId
//
//	n, err := conn.Write(data);
//	if err != nil || n != len(data){
//		fmt.Println("err", err)
//	}
//	SequenceId ++
//	return nil
//}
//
//// 权限验证
//func authPluginAllowed(pluginName string) bool {
//	// defines the supported auth plugins
//	var supportedAuthPlugins = []string{"mysql_native_password"}
//	for _, p := range supportedAuthPlugins {
//		if pluginName == p {
//			return true
//		}
//	}
//	return false
//}
//
//func CalcPassword(scramble, password []byte) []byte {
//	if len(password) == 0 {
//		return nil
//	}
//
//	// stage1Hash = SHA1(password)
//	crypt := sha1.New()
//	crypt.Write(password)
//	stage1 := crypt.Sum(nil)
//
//	// scrambleHash = SHA1(scramble + SHA1(stage1Hash))
//	// inner Hash
//	crypt.Reset()
//	crypt.Write(stage1)
//	hash := crypt.Sum(nil)
//
//	// outer Hash
//	crypt.Reset()
//	crypt.Write(scramble)
//	crypt.Write(hash)
//	scramble = crypt.Sum(nil)
//
//	// token = scrambleHash XOR stage1Hash
//	for i := range scramble {
//		scramble[i] ^= stage1[i]
//	}
//	return scramble
//}
