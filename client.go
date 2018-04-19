package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9999")
	defer conn.Close()

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	writer.WriteString("Hello!\n")
	writer.Flush()

	for {
		line, err := reader.ReadString('\n')
		if len(line) > 0 {
			fmt.Println("Read line from server: ", line)
			writer.WriteString(line)
			writer.Flush()
		}
		if err != nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

}
