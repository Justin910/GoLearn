package work1

import (
	"context"
	"net"
	"time"
)

type ServerProxy interface {
	ServerHandler(conn net.Conn)
}

type ClientProxy interface {
	ClientHandle(conn net.Conn)
}

func ListenServer(ctx context.Context, addr string, proxy ServerProxy) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		l.Close()
	}()

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				time.Sleep(time.Second)
				continue
			}

			go proxy.ServerHandler(conn)
		}
	}()

	return nil
}
