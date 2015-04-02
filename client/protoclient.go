package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"

	"../PbTest"

	"github.com/golang/protobuf/proto"
)

type Headers []string

func main() {
	filename := flag.String("f", "CSVValue.csv", "Enter the filename of CSV to read from")
	dest := flag.String("d", "127.0.0.1:2110", "Enter the destination socket address")
	flag.Parse()
	data, err := retrieveDataFromFile(filename)
	checkError(err)
	sendDataToDest(data, dest)
}

func checkError(err Error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func retrieveDataFromFile(fname *string) ([]byte, error) {
	file, err := os.Open(fname)
	checkError(err)
	defer file.Close()
	csvreader := csv.NewReader(file)
	var headers Headers
	headers, err = csvreader.Read()
	checkError(err)
	itemIdIndex := headers.getHeaderIndex("itemid")
	itemNameIndex := headers.getHeaderIndex("itemname")
	itemValueIndex := headers.getHeaderIndex("itemvalue")
	itemTypeIndex := headers.getHeaderIndex("itemType")
	Testmessage := new(PbTest.TestMessage)
	for {
		record, err := csvreader.Read()
		if err != io.EOF {
			checkError(err)
		} else {
			break
		}
		msgItem := new(PbTest.TestMessage_MsgItem)
		itemid, err := strconv.Atoi(record[itemIdIndex])
		checkError(err)
		msgItem.Id = proto.Int32(itemid)
		msgItem.ItemName = &record[itemNameIndex]
		itemValue, err := strconv.Atoi(record[itemValueIndex])
		checkError(err)
		msgItem.ItemValue = proto.Int32(itemValue)
		ttype, err := strconv.Atoi(record[itemTypeIndex])
		checkError(err)
		converted_ttype := PbTest.TestMessage_TType(ttype)
		msgItem.TransactionType = &converted_ttype
		TestMessage.MessageItems = append(TestMessage.MessageItems, msgItem)
		fmt.Println(record)
	}
	return proto.Marshal(TestMessage)
}

func (h Headers) getHeaderIndex(headername string) int {
	if len(headername) >= 2 {
		for index, s := range h {
			if s == headername {
				return index
			}
		}
	}
	return -1
}

func sendDataToDest(data []byte, dst *string) {
	conn, err := net.Dial("tcp", dst)
	checkError(err)
	n, err := conn.Write(data)
	checkError(err)
	fmt.Println("Sent " + strconv.Itoa(n) + " bytes")
}
