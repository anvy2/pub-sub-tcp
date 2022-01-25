package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"test/jsonHelper"
	"time"
)

type RecievedData struct {
	Data      string
	Timestamp time.Time
}

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
		reader := bufio.NewReader(conn)
		text, _ := reader.ReadBytes('\n')
		if string(text) == "STOP" {
			return
		}
		var data RecievedData
		err := jsonHelper.DecodeJson(text, &data)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		} else {
			fmt.Println(data.Data)
			fmt.Println("Recieved at: ", data.Timestamp)
		}

	}
}
