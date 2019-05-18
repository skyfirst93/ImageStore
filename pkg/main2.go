package main

import (
	"ImageStore/pkg/messaging"
)

func main() {
	messaging.InitConsumer("127.0.0.1:9092", "group", "akash")
	var topics []string
	topics = append(topics, "akash")
	messaging.ReadMessage(topics)
}
