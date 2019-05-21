package apihandler

import (
	"ImageStore/pkg/messaging"
	"ImageStore/pkg/utils"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

var url = "whocares"

// init function initializes Message Bus
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
	}

	messaging.InitProducer(utils.MessageQueueAddr)
	messaging.InitConsumer(utils.MessageQueueAddr, "group")
}

//createAlbum is function used for creating test albums
func createAlbum(albumName string) {
	req, _ := http.NewRequest("POST", url, nil)
	req = mux.SetURLVars(req, map[string]string{"albumname": albumName})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateAlbumHandler)
	handler.ServeHTTP(rr, req)
}

//createImage is function used for creating test image in test albums
func createImage(albumName, imageName string) {
	req, _ := http.NewRequest("POST", url, nil)
	req = mux.SetURLVars(req, map[string]string{"albumname": albumName, "imagename": imageName})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateImageHandler)
	handler.ServeHTTP(rr, req)
}

//deleteAlbum is function used for deleting test albums
func deleteAlbum(albumName string) {
	req, _ := http.NewRequest("POST", url, nil)
	req = mux.SetURLVars(req, map[string]string{"albumname": albumName})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteAlbumHandler)
	handler.ServeHTTP(rr, req)
}

//deleteImage is function used for creating test image in test albums
func deleteImage(albumName, imageName string) {
	req, _ := http.NewRequest("POST", url, nil)
	req = mux.SetURLVars(req, map[string]string{"albumname": albumName, "imagename": imageName})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteImageHandler)
	handler.ServeHTTP(rr, req)
}

//TestCreateAlbumHandler for testing album creation
func TestCreateAlbumHandler(t *testing.T) {
	createAlbum("Test1")
	testCases := []struct {
		albumname string
		expected  int
	}{
		{"Test1", http.StatusConflict},
	}
	for _, tc := range testCases {
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			t.Fatal(err)
		}
		req = mux.SetURLVars(req, map[string]string{"albumname": tc.albumname})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateAlbumHandler)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != tc.expected {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, tc.expected)
		}
	}
	deleteAlbum("test1")
}

//TestCreateImageHandler for testing image creation
func TestCreateImageHandler(t *testing.T) {
	createAlbum("Test1")
	createImage("Test1", "image.png")
	testCases := []struct {
		albumname string
		imagename string
		expected  int
	}{
		{"Test1", "image.png", http.StatusConflict},
		{"Test20", "image.png", http.StatusNotFound},
	}
	for _, tc := range testCases {
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			t.Fatal(err)
		}
		req = mux.SetURLVars(req, map[string]string{"albumname": tc.albumname, "imagename": tc.imagename})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateImageHandler)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != tc.expected {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, tc.expected)
		}
	}
	deleteImage("Test1", "image.png")
	deleteAlbum("Test1")
}

//TestDeleteAlbumHandler is handler function for testing deleting an album
//and return suucess or failure
func TestDeleteAlbumHandler(t *testing.T) {
	createAlbum("Test2")
	createAlbum("Test3")
	createImage("Test3", "image.png")
	testCases := []struct {
		albumname string
		expected  int
	}{
		{"Test2", http.StatusOK},
		{"Test2", http.StatusNotFound},
		{"Test3", http.StatusInternalServerError},
	}
	for _, tc := range testCases {
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			t.Fatal(err)
		}
		req = mux.SetURLVars(req, map[string]string{"albumname": tc.albumname})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(DeleteAlbumHandler)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != tc.expected {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, tc.expected)
		}
	}
	deleteImage("Test3", "image.png")
	deleteAlbum("Test3")
}

//TestDeleteImageHandler is handler function for testing deleting an image
//and return suucess or failure
func TestDeleteImageHandler(t *testing.T) {
	createAlbum("Test1")
	createImage("Test1", "image.png")
	testCases := []struct {
		albumname string
		imagename string
		expected  int
	}{
		{"Test1", "image.png", http.StatusOK},
		{"Test1", "image.png", http.StatusNotFound},
		{"Test20", "image.png", http.StatusNotFound},
	}
	for _, tc := range testCases {
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			t.Fatal(err)
		}
		req = mux.SetURLVars(req, map[string]string{"albumname": tc.albumname, "imagename": tc.imagename})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(DeleteImageHandler)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != tc.expected {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, tc.expected)
		}
	}
	deleteAlbum("Test1")
}

//GetAlbumsList is handler function for testing getting list of albums
func TestGetAlbumsList(t *testing.T) {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAlbumsList)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//TestGetImages is handler function for testing getting list of image
//and returning the list of images
func TestGetImages(t *testing.T) {
	createAlbum("Test1")
	createImage("Test1", "image.png")
	testCases := []struct {
		albumname string
		expected  int
	}{
		{"Test1", http.StatusOK},
		{"Test11", http.StatusNotFound},
	}
	for _, tc := range testCases {
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			t.Fatal(err)
		}
		req = mux.SetURLVars(req, map[string]string{"albumname": tc.albumname})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetImages)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != tc.expected {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, tc.expected)
		}
	}
	deleteImage("Test1", "image.png")
	deleteAlbum("Test1")
}

//GetImagesByName is handler function for testing  getting an image
//and returning the image
func TestGetImagesByName(t *testing.T) {
	createAlbum("Test1")
	createImage("Test1", "image.png")
	testCases := []struct {
		albumname string
		imagename string
		expected  int
	}{
		//		{"Test1", "image.png", http.StatusOK},
		{"Test1", "image12.png", http.StatusNotFound},
		{"Test20", "image.png", http.StatusNotFound},
	}
	for _, tc := range testCases {
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			t.Fatal(err)
		}
		req = mux.SetURLVars(req, map[string]string{"albumname": tc.albumname, "imagename": tc.imagename})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetImagesByName)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != tc.expected {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, tc.expected)
		}
	}
	deleteImage("Test1", "image.png")
	deleteAlbum("Test1")
}

//TestGetCreateNotification is handler function for testing getting the list of notification
//of Images created
func TestGetCreateNotification(t *testing.T) {
	createAlbum("Test1")
	createImage("Test1", "image.png")
	testCases := []struct {
		expected int
	}{
		{http.StatusOK},
		{http.StatusNoContent},
	}
	for _, tc := range testCases {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetCreateNotification)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != tc.expected {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, tc.expected)
		}
	}
	deleteImage("Test1", "image.png")
	deleteAlbum("Test1")
}

//TestGetDeleteNotification is handler function for testing getting the list of notification of
//Images deleted
func TestGetDeleteNotification(t *testing.T) {
	createAlbum("Test1")
	createImage("Test1", "image.png")
	deleteImage("Test1", "image.png")
	deleteAlbum("Test1")
	testCases := []struct {
		expected int
	}{
		{http.StatusOK},
		{http.StatusNoContent},
	}
	for _, tc := range testCases {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetDeleteNotification)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != tc.expected {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, tc.expected)
		}
	}
}
