package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	conn, err := net.Dial("tcp", "127.0.0.1:"+arguments[1])

	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {
		fmt.Print("Enter JSON to send : ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadBytes('\n')
		if string(text) == "STOP" {
			return
		}
		conn.Write(text)
	}
}
