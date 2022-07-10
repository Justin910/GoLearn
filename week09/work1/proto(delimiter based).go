package work1

import (
	"bufio"
	"log"
	"net"
)

type DelimiterBased struct {
	Sep byte
	Num int
}

func (f *DelimiterBased) ServerHandler(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {

		buf, err := reader.ReadBytes(f.Sep)
		if err != nil {
			break
		}

		if len(buf) > 0 {
			buf = buf[:len(buf)-1]
		}

		log.Printf("Recv Length: %d, Buf: %s\n", len(buf), buf)
		f.Num++
	}
}

func (f *DelimiterBased) GetRecvPackageNum() int {
	return f.Num
}
