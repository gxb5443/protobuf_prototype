package main

import (
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	test := &Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Optionalgroup: &Test_OptionalGroup{
			RequiredField: proto.String("goodbye"),
		},
	}

	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("Marshal error: ", err)
	}
	newTest := &Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("UnMarshal error: ", err)
	}
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}

	log.Printf("Unmarshalled to: %+v", newTest)
}
