package main

import (
	api "ImageStore/pkg/apiroutes"
	"ImageStore/pkg/messaging"
	"fmt"
	"sync"
)
func init(){
	messaging.InitProducer("127.0.0.1:9092")
	messaging.InitConsumer("127.0.0.1:9092", "group")

}

func main() {
	fmt.Println("magic is happening on port 8081")
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go api.RunApi(&waitgroup, "127.0.0.1:8081")
	waitgroup.Wait()

}
