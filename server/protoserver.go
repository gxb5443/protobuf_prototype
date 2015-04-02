package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"

	"../PbTest"

	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Staring Server..")
	c := make(chan *PbTest.TestMessage)
	go func() {
		for {
			message := <-c
			ReadReceivedData(message)
		}
	}()
	listener, err := net.Listen("tcp", ":8080")
	checkError(err)
	for {
		if conn, err := listener.Accept(); err == nil {
			go handleProtoClient(conn, c)
		} else {
			continue
		}
	}
}

func ReadReceivedData(data *PbTest.TestMessage) {
	msgItems := items.GetMessageItems()
	fmt.Println("Receiving data...")
	for _, item := range msgItems {
		fmt.Println(item)
	}
}

func checkError(err Error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handleProtoClient(conn net.Conn, c chan *PbTest.Testmessage) {
	fmt.Println("Connected!")
	defer conn.Close()
	var buf bytes.Buffer
	io.Copy(&buf, conn)
	pdata := new(PbTest.TestMessage)
	err = proto.Unmarshal(buf, pdata)
	checkError(err)
	c <- pdata
}
