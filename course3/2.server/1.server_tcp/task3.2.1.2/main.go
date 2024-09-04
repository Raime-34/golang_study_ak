package main

import (
	"fmt"
	"net"
	"strings"
)

const (
	standartRes = "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html\r\n" +
		"Connection: close\r\n\r\n" +
		"<!DOCTYPE html>\n<html>\n<head>\n    <title>Webserver</title>\n</head>\n<body>\n    hello world\n</body>\n</html>\n"
	notFoundRes = "HTTP/1.1 404 Not Found\r\n" +
		"Content-Type: text/html\r\n" +
		"Connection: close\r\n\r\n" +
		"<!DOCTYPE html>\n<html>\n<head>\n    <title>404 Not Found</title>\n</head>\n<body>\n    <h1>404 Not Found</h1>\n    <p>The page you are looking for does not exist.</p>\n</body>\n</html>\n"
)

func main() {
	listener, _ := net.Listen("tcp", ":8000")
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil {
		return
	}

	if req := string(buff[:n]); !strings.Contains(req, "GET / HTTP/1.1") {
		_, err = conn.Write([]byte(notFoundRes))
	} else {
		_, err = conn.Write([]byte(standartRes))
	}

	if err != nil {
		fmt.Println(err)
	}
}
