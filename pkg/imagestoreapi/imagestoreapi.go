package imagestoreapi

import (
	"fmt"
	"net/http"
	"sync"
	api "ImageStore/pkg/apihandler"

	"github.com/gorilla/mux"
)

//RunApi is used to start listioning on the endpoint and
// Call RunAPIOnRouter function
func RunApi(waitgroup *sync.WaitGroup, endpoint string) {
	r := mux.NewRouter()
	RunAPIOnRouter(r)
	err := http.ListenAndServe(endpoint, r)
	if err != nil {
		fmt.Println("Error-in-REST-Server")
	}
	waitgroup.Done()
}

//RunAPIOnRouter is used for setting the routes for the API's
func RunAPIOnRouter(r *mux.Router) {

	apirouter := r.PathPrefix("/api/store").Subrouter()

	apirouter.Methods("POST").Path("/create/album").HandlerFunc(api.CreateAlbumHandler)
	apirouter.Methods("DELETE").Path("/delete/album").HandlerFunc(api.DeleteAlbumHandler)

	apirouter.Methods("POST").Path("/create/image").HandlerFunc(api.CreateImageHandler)
	apirouter.Methods("DELETE").Path("/delete/image").HandlerFunc(api.DeleteImageHandler)

	apirouter.Methods("GET").Path("/images").HandlerFunc(api.GetImages)
	apirouter.Methods("GET").Path("/images/{id}").HandlerFunc(api.GetImagesById)
	apirouter.Methods("GET").Path("/albums").HandlerFunc(api.GetAlbums)
	apirouter.Methods("GET").Path("/albums/{id}").HandlerFunc(api.GetAlbumsById)

	apirouter.Methods("GET").Path("/notification/create").HandlerFunc(api.GetCreateNotification)
	apirouter.Methods("GET").Path("/notification/delete").HandlerFunc(api.GetDeleteNotification)
}
