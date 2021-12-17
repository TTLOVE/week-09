package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"week-09/decoder"
)

func Serve() {
	listener, err := net.Listen("tcp", ":20000")
	if err != nil {
		fmt.Printf("listen fail,err:%s\n", err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept fail,err:%s\n", err)
			continue
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	data := &decoder.Data{}
	for {
		err := decoder.Decode(reader, data)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode fail, err:", err)
		}

		// 打印获取到的data信息
		fmt.Printf("data:%#v \n", data)
	}
}
