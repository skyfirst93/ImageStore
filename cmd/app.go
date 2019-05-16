package main

import (
	"fmt"
	"sync"
	api "ImageStore/pkg/imagestoreapi"
)

func main() {
	fmt.Println("magic is happening on port 8081")
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go api.RunApi(&waitgroup, "127.0.0.1:8081")
	waitgroup.Wait()

}
