package apihandler

import (
	"ImageStore/pkg/messaging"
	"ImageStore/pkg/utils"
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

	"github.com/gorilla/mux"
)

//Notification struct defines the structure for create and delete notification
type Notification struct {
	NotiType  string
	info      string
	timestamp string
}

//MultiValuesResponse defines structure for multiple values of Images and Albums
type MultiValuesResponse struct {
	Values     []string `json:"values"`
	HTTPStatus int      `json:"HttpStatus"`
}

//Response structure defines basic message and status of response
type Response struct {
	Message    string `json:"message"`
	HTTPStatus int    `json:"HttpStatus"`
}

//ErrorResponse structure defines error message and status of response
type ErrorResponse struct {
	Error      error `json:"error"`
	HTTPStatus int   `json:"HttpStatus"`
}

//writeResponse write the message passed and status to http Response writer
func writeResponse(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{Message: message, HTTPStatus: status})
}

//writeMultiValuesResponse write a slice of values(images/albums) and http
//to http Response writer
func writeMultiValuesResponse(w http.ResponseWriter, values []string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(MultiValuesResponse{Values: values, HTTPStatus: status})
}

//writeErrorResponse writes error message and http status
//to http Response writer
func writeErrorResponse(w http.ResponseWriter, errorMessage error, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Error: errorMessage, HTTPStatus: status})
}

//checkIfPathExists function check if the path passed exist and return true
// else return false
func checkIfPathExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

//CreateAlbumHandler is handler function for creating an Album
func CreateAlbumHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	albumPath := utils.StoragePath + "/" + params["albumname"]
	if albumPresent := checkIfPathExists(albumPath); !albumPresent {
		if err := os.MkdirAll(albumPath, 0755); err != nil {
			log.Fatal(err)
		}
		writeResponse(w, "Album Created", http.StatusOK)
	} else {
		writeResponse(w, "Album Already Present", http.StatusConflict)
	}
}

func CreateImageHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	albumPath := utils.StoragePath + "/" + params["albumname"]
	imagePath := albumPath + "/" + params["imagename"]

	if albumpresent := checkIfPathExists(albumPath); albumpresent {
		if imagePresent := checkIfPathExists(imagePath); !imagePresent {
			//Create new image with the values being top left corner x,y and bottem right x,y
			newImage := image.NewRGBA(image.Rect(0, 0, 240, 240))
			//Create a variable with RGBA combinations
			blue := color.RGBA{0, 0, 255, 255}
			//Use Draw function to create the laypout and fill colors
			draw.Draw(newImage, newImage.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
			//Create a file with the Name given by user
			filepointer, err := os.Create(imagePath)
			if err != nil {
				log.Fatal(err)
			}
			if err := png.Encode(filepointer, newImage); err != nil {
				filepointer.Close()
				log.Fatal(err)
			}
			writeResponse(w, "Image created", http.StatusOK)
			messaging.WriteMessage("Image Created", "IMAGE")
			return
		}
		writeResponse(w, "Image Already Present", http.StatusConflict)
		return
	}
	writeResponse(w, "Album not Present", http.StatusConflict)
}

func DeleteAlbumHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	albumPath := utils.StoragePath + "/" + params["albumname"]
	if present := checkIfPathExists(albumPath); present {
		// Remove the file.
		if err := os.Remove(albumPath); err != nil {
			writeErrorResponse(w, err, http.StatusInternalServerError)
			return
		}
		writeResponse(w, "Album Deleted", http.StatusOK)
		return
	}
	writeResponse(w, "Album Not Present", http.StatusConflict)
}

func DeleteImageHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	albumPath := utils.StoragePath + "/" + params["albumname"]
	imagePath := albumPath + "/" + params["imagename"]

	if albumpresent := checkIfPathExists(albumPath); albumpresent {
		if imagePresent := checkIfPathExists(imagePath); imagePresent {
			// Remove the file.
			if err := os.Remove(imagePath); err != nil {
				writeErrorResponse(w, err, http.StatusInternalServerError)
				return
			}
			writeResponse(w, "Image deleted", http.StatusOK)
			messaging.WriteMessage("Image Deleted", "DELETE-IMAGE")
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
	files, err := ioutil.ReadDir(utils.StoragePath)
	if err != nil {
		writeErrorResponse(w, err, http.StatusInternalServerError)
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

	params := mux.Vars(req)
	albumPath := utils.StoragePath + "/" + params["albumname"]

	if albumpresent := checkIfPathExists(albumPath); albumpresent {
		files, err := ioutil.ReadDir(albumPath)
		if err != nil {
			writeErrorResponse(w, err, http.StatusInternalServerError)
			return
		}
		for _, filePointer := range files {
			images = append(images, filePointer.Name())
		}
		writeMultiValuesResponse(w, images, http.StatusOK)
		return
		//Note return images not list of images
	}
	writeResponse(w, "Album not Present", http.StatusConflict)
}

func GetImagesByName(w http.ResponseWriter, req *http.Request) {
	var image []string

	params := mux.Vars(req)
	albumPath := utils.StoragePath + "/" + params["albumname"]
	imagePath := albumPath + "/" + params["imagename"]

	if albumpresent := checkIfPathExists(albumPath); albumpresent {
		if imagePresent := checkIfPathExists(imagePath); imagePresent {
			// Note return image not name of images
			image = append(image, params["imagename"])
			http.ServeFile(w, req, imagePath)
			//writeMultiValuesResponse(w, image, http.StatusOK)
			return
		}
		writeResponse(w, "Image not Present", http.StatusConflict)
		return
	}
	writeResponse(w, "Album not Present", http.StatusConflict)
	//Note the return status

}

// Note to put check on startup for storage path
func GetCreateNotification(w http.ResponseWriter, req *http.Request) {
	message := messaging.ReadMessage("IMAGE")
	if message != nil {
		writeMultiValuesResponse(w, message, http.StatusOK)
		return
	}
	writeResponse(w, "No Create Notification", http.StatusConflict)
	//Note the return status
}

func GetDeleteNotification(w http.ResponseWriter, req *http.Request) {
	if message := messaging.ReadMessage("DELETE-IMAGE"); message != nil {
		writeMultiValuesResponse(w, message, http.StatusOK)
		return
	}
	writeResponse(w, "No Delete Notification", http.StatusConflict)
	//Note the return status
}

func Swagger(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hepepepepeppepep")
	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, r, "swagger.json")
}
