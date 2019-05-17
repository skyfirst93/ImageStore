package apihandler

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)
//Set this as a Environment variable
var storagePath = "/akash"

type MultiValuesResponse struct {
	Values    []string
	HTTPStatus int
}

type Response struct {
        Message    string
        HTTPStatus int
}

type ErrorResponse struct {
	Error      error
	HTTPStatus int
}

func writeResponse(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{Message: message, HTTPStatus: status})
}
func writeMultiValuesResponse(w http.ResponseWriter, values []string, status int) {
        w.WriteHeader(status)
        json.NewEncoder(w).Encode(MultiValuesResponse{Values: values, HTTPStatus: status})
}

func writeErrorResponse(w http.ResponseWriter, errorMessage error, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Error: errorMessage, HTTPStatus: status})
}

func checkIfPathExists(dir string) bool {
	if _, err := os.Stat(dir); err == nil {
		return true
	}
	return false
}

func CreateAlbumHandler(w http.ResponseWriter, req *http.Request) {
	dir := "akash"
	if present := checkIfPathExists(storagePath+"/"+dir); !present {
		if err := os.MkdirAll(storagePath+"/"+dir, 0755); err != nil {
			log.Fatal(err)
		}
		writeResponse(w, "Album Created", http.StatusOK)
	} else {
		writeResponse(w, "Album Already Present", http.StatusConflict)
	}
}

func CreateImageHandler(w http.ResponseWriter, req *http.Request) {
	dir := "akash"
	imageName := "akash.png"
	if albumpresent := checkIfPathExists(storagePath + "/" +dir); albumpresent {
		if imagePresent := checkIfPathExists(storagePath + "/" +dir + "/" + imageName); !imagePresent {
			//Create new image with the values being top left corner x,y and bottem right x,y
			newImage := image.NewRGBA(image.Rect(0, 0, 240, 240))
			//Create a variable with RGBA combinations
			blue := color.RGBA{0, 0, 255, 255}
			//Use Draw function to create the laypout and fill colors
			draw.Draw(newImage, newImage.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
			//Create a file with the Name given by user
			filepointer, err := os.Create(storagePath + "/" +dir +"/"+ imageName)
			if err != nil {
				log.Fatal(err)
			}
			if err := png.Encode(filepointer, newImage); err != nil {
				filepointer.Close()
				log.Fatal(err)
			}
			writeResponse(w, "Image created", http.StatusOK)
			return
		}
		writeResponse(w, "Image Already Present", http.StatusConflict)
		return
	}
	writeResponse(w, "Album not Present", http.StatusConflict)
}

func DeleteAlbumHandler(w http.ResponseWriter, req *http.Request) {
	dir := "akash"
	if present := checkIfPathExists(storagePath + "/" +dir); present {
		// Remove the file.
		if err := os.Remove(storagePath + "/" +dir); err != nil {
			writeErrorResponse(w,err,http.StatusInternalServerError)
			return
		}
		writeResponse(w, "Album Deleted", http.StatusOK)
		return
	}
	writeResponse(w, "Album Not Present", http.StatusConflict)
}

func DeleteImageHandler(w http.ResponseWriter, req *http.Request) {
	dir := "akash"
	imageName := "akash.png"
	if albumpresent := checkIfPathExists(storagePath + "/" +dir); albumpresent {
		if imagePresent := checkIfPathExists(storagePath + "/" + dir + "/" + imageName); imagePresent {
			// Remove the file.
			if err := os.Remove(storagePath + "/" +dir+ "/" + imageName); err != nil {
				writeErrorResponse(w,err,http.StatusInternalServerError)
				return
			}
			writeResponse(w, "Image deleted", http.StatusOK)
			return
		} else {
			writeResponse(w, "Image Not Present", http.StatusConflict)
			return
		}
	}
	writeResponse(w, "Image not Present", http.StatusConflict)
}

func GetAlbumsList(w http.ResponseWriter, req *http.Request) {
	var albums []string
	files, err := ioutil.ReadDir(storagePath)
	if err != nil {
		writeErrorResponse(w,err,http.StatusInternalServerError)
                return
	}
	for _, filePointer := range files {
		albums = append(albums, filePointer.Name())
	}
	writeMultiValuesResponse(w, albums, http.StatusOK)
	//Note return all albums 
}

func GetImages(w http.ResponseWriter, req *http.Request) {
        var images []string
	albumName:="akash"
	if albumpresent := checkIfPathExists(storagePath + "/" +dir); albumpresent {
		files, err := ioutil.ReadDir(storagePath+"/"+albumName)
		if err != nil {
			writeErrorResponse(w,err,http.StatusInternalServerError)
			return
		}
		for _, filePointer := range files {
			images = append(images, filePointer.Name())
		}
		writeMultiValuesResponse(w, images, http.StatusOK)
		//Note return images not list of images 
	}
	writeResponse(w, "Album not Present", http.StatusConflict)

}

func GetImagesByName(w http.ResponseWriter, req *http.Request) {
        dir := "akash"
        imageName := "akash.png"
	var image []string
        if albumpresent := checkIfPathExists(storagePath + "/" +dir); albumpresent {
                if imagePresent := checkIfPathExists(storagePath + "/" +dir + "/" + imageName); imagePresent {
			// Note return image not name of images 
			image = append(image,imageName)
			writeMultiValuesResponse(w, image, http.StatusOK)
		        return
		}
		writeResponse(w, "Image not Present", http.StatusConflict)
		return
        }
        writeResponse(w, "Album not Present", http.StatusConflict)

}
// Note to put check on startup for storage path 
func GetCreateNotification(w http.ResponseWriter, req *http.Request) {
	fmt.Println("222")

}

func GetDeleteNotification(w http.ResponseWriter, req *http.Request) {
	fmt.Println("222")

}
