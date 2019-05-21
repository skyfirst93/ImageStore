package messaging

import (
	"ImageStore/pkg/utils"
	"fmt"
	"os"
	"testing"
)

// init function initializes Message Bus
func init() {
	utils.MessageQueueAddr = os.Getenv("KAFKA_SERVICE")
	if utils.MessageQueueAddr == "" {
		fmt.Println("Environment variable KAFKA_SERVICE undefined")
		os.Exit(1)
	}
	utils.StoragePath = os.Getenv("STORAGE_PATH")
	if utils.StoragePath == "" {
		fmt.Println("Environment variable STORAGE_PATH undefined")
		os.Exit(1)
	}
	InitProducer(utils.MessageQueueAddr)
	InitConsumer(utils.MessageQueueAddr, "group")
}

//TestWriteMessage is used for testing writing messages onto messages bus
func TestWriteMessage(t *testing.T) {
	WriteMessage("Test", "Test")
	v := ReadMessage("Test")
	_ = v
}

//BenchmarkReadMessage is used for benchmarking Readmessage function
func BenchmarkReadMessage(b *testing.B) {
	WriteMessage("Test", "Test")
	for n := 0; n < b.N; n++ {
		ReadMessage("Test")
	}
}

//BenchmarkWriteMessage is used for benchmarking Readmessage function
func BenchmarkWriteMessage(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WriteMessage("Test", "Test")
	}
}
