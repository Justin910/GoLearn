package work1

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"
)

func Test_FixLength(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	length := 100
	sendPackageNum := 10

	f := new(FixLength)
	f.SetRecvBufSize = length
	err := ListenServer(ctx, ":9999", f)
	if err != nil {
		t.Fatalf("Listen Error. Err: %s", err.Error())
	}

	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()

	sb := strings.Builder{}
	for i := 0; i < sendPackageNum; i++ {
		//writer := bufio.NewWriter(conn)
		sb.WriteString(strings.Repeat("1", f.SetRecvBufSize))
	}

	writer := bufio.NewWriter(conn)
	writer.WriteString(sb.String())
	writer.Flush()

	time.Sleep(time.Second * 1)

	fmt.Println(f.GetRecvPackageNum())

	if f.GetRecvPackageNum() != sendPackageNum {
		t.Fatalf("Recv Package Error")
	}
}

func Test_DelimiterBased(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sendPackageNum := 10

	f := new(DelimiterBased)
	f.Sep = '\n'
	err := ListenServer(ctx, ":9998", f)
	if err != nil {
		t.Fatalf("Listen Error. Err: %s", err.Error())
	}

	conn, err := net.Dial("tcp", "127.0.0.1:9998")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()

	sb := strings.Builder{}
	for i := 0; i < sendPackageNum; i++ {
		//writer := bufio.NewWriter(conn)
		sb.WriteString(strings.Repeat("1", i+100))
		sb.WriteString("\n")
	}

	writer := bufio.NewWriter(conn)
	writer.WriteString(sb.String())
	writer.Flush()

	time.Sleep(time.Second * 1)

	fmt.Println(f.GetRecvPackageNum())

	if f.GetRecvPackageNum() != sendPackageNum {
		t.Fatalf("Recv Package Error")
	}
}

func Test_LengthFieldBased(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sendPackageNum := 10

	f := new(LengthFieldBased)
	err := ListenServer(ctx, ":9997", f)
	if err != nil {
		t.Fatalf("Listen Error. Err: %s", err.Error())
	}

	conn, err := net.Dial("tcp", "127.0.0.1:9997")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()

	b := make([]byte, 0)
	for i := 0; i < sendPackageNum; i++ {
		b = append(b, f.Encode([]byte(strings.Repeat("1", 250+i)))...)
	}

	// 一次性发10个包，确定是否有粘包
	writer := bufio.NewWriter(conn)
	writer.Write(b)
	writer.Flush()

	time.Sleep(time.Second * 1)

	fmt.Println(f.GetRecvPackageNum())

	if f.GetRecvPackageNum() != sendPackageNum {
		t.Fatalf("Recv Package Error")
	}
}
