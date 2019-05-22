# IMAGE STORE

An Image store for creation and deletion on Albums and images. 

Features: 
1) Create/Delete Image Album using REST API
2) Create/Delete Image in an Album using REST API
3) Get an Image and all images in an Album using REST API
4) Produce Notification Message whenever Create/Delete Image using message broker like Kafka
5) Get the Create/Delete Image Notification Messages using REST API
  

Non-Function Requirements:
1. Documentation using Swagger
2. Performance Metrics
3. Unit testcases


Benchmark Result and Code coverage
```
goos: linux
goarch: amd64
pkg: ImageStore/pkg/apihandler
BenchmarkCreateAlbumHandler-4      	 1000000	      1722 ns/op	     394 B/op	       4 allocs/op
BenchmarkCreateImageHandler-4      	  500000	      2816 ns/op	     668 B/op	       7 allocs/op
BenchmarkDeleteAlbumHandler-4      	 1000000	      1707 ns/op	     438 B/op	       5 allocs/op
BenchmarkDeleteImageHandler-4      	  500000	      2877 ns/op	     712 B/op	       8 allocs/op
BenchmarkGetAlbumsList-4           	  100000	     16172 ns/op	   10638 B/op	      15 allocs/op
BenchmarkGetImages-4               	  100000	     16464 ns/op	   10934 B/op	      18 allocs/op
BenchmarkGetImagesByName-4         	  200000	     15054 ns/op	    3810 B/op	      24 allocs/op
BenchmarkGetCreateNotification-4   	       1	2705000417 ns/op	 1116352 B/op	     575 allocs/op
BenchmarkGetDeleteNotification-4   	       1	2966355847 ns/op	 1116160 B/op	     597 allocs/op
PASS
coverage: 89.9% of statements
ok  	ImageStore/pkg/apihandler	32.201s
```
```
pkg: ImageStore/pkg/messaging
BenchmarkReadMessage-4    	       1	2953916904 ns/op
BenchmarkWriteMessage-4   	    2000	   1032789 ns/op
PASS
coverage: 72.7% of statements
ok  	ImageStore/pkg/messaging	11.286s
```


make unit-test
```
go test -cover ./...
ok  	ImageStore/cmd	0.022s	coverage: 40.9% of statements [no tests to run]
ok  	ImageStore/pkg/apihandler	17.150s	coverage: 86.9% of statements
ok  	ImageStore/pkg/apiroutes	0.005s	coverage: 62.5% of statements
ok  	ImageStore/pkg/messaging	2.239s	coverage: 69.7% of statements
ok  	ImageStore/pkg/utils	0.004s	coverage: 0.0% of statements [no tests to run
```
