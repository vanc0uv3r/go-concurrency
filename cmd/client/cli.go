package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	addr string
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	buff := make([]byte, 1024)

	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Occured error with reading stdin: ", err.Error())
		}
		fmt.Println("MSG to serv", line)
		_, err = conn.Write(line)
		if err != nil {
			fmt.Println("Occured error while sending request: ", err.Error())
		}

		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println("Occured error while sending request: ", err.Error())
		}
		fmt.Println("RESPONSE ", string(buff[:n]))
	}
}

func init() {
	flag.StringVar(&addr, "addr", "127.0.0.1:8888", "host:port to connect to server")
}