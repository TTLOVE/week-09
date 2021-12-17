package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

const (
	Package_length = iota
	Header_length
	Protocol_version
	Operation
	Sequence
)

var LengthMaps = [5]int{4, 2, 2, 4, 4}

// 客户端模拟请求
func Request() {
	conn, err := net.Dial("tcp", ":20000")
	if err != nil {
		fmt.Printf("conn fail, err:%s \n", err)
	}
	msgs := [5]string{
		"我是模拟发送goim数据",
		"我是模拟发送goim数据，我是模拟发送goim数据",
		"我是模拟发送goim数据，我是模拟发送goim数据我，是模拟发送goim数据",
		"我是模拟发送goim数据，我是模拟发送goim数据，我是模拟发送goim数据，我是模拟发送goim数据",
		"我是模拟发送goim数据，我是模拟发送goim数据，我是模拟发送goim数据，我是模拟发送goim数据，我是模拟发送goim数据",
	}
	for _, msg := range msgs {
		body, err := encodeData(msg)
		if err != nil {
			fmt.Printf("encode body fail, err:%s \n", err)
			return
		}

		log.Println("len:", len(body), "string:", string(body))
		_, err = conn.Write(body)
		if err != nil {
			fmt.Printf("write fail,err:%s\n", err)
			return
		}
	}
}

// 对数据进行编码处理
func encodeData(body string) ([]byte, error) {
	var pkg = new(bytes.Buffer)
	var err error
	var bodyByte = []byte(body)
	// 循环写入消息头
	for key, length := range LengthMaps {
		if key == Package_length {
			err = binary.Write(pkg, binary.LittleEndian, int32(len(bodyByte)))
		} else {
			switch length {
			case 2:
				err = binary.Write(pkg, binary.LittleEndian, int16(length))
			case 4:
				err = binary.Write(pkg, binary.LittleEndian, int32(length))
			default:
				continue
			}
		}

		if err != nil {
			fmt.Printf("write key %d fail, err :%s\n", key, err)
			return nil, err
		}
	}

	// 写入消息体
	err = binary.Write(pkg, binary.LittleEndian, bodyByte)
	if err != nil {
		fmt.Printf("write body fail, err :%s\n", err)
		return nil, err
	}

	return pkg.Bytes(), nil
}
