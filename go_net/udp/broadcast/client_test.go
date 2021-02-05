package broadcast

import (
	"fmt"
	"net"
	"os"
	"testing"
)

/*
ParseIP将s解析为IP地址，并返回结果。
字符串s可以采用IPv4点分十进制（“ 192.0.2.1”），
IPv6（“ 2001：db8 :: 68”）或IPv4映射的IPv6（“ :: ffff：192.0.2.1”）形式。
如果s不是IP地址的有效文本表示形式，则ParseIP返回nil
*/
func TestCli(t *testing.T) {
	ip := net.ParseIP("255.255.255.255")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: 9981}
	conn, err := net.ListenUDP("udp", srcAddr)
	if err != nil {
		fmt.Println(err)
	}
	n, err := conn.WriteToUDP([]byte("hello"), dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	data := make([]byte, 1024)
	n, _, err = conn.ReadFrom(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("read %s from : %s \n", data[:n], conn.RemoteAddr())
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
