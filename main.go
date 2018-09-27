package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/LeeTrent/protobuf-example-go/src/simple"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()
	readAndWriteDemo(sm)
	jsonDemo(sm)
}

func readAndWriteDemo(sm proto.Message) {
	fmt.Println("\nWrite simple message to file ...")
	writeToFile("simple.bin", sm)

	fmt.Println("\nRead simple message from file ...")
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println(sm2)
}

func jsonDemo(sm proto.Message) {
	fmt.Println("\nConvert simple message to JSON ...")
	smAsJSON := toJSON(sm)
	fmt.Println(smAsJSON)

	fmt.Println("\nConvert JSON back to simple message ...")
	smFromJSON := &simplepb.SimpleMessage{}
	fromJSON(smAsJSON, smFromJSON)
	fmt.Println(smFromJSON)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("[toJSON] - Error encountered when invoking marshaler.MarshalToString", err)
		return ""
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("[fromJSON] - Error encountered when invoking marshaler.UnmarshalString", err)
	}
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("[writeToFile] - Error encountered when invoking proto.Marshal", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("[writeToFile] - Error encountered when invoking ioutil.WriteFile", err)
		return err
	}

	fmt.Println("Data has been successfully written to disk")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("[readFromFile] - Error encountered when invoking ioutil.ReadFile", err)
		return err
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("[readFromFile] - Error encountered when invoking ioutil.Unmarshal", err)
		return err
	}

	fmt.Println("Data has successfully been read from file on disk")
	return nil
}

func doSimple() *simplepb.SimpleMessage {

	fmt.Println("\nCreate simple message ...")
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}
	fmt.Println(&sm)

	// fmt.Println(sm)
	// sm.Name = "I renamed you"
	// fmt.Println(sm)
	// fmt.Println("The ID is: ", sm.GetId())

	return &sm
}
