package main

import (
	"ImageStore/pkg/messaging"
)

func main() {
	messaging.InitProducer("127.0.0.1:9092", "akash")
	messaging.WriteMessage("this is a message", "akash")
}
