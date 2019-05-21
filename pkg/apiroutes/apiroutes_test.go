package apiroutes

import (
	"testing"

	"github.com/gorilla/mux"
)

//TestRunAPIOnRouter is used to test for setting the routes for the API's
func TestRunAPIOnRouter(t *testing.T) {
	r := mux.NewRouter()
	RunAPIOnRouter(r)
}
