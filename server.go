package main;

import (
	"net"
	"strconv"
	"bufio"
	"fmt"
)

func main() {
	startServer(9999)
}

func startServer(port int) {
	server, _ := net.Listen("tcp", "localhost:"+strconv.Itoa(port))
	defer server.Close()

	conn, _ := server.Accept()

	handleRequest(conn)
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	bufferedReader := bufio.NewReader(conn)
	bufferedWriter := bufio.NewWriter(conn)

	for {
		// Note: .ReadLine() is complex than this, because it may read part of line when the line is long
		line, err := bufferedReader.ReadString('\n')
		if len(line) > 0 {
			fmt.Println("Read line from client: ", line)
			bufferedWriter.WriteString(string(line))
			bufferedWriter.Flush()
		}
		if err != nil {
			break
		}
	}
}
