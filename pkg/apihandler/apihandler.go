package apihandler

import "net/http"

func CreateAlbumHandler(w http.ResponseWriter, req *http.Request) {
	//Check if an album exits with the same name
	//
}

func CreateImageHandler(w http.ResponseWriter, req *http.Request) {
	//Checkif image is present with same name
	//Check if all measurements are present

}

func DeleteAlbumHandler(w http.ResponseWriter, req *http.Request) {
}

func DeleteImageHandler(w http.ResponseWriter, req *http.Request) {

}

func GetAlbums(w http.ResponseWriter, req *http.Request) {
}

func GetImages(w http.ResponseWriter, req *http.Request) {
}

func GetAlbumsById(w http.ResponseWriter, req *http.Request) {

}

func GetImagesById(w http.ResponseWriter, req *http.Request) {

}

func GetCreateNotification(w http.ResponseWriter, req *http.Request) {

}

func GetDeleteNotification(w http.ResponseWriter, req *http.Request) {

}
