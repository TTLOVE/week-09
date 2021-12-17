package decoder

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"week-09/client"
)

type Data struct {
	Package_length   int32
	Header_length    int16
	Protocol_version int16
	Operation        int32
	Sequence         int32
	Body             string
}

// Decode 是从reader里面读取解码数据信息
// goim的数据结构分为两部分：一部分为内容信息，一部分为非内容信息
// 内容信息分布：
// Package Length，包长度(4 byte)
// Header Length，头长度(2 byte)
// Protocol Version，协议版本(2 byte)
// Operation，操作码(4 byte)
// Sequence 请求序号 ID(4 byte)
// 这里是分开步骤去一步步获取对应的信息的长度，然后在最后获取body的信息
// 后面可以根据Data信息去读取header的信息
func Decode(reader *bufio.Reader, data *Data) error {
	err := handleOutBody(reader, data)
	if err != nil {
		return err
	}

	// 读取body信息
	bodyByte := make([]byte, data.Package_length)
	_, err = reader.Read(bodyByte)
	if err != nil {
		return err
	}

	data.Body = string(bodyByte)
	return nil
}

// 处理非body的数据
func handleOutBody(reader *bufio.Reader, data *Data) error {
	bufferByte := make([]byte, client.LengthMaps[client.Package_length])
	for column, length := range client.LengthMaps {
		setByte := bufferByte[:length]
		_, err := reader.Read(setByte)
		if err != nil {
			return err
		}

		lengthBuff := bytes.NewBuffer(setByte)
		switch column {
		case client.Package_length:
			err = binary.Read(lengthBuff, binary.LittleEndian, &data.Package_length)
		case client.Header_length:
			err = binary.Read(lengthBuff, binary.LittleEndian, &data.Header_length)
		case client.Protocol_version:
			err = binary.Read(lengthBuff, binary.LittleEndian, &data.Protocol_version)
		case client.Operation:
			err = binary.Read(lengthBuff, binary.LittleEndian, &data.Operation)
		case client.Sequence:
			err = binary.Read(lengthBuff, binary.LittleEndian, &data.Sequence)
		}
		if err != nil {
			return err
		}
	}

	return nil
}
