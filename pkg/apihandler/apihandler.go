package apihandler

import (
	"ImageStore/pkg/messaging"
	"ImageStore/pkg/utils"
	"encoding/json"
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
// @SubApi Create Album API [/create/album/{albumname}]
// @Title CreateAlbumHandler
// @Description Create album handler creates the album by name
// @Accept  json
// @Param   albumname     path    string     true        "Album Name"
// @Success 200 {object}  Response
// @Failure 209 {object} Response    "Album Name already present"
// @Resource /api/store
// @Router /api/store/create/album/{albumname} [POST]
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

//CreateImageHandler is handler function for creating an image
//and return success or failure
//Note add .png
// @SubApi Create Image API [/create/image/{albumname}/{imagename}]
// @Title Create Image Handler
// @Description Create Image handler creates the image by name
// @Accept  json
// @Param   albumname     path    string     true        "Album Name"
// @Param   imagename     path    string     true        "Image Name"
// @Success 200 {object}  Response
// @Failure 404 {object} Response    "Album Name already present"
// @Resource /api/store
// @Router /api/store/create/image/{albumname}/{imagename} [POST]
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
	writeResponse(w, "Album not Present", http.StatusNotFound)
}

//DeleteAlbumHandler is handler function for deleting an album
//and return suucess or failure
// @SubApi Delete Album API [/delete/image/{albumname}]
// @Title Delete Album Handler
// @Description DeleteAlbumHandler is handler function for deleting an album
// @Accept  json
// @Param   albumname     path    string     true        "Album Name"
// @Success 200 {object}  Response
// @Failure 404 {object} Response    "Album Name already present"
// @Resource /api/store
// @Router /api/store/delete/album/{albumname} [DELETE]
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
	writeResponse(w, "Album Not Present", http.StatusNotFound)
}

//DeleteImageHandler is handler function for deleting an image
//and return suucess or failure
// @SubApi Delete Image API [/delete/image/{albumname}/{imagename}]
// @Title Delete Image Handler
// @Description Create Image handler creates the image by name
// @Accept  json
// @Param   albumname     path    string     true        "Album Name"
// @Param   imagename     path    string     true        "Image Name"
// @Success 200 {object}  Response
// @Failure 404 {object} Response
// @Resource /api/store
// @Router /api/store/delete/image/{albumname}/{imagename} [DELETE]
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
		}
		writeResponse(w, "Image Not Present", http.StatusNotFound)
		return
	}
	writeResponse(w, "Album not Present", http.StatusNotFound)
}

//GetAlbumsList is handler function for getting list of albums
// @SubApi Get Albums list API [/albums]
// @Title Get Albums List Handler
// @Description GetAlbumsList is handler function for getting list of albums
// @Accept  json
// @Success 200 {object}  MultiValuesResponse
// @Resource /api/store
// @Router /api/store/albums [GET]
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

//GetImages is handler function for getting list of image
//and returning the list of images
// @SubApi Get Images list API [/images/{albumname}/{imagename}]
// @Title Get Images List Handler
// @Description GetImages is handler function for getting list of image in an album
// @Accept  json
// @Param   albumname     path    string     true        "Album Name"
// @Success 200 {object}  Response
// @Failure 404 {object} Response
// @Resource /api/store
// @Router /api/store/images/{albumname}/ [GET]
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
	writeResponse(w, "Album not Present", http.StatusNotFound)
}

//GetImagesByName is handler function for getting an image
//and returning the image
// @SubApi GetImagesByName API [/images/{albumname}/{imagename}]
// @Title GetImagesByName Handler
// @Description GetImagesByName is handler function for getting an image in an album
// @Accept  json
// @Param   albumname     path    string     true        "Album Name"
// @Param   imagename     path    string     true        "Image Name"
// @Success 200 {object}  Response
// @Failure 404 {object} Response
// @Resource /api/store
// @Router /api/store/images/{albumname}/{imagename} [GET]
func GetImagesByName(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	albumPath := utils.StoragePath + "/" + params["albumname"]
	imagePath := albumPath + "/" + params["imagename"]

	if albumpresent := checkIfPathExists(albumPath); albumpresent {
		if imagePresent := checkIfPathExists(imagePath); imagePresent {
			// Note return image not name of images
			http.ServeFile(w, req, imagePath)
			return
		}
		writeResponse(w, "Image not Present", http.StatusNotFound)
		return
	}
	writeResponse(w, "Album not Present", http.StatusNotFound)
	//Note the return status

}

// Note to put check on startup for storage path

//GetCreateNotification is handler function for getting the list of notification
//of Images created
// @SubApi Get Create Notification API [/api/store/notification/create]
// @Title Get Create Notification Handler
// @Description Get Create Notification is handler function for getting the list of notification
// @Accept  json
// @success 200 {object} MultiValuesResponse   "Create Notifications"
// @Failure 204 {object} Response    "No more Create notification"
// @Resource /api/store
// @Router /api/store/notification/create [GET]
func GetCreateNotification(w http.ResponseWriter, req *http.Request) {
	message := messaging.ReadMessage("IMAGE")
	if message != nil {
		writeMultiValuesResponse(w, message, http.StatusOK)
		return
	}
	writeResponse(w, "No Create Notification", http.StatusNoContent)
	//Note the return status
}

//GetDeleteNotification is handler function for getting the list of notification of
//Images deleted
// @SubApi GetDeleteNotification API [/notification/delete]
// @Title Get Delete Notification Handler
// @Description GetDeleteNotification is handler function for getting the list of notification of Images deleted
// @Accept  json
// @Success 200 {object}  MultiValuesResponse	"Delete Notifications"
// @Failure 204 {object} Response    "No more Delete notification"
// @Resource /api/store
// @Router /api/store/notification/delete [GET]
func GetDeleteNotification(w http.ResponseWriter, req *http.Request) {
	if message := messaging.ReadMessage("DELETE-IMAGE"); message != nil {
		writeMultiValuesResponse(w, message, http.StatusOK)
		return
	}
	writeResponse(w, "No Delete Notification", http.StatusNoContent)
	//Note the return status
}
