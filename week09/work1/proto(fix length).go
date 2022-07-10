package work1

import (
	"bufio"
	"log"
	"net"
	"time"
)

type FixLength struct {
	SetRecvBufSize int
	Num            int
}

func (f *FixLength) ServerHandler(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	buf := make([]byte, f.SetRecvBufSize)
	for {
		conn.SetDeadline(time.Now().Add(time.Second * 60))
		n, err := reader.Read(buf)
		if err != nil {
			break
		}

		log.Printf("Recv Length: %d, Buf: %s\n", n, string(buf[:n]))
		if n != f.SetRecvBufSize {
			break
		}
		f.Num++
	}
}

func (f *FixLength) GetRecvPackageNum() int {
	return f.Num
}
