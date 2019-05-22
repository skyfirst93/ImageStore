package main

import (
	api "ImageStore/pkg/apiroutes"
	"ImageStore/pkg/messaging"
	"ImageStore/pkg/utils"
	"fmt"
	"log"
	"os"
	"sync"
)

// @APIVersion 1.0.0
// @APITitle IMAGE STORE SWAGGER API
// @APIDescription Image store for albums and images
// @Contact.name akash Pahal
// @BasePath http://127.0.0.1:8081/api/store

// @SubApi CREATE Album API [/create/album/{albumname}]
// @Title CreateAlbumHandler
// @Description Create album handler creates the album by name
// @Accept  json
// @Param   albumname     path    string     true        "Album Name"
// @Success 200 {array}  writeResponse
// @Failure 209 {object} writeResponse    "Album Name already present"
// @Router /api/store/create/album/{albumname} [get]

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
	} else {
		//Check if path exists and if not create it
		if _, err := os.Stat(utils.StoragePath); err != nil {
			if err := os.MkdirAll(utils.StoragePath, 0755); err != nil {
				log.Fatal(err)
			}
		}
	}

	utils.ServicePort = os.Getenv("SERVICE_PORT")
	if utils.ServicePort == "" {
		fmt.Println("Environment variable SERVICE_PORT undefined")
		os.Exit(1)
	}
	messaging.InitProducer(utils.MessageQueueAddr)
	messaging.InitConsumer(utils.MessageQueueAddr, "group")
}

func main() {
	fmt.Println("magic is happening on port 8081")
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go api.RunAPI(&waitgroup, ":"+utils.ServicePort)
	waitgroup.Wait()

}
