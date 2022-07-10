package work1

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type LengthFieldBased struct {
	Num int
}

type Payload struct {
}

var (
	ErrDataLenError   = errors.New("Data Length Error")
	ErrDataIncomplete = errors.New("Data Incomplete")
)

func (f *LengthFieldBased) ServerHandler(conn net.Conn) {
	defer conn.Close()
	for {
		conn.SetDeadline(time.Now().Add(time.Second * 60))
		buf, err := f.Read(conn)
		if err != nil {
			fmt.Println(err)
			break
		}

		log.Printf("Recv Length: %d, Buf: %s\n", len(buf), buf)
		f.Num++
	}
}

func (f *LengthFieldBased) GetRecvPackageNum() int {
	return f.Num
}

func (f *LengthFieldBased) Read(reader io.Reader) ([]byte, error) {

	// 读取Header长度字段
	buf := make([]byte, 2)
	n, err := reader.Read(buf)
	if err != nil {
		return nil, err
	}

	if n != 2 {
		return nil, ErrDataIncomplete
	}

	l := binary.BigEndian.Uint16(buf)

	nbuf := make([]byte, l)
	n, err = io.ReadFull(reader, nbuf)
	if err != nil {
		return nil, err
	}

	if uint16(n) != l {
		return nil, ErrDataIncomplete
	}

	return nbuf[:n], nil
}

func (f *LengthFieldBased) Encode(b []byte) []byte {

	//前两个字节标识body长度
	headerBuf := make([]byte, 2)

	binary.BigEndian.PutUint16(headerBuf, uint16(len(b)))

	buf := append(headerBuf, b...)

	return buf
}
