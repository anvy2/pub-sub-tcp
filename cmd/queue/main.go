package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"test/jsonHelper"
	"test/tcp"
	"time"
)

var waitTime int

func main() {
	c := make(chan jsonHelper.JSON)
	var wg sync.WaitGroup
	wg.Add(2)
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	inputPORT := ":" + arguments[1]
	inputServer, err := tcp.OpenConnection(inputPORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputPORT := ":" + arguments[2]
	outputServer, er := tcp.OpenConnection(outputPORT)

	if er != nil {
		fmt.Println(er)
		return
	}

	waitTime, err = strconv.Atoi(arguments[3])

	if err != nil {
		fmt.Println("Enter valid wait time in milliseconds")
	}

	defer inputServer.Listener.Close()
	defer outputServer.Listener.Close()
	go acceptInputConnection(&inputServer, c, &wg)
	go acceptOutputConnection(&outputServer, c, &wg)
	wg.Wait()
}

func acceptInputConnection(server *tcp.Server, c chan jsonHelper.JSON, wg *sync.WaitGroup) {
	defer wg.Done()
	defer server.Wg.Wait()
	for {
		conn, err := server.AcceptNewConnection()
		if err != nil {
			fmt.Println(err)
			return
		}
		go readFromConnection(conn, c, server)
	}

}

func readFromConnection(conn net.Conn, c chan jsonHelper.JSON, server *tcp.Server) {
	defer conn.Close()
	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		inputString := strings.TrimSpace(netData)
		if inputString == "STOP" {
			fmt.Println("Exiting incoming TCP server!")
			return
		}
		if jsonHelper.IsJSON(inputString) {
			data := jsonHelper.ChannelData{
				Data:      inputString,
				Timestamp: time.Now(),
			}
			jsonData, err := jsonHelper.EncodeToJson(data)
			if err != nil {
				fmt.Println(err)
			} else {
				c <- jsonData
			}
		} else {
			fmt.Println("Recieved data is not valid json")
		}
	}
}

func acceptOutputConnection(server *tcp.Server, c chan jsonHelper.JSON, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		conn, err := server.AcceptNewConnection()
		if err != nil {
			fmt.Println(err)
			return
		}
		go writeToConnection(conn, c, server)
	}
}

func writeToConnection(conn net.Conn, c chan jsonHelper.JSON, server *tcp.Server) {
	server.Wg.Add(1)
	defer conn.Close()
	defer server.Wg.Done()
	for {
		select {
		case o := <-c:
			conn.Write(o)
			conn.Write([]byte(string([]rune{'\n'})))
		default:
			if len(c) == 0 {
				time.Sleep(time.Millisecond * time.Duration(waitTime))
			}

		}
	}
}
