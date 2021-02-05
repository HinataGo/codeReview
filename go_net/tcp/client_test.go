package tcp

import (
	"fmt"
	"net"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	var buffers []byte
	hostPort := "localhost:8081"
	tcpAdd, err := net.ResolveTCPAddr("tcp4", hostPort)
	checkTcpErr(err)
	conn, err := net.DialTCP("tcp", nil, tcpAdd)
	defer conn.Close()
	rAddr := conn.RemoteAddr()

	w, err := conn.Write(buffers[0:])
	checkTcpErr(err)
	fmt.Println("Receive from server ", rAddr.String(), string(buffers[0:w]))

	r, err := conn.Read(buffers[0:])
	checkTcpErr(err)
	fmt.Println("Reply from server ", rAddr.String(), string(buffers[0:r]))
	os.Exit(0)
}

func checkTcpErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
