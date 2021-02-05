package tcp

import (
	"fmt"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	service := "localhost:8081"
	sAdd, err := net.ResolveTCPAddr("tcp4", service)
	checkTcpErr(err)

	listen, err := net.ListenTCP("tcp", sAdd)
	checkTcpErr(err)

	for {
		if conn, err := listen.Accept(); err != nil {
			continue
		} else {
			go handle(conn)
		}

	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	var buffer [512]byte
	for {
		r, err := conn.Read(buffer[0:])
		if err != nil {
			return
		}
		rAdd := conn.RemoteAddr()
		fmt.Println("Receive from client", rAdd.String(), string(buffer[0:r]))
		_, err2 := conn.Write([]byte("Welcome client!"))
		if err2 != nil {
			return
		}

	}

}
