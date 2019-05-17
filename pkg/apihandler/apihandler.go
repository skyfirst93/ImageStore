package apihandler

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message    string
	HTTPStatus uint
}

func CreateAlbumHandler(w http.ResponseWriter, req *http.Request) {
	//Check if an album exits with the same name
	//
	fmt.Println("222")
}

func CreateImageHandler(w http.ResponseWriter, req *http.Request) {
	//Checkif image is present with same name
	//Check if all measurements are present

	//Create new image with the values being top left corner x,y and bottem right x,y
	newimage := image.NewRGBA(image.Rect(0, 0, 240, 240))
	//Create a variable with RGBA combinations
	blue := color.RGBA{0, 0, 255, 255}
	//Use Draw function to create the laypout and fill colors
	draw.Draw(newimage, newimage.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	//Create a file with the Name given by user
	filepointer, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(filepointer, newimage); err != nil {
		filepointer.Close()
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: "Image created", HTTPStatus: http.StatusOK})
}

func DeleteAlbumHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("222")

}

func DeleteImageHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("222")

}

func GetAlbums(w http.ResponseWriter, req *http.Request) {
	fmt.Println("222")

}

func GetImages(w http.ResponseWriter, req *http.Request) {
	fmt.Println("222")

}

func GetAlbumsById(w http.ResponseWriter, req *http.Request) {
	fmt.Println("222")

}

func GetImagesById(w http.ResponseWriter, req *http.Request) {
	fmt.Println("222")

}

func GetCreateNotification(w http.ResponseWriter, req *http.Request) {
	fmt.Println("222")

}

func GetDeleteNotification(w http.ResponseWriter, req *http.Request) {
	fmt.Println("222")

}
