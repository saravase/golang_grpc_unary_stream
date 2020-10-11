package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

// Write protocol buffer data into json file
func WriteProtoBufToJSONFile(message proto.Message, filename string) error {

	data, err := ProtoBufToJSON(message)
	if err != nil {
		return fmt.Errorf("Cannot marshal proto buf message into json file : %v", err)
	}

	// WriteFile writes data to a file named by filename.
	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("Cannot write json data to file : %v", err)
	}

	return nil
}

// Write protocol buffer data into binary file
func WriteProtoBufToBinaryFile(message proto.Message, filename string) error {

	// Marshal returns the wire-format encoding of message.
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("Cannot marshal proto buf message into binary file : %v", err)
	}

	// WriteFile writes data to a file named by filename.
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("Cannot write binary data to file : %v", err)
	}

	return nil
}

// Read protocol buffer data from binary file
func ReadProtoBufFromBinaryFile(filename string, message proto.Message) error {

	// ReadFile reads the file named by filename and returns the contents
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Cannot read binary data from file : %v", err)
	}

	// Unmarshal parses a wire-format message in data and places the decoded results in messagge
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("Cannot Unmarshal proto buf message from binary file : %v", err)
	}

	return nil
}
