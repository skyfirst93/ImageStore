package apiroutes

import (
	"ImageStore/pkg/apihandler"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

//RunApi is used to start listioning on the endpoint and
// Call RunAPIOnRouter function
func RunApi(waitgroup *sync.WaitGroup, endpoint string) {
	r := mux.NewRouter()
	RunAPIOnRouter(r)
	err := http.ListenAndServe(endpoint, r)
	if err != nil {
		fmt.Println("Error-in-REST-Server", err)
	}
	waitgroup.Done()
}

//RunAPIOnRouter is used for setting the routes for the API's
func RunAPIOnRouter(r *mux.Router) {

	apirouter := r.PathPrefix("/api/store").Subrouter()

	apirouter.Methods("POST").Path("/create/album/{albumname}").HandlerFunc(apihandler.CreateAlbumHandler)
	apirouter.Methods("DELETE").Path("/delete/album/{albumname}").HandlerFunc(apihandler.DeleteAlbumHandler)

	apirouter.Methods("POST").Path("/create/image/{albumname}/{imagename}").HandlerFunc(apihandler.CreateImageHandler)
	apirouter.Methods("DELETE").Path("/delete/image/{albumname}/{imagename}").HandlerFunc(apihandler.DeleteImageHandler)

	apirouter.Methods("GET").Path("/images/{albumname}").HandlerFunc(apihandler.GetImages)
	apirouter.Methods("GET").Path("/images/{albumname}/{imagename}").HandlerFunc(apihandler.GetImagesByName)
	apirouter.Methods("GET").Path("/albums").HandlerFunc(apihandler.GetAlbumsList)

	apirouter.Methods("GET").Path("/notification/create").HandlerFunc(apihandler.GetCreateNotification)
	apirouter.Methods("GET").Path("/notification/delete").HandlerFunc(apihandler.GetDeleteNotification)

	apirouter.Methods("GET").Path("/swagger.json").HandlerFunc(apihandler.Swagger)

}
