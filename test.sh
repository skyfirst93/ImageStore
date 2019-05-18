#Create album
echo "Creating album test"
curl -X POST http://localhost:8081/api/store/create/album/akash
curl -X POST http://localhost:8081/api/store/create/album/akash1
curl -X POST http://localhost:8081/api/store/create/album/akash1


#create image
echo "Creating images test"
curl -X POST http://localhost:8081/api/store/create/image/akash/image1
curl -X POST http://localhost:8081/api/store/create/image/akash/image2
curl -X POST http://localhost:8081/api/store/create/image/akash1/image1
curl -X POST http://localhost:8081/api/store/create/image/akash1/image2
curl -X POST http://localhost:8081/api/store/create/image/akash1/image2
curl -X POST http://localhost:8081/api/store/create/image/akash2/image2


#get images
echo "Get images test"
curl -X GET http://localhost:8081/api/store/images/akash
curl -X GET http://localhost:8081/api/store/images/akash1
curl -X GET http://localhost:8081/api/store/images/akash2

#get albums
echo "Getting album list test"
curl -X GET http://localhost:8081/api/store/albums

#get image by name 
echo "get individual image test"
curl -X GET http://localhost:8081/api/store/images/akash/image1
curl -X GET http://localhost:8081/api/store/images/akash/image2
curl -X GET http://localhost:8081/api/store/images/akash1/image1
curl -X GET http://localhost:8081/api/store/images/akash1/image2
curl -X GET http://localhost:8081/api/store/images/akash/image3

#delete album
echo "Delete album test"
curl -X DELETE http://localhost:8081/api/store/delete/album/akash
curl -X DELETE http://localhost:8081/api/store/delete/album/akash1
curl -X DELETE http://localhost:8081/api/store/delete/album/akash2

#delete images
echo "Delete image test"
curl -X DELETE  http://localhost:8081/api/store/delete/image/akash/image1
curl -X DELETE  http://localhost:8081/api/store/delete/image/akash/image1
curl -X DELETE  http://localhost:8081/api/store/delete/image/akash/image2
curl -X DELETE  http://localhost:8081/api/store/delete/image/akash1/image1
curl -X DELETE  http://localhost:8081/api/store/delete/image/akash1/image2
curl -X DELETE  http://localhost:8081/api/store/delete/image/akash2/image1

#delete album
echo "Delete album test"
curl -X DELETE http://localhost:8081/api/store/delete/album/akash
curl -X DELETE http://localhost:8081/api/store/delete/album/akash1
curl -X DELETE http://localhost:8081/api/store/delete/album/akash2

