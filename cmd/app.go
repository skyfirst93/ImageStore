package main

import (
	api "ImageStore/pkg/apiroutes"
	"ImageStore/pkg/messaging"
	"ImageStore/pkg/utils"
	"fmt"
	"os"
	"sync"
)

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

	messaging.InitProducer(utils.MessageQueueAddr)
	messaging.InitConsumer(utils.MessageQueueAddr, "group")
}

func main() {
	fmt.Println("magic is happening on port 8081")
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go api.RunApi(&waitgroup, "127.0.0.1:8081")
	waitgroup.Wait()

}
