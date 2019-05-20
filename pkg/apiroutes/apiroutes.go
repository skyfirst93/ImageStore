package apiroutes

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

//RunAPI is used to start listioning on the endpoint and
// Call RunAPIOnRouter function
func RunAPI(waitgroup *sync.WaitGroup, endpoint string) {
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

	//apirouter := r.PathPrefix("/api/store").Subrouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8081/docs/swagger.json"), //The url pointing to API definition"
	))
	/*
		// Swagger operation POST /create/album/{albumname}
		//
		// add here the decsription
		// ---
		// @Summary CREATE album
		// @ID CreateAlbumHandler
		// @Produce application/json
		// @Param albumname path string true "The Name for the album that is to be created"
		// @responses:
		//   200:
		//     description: successful operation
		//   209:
		//     description: Data Already present
		//   400:
		//     description: Invalid name supplied
		// @Router /create/album/{albumname} [POST]
		apirouter.Methods("POST").Path("/create/album/{albumname}").HandlerFunc(apihandler.CreateAlbumHandler)

		// swagger:operation DELETE /delete/album/{albumname}
		//
		// add here the decsription
		// ---
		// summary: DELETE album
		// operationId: DeleteAlbumHandler
		// produces:
		// - application/json
		// parameters:
		// - name: albumname
		//   in: path
		//   description: The Name for the album that is to be deleted
		//   required: true
		//   type: string
		// responses:
		//   200:
		//     description: successful operation
		//   404:
		//     description: No data present
		//   400:
		//     description: Invalid name supplied
		apirouter.Methods("DELETE").Path("/delete/album/{albumname}").HandlerFunc(apihandler.DeleteAlbumHandler)

		// swagger:operation POST /create/image/{albumname}/{imagename}
		//
		// add here the decsription
		// ---
		// summary: Create image
		// operationId: CreateImageHandler
		// produces:
		// - application/json
		// parameters:
		// - name: albumname
		//   in: path
		//   description: The Name for the album that containes the image
		//   required: true
		//   type: string
		// - name: imagename
		//   in: path
		//   description: The name for the image that needs to be created
		//   required: true
		//   type: string
		// responses:
		//   200:
		//     description: successful operation
		//   404:
		//     description: No data present
		//   400:
		//     description: Invalid name supplied
		apirouter.Methods("POST").Path("/create/image/{albumname}/{imagename}").HandlerFunc(apihandler.CreateImageHandler)

		// swagger:operation DELETE /delete/image/{albumname}/{imagename}
		//
		// add here the decsription
		// ---
		// summary: Delete the image by image
		// operationId: DeleteImageHandler
		// produces:
		// - application/json
		// parameters:
		// - name: albumname
		//   in: path
		//   description: The Name for the album that containes the image
		//   required: true
		//   type: string
		// - name: imagename
		//   in: path
		//   description: The name for the image that needs to be deleted
		//   required: true
		//   type: string
		// responses:
		//   200:
		//     description: successful operation
		//   404:
		//     description: No data present
		//   400:
		//     description: Invalid name supplied
		apirouter.Methods("DELETE").Path("/delete/image/{albumname}/{imagename}").HandlerFunc(apihandler.DeleteImageHandler)

		// swagger:operation GET /images/{albumname}
		//
		// add here the decsription
		// ---
		// summary: Get the image by image
		// operationId: GetImages
		// produces:
		// - application/json
		// parameters:
		// - name: albumname
		//   in: path
		//   description: The Name for the album that containes the image
		//   required: true
		//   type: string
		// responses:
		//   200:
		//     description: successful operation
		//   404:
		//     description: No data present
		//   400:
		//     description: Invalid name supplied
		apirouter.Methods("GET").Path("/images/{albumname}").HandlerFunc(apihandler.GetImages)

		// swagger:operation GET /images/{albumname}/{imagename}
		//
		// add here the decsription
		// ---
		// summary: Get the image by image
		// operationId: GetImagesByName
		// produces:
		// - application/json
		// parameters:
		// - name: albumname
		//   in: path
		//   description: The Name for the album that containes the image
		//   required: true
		//   type: string
		// - name: imagename
		//   in: path
		//   description: The name for the image that needs to be fetched.
		//   required: true
		//   type: string
		// responses:
		//   200:
		//     description: successful operation
		//   404:
		//     description: Data not found
		//   400:
		//     description: Invalid name supplied

		apirouter.Methods("GET").Path("/images/{albumname}/{imagename}").HandlerFunc(apihandler.GetImagesByName)
		// swagger:operation GET /albums
		//
		// add here the decsription
		// ---
		// summary: Get the list of albums present
		// operationId: GetAlbumslist
		// produces:
		// - application/json
		// responses:
		//   200:
		//     description: successful operation
		//   404:
		//     description: No Data present
		apirouter.Methods("GET").Path("/albums").HandlerFunc(apihandler.GetAlbumsList)
		// swagger:operation GET /notification/create
		//
		// add here the decsription
		// ---
		// summary: Get Creation notification
		// operationId: GetCreateNotiffication
		// produces:
		// - application/json
		// responses:
		//   200:
		//     description: successful operation
		//   204:
		//     description: No more Notification
		apirouter.Methods("GET").Path("/notification/create").HandlerFunc(apihandler.GetCreateNotification)
		// swagger:operation GET /notification/delete
		//
		// add here the decsription
		// ---
		// summary: Get deletion notification
		// operationId: getDeleteNotiffication
		// produces:
		// - application/json
		// responses:
		//   200:
		//     description: successful operation
		//   204:
		//     description: No more Notification
		apirouter.Methods("GET").Path("/notification/delete").HandlerFunc(apihandler.GetDeleteNotification)
	*/
}
