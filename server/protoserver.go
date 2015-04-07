package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/rpc"
	"runtime"

	"../PbTest"

	"github.com/golang/protobuf/proto"
)

type ServiceName struct{}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	fmt.Println("Staring Server..")
	listenAndServeRPC("tcp", ":8080")
	/*
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
	*/
}

func ReadReceivedData(data *PbTest.TestMessage) {
	msgItems := data.GetMessageItems()
	fmt.Println("Receiving data...")
	for _, item := range msgItems {
		fmt.Println(item)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Fatal error: ", err.Error())
	}
}

func handleProtoClient(conn net.Conn, c chan *PbTest.TestMessage) {
	fmt.Println("Connected!")
	defer conn.Close()
	var buf bytes.Buffer
	io.Copy(&buf, conn)
	pdata := new(PbTest.TestMessage)
	err := proto.Unmarshal(buf.Bytes(), pdata)
	checkError(err)
	c <- pdata
}

func listenAndServeRPC(network, addr string) {
	sn := new(ServiceName)
	rpc.Register(sn)
	listener, err := net.Listen("tcp", ":8080")
	checkError(err)
	defer listener.Close()
	rpc.Accept(listener)
}

func (s *ServiceName) Summate(tm *PbTest.TestMessage, response *PbTest.TestResponse) error {
	log.Println("SUMMING!")
	//response := new(PbTest.TestResponse)
	var sum int32
	sum = 0
	for _, msg := range tm.MessageItems {
		sum += *msg.ItemValue
	}

	response.FunctionName = proto.String("Summate")
	conversion := PbTest.TestResponse_StatusType(0)
	response.Status = &conversion
	response.Solution = proto.Int32(sum)
	return nil
}
