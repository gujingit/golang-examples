package main

import (
	"io"
	"log"
	"net"
	"os"
)

/*
客户端
1. 从标准输入读入，发送给服务端；
2. 从服务端读取数据，写入到标准输出
*/

func main() {
	conn, e := net.Dial("tcp", "localhost:8080")
	if e != nil {
		log.Fatal(e)
	}
	defer func() {
		conn.Close()
	}()

	// 2. 从服务端读取数据，写入到标准输出
	// io读取是阻塞的
	go mustCopy(os.Stdout, conn)
	// 1. 从标准输入读入，发送给服务端
	mustCopy(conn, os.Stdin)


	/*
	1. 第一句加go，第二句不加：work
	2. 第一句不加go，第二句加： 一直阻塞在第一句，不会执行第二句
	3。 两句都加go，主程序瞬间退出
	*/

}

func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}
