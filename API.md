
# IMAGE STORE SWAGGER API
Image store for albums and images

Table of Contents

1. [Get Albums list API](#albums)
1. [Create album handler creates the album by name](#api/store)
1. [Get Create Notification API](#api/store/notification/create)
1. [Create Album API](#create/album/)
1. [Create Image API](#create/image/)
1. [Delete Image API](#delete/image/)
1. [GetImagesByName API](#images/)
1. [GetDeleteNotification API](#notification/delete)

<a name="api/store"></a>

## api/store

| Specification | Value |
|-----|-----|
| Resource Path | /api/store |
| API Version | 1.0.0 |
| BasePath for the API | http://127.0.0.1:8081/api/store |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /api/store/create/album/\{albumname\} | [POST](#CreateAlbumHandler) | Create album handler creates the album by name |
| /api/store/create/image/\{albumname\}/\{imagename\} | [POST](#Create Image Handler) | Create Image handler creates the image by name |
| /api/store/delete/album/\{albumname\} | [DELETE](#Delete Album Handler) | DeleteAlbumHandler is handler function for deleting an album |
| /api/store/delete/image/\{albumname\}/\{imagename\} | [DELETE](#Delete Image Handler) | Create Image handler creates the image by name |
| /api/store/albums | [GET](#Get Albums List Handler) | GetAlbumsList is handler function for getting list of albums |
| /api/store/images/\{albumname\}/ | [GET](#Get Images List Handler) | GetImages is handler function for getting list of image in an album |
| /api/store/images/\{albumname\}/\{imagename\} | [GET](#GetImagesByName Handler) | GetImagesByName is handler function for getting an image in an album |
| /api/store/notification/create | [GET](#Get Create Notification Handler) | Get Create Notification is handler function for getting the list of notification |
| /api/store/notification/delete | [GET](#Get Delete Notification Handler) | GetDeleteNotification is handler function for getting the list of notification of Images deleted |



<a name="CreateAlbumHandler"></a>

#### API: /api/store/create/album/\{albumname\} (POST)


Create album handler creates the album by name



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| albumname | path | string | Album Name | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Response](#ImageStore.pkg.apihandler..Response) |  |
| 209 | object | [Response](#ImageStore.pkg.apihandler..Response) | Album Name already present |


<a name="Create Image Handler"></a>

#### API: /api/store/create/image/\{albumname\}/\{imagename\} (POST)


Create Image handler creates the image by name



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| albumname | path | string | Album Name | Yes |
| imagename | path | string | Image Name | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Response](#ImageStore.pkg.apihandler..Response) |  |
| 404 | object | [Response](#ImageStore.pkg.apihandler..Response) | Album Name already present |


<a name="Delete Album Handler"></a>

#### API: /api/store/delete/album/\{albumname\} (DELETE)


DeleteAlbumHandler is handler function for deleting an album



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| albumname | path | string | Album Name | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Response](#ImageStore.pkg.apihandler..Response) |  |
| 404 | object | [Response](#ImageStore.pkg.apihandler..Response) | Album Name already present |


<a name="Delete Image Handler"></a>

#### API: /api/store/delete/image/\{albumname\}/\{imagename\} (DELETE)


Create Image handler creates the image by name



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| albumname | path | string | Album Name | Yes |
| imagename | path | string | Image Name | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Response](#ImageStore.pkg.apihandler..Response) |  |
| 404 | object | [Response](#ImageStore.pkg.apihandler..Response) |  |


<a name="Get Albums List Handler"></a>

#### API: /api/store/albums (GET)


GetAlbumsList is handler function for getting list of albums



| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [MultiValuesResponse](#ImageStore.pkg.apihandler..MultiValuesResponse) |  |


<a name="Get Images List Handler"></a>

#### API: /api/store/images/\{albumname\}/ (GET)


GetImages is handler function for getting list of image in an album



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| albumname | path | string | Album Name | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Response](#ImageStore.pkg.apihandler..Response) |  |
| 404 | object | [Response](#ImageStore.pkg.apihandler..Response) |  |


<a name="GetImagesByName Handler"></a>

#### API: /api/store/images/\{albumname\}/\{imagename\} (GET)


GetImagesByName is handler function for getting an image in an album



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| albumname | path | string | Album Name | Yes |
| imagename | path | string | Image Name | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Response](#ImageStore.pkg.apihandler..Response) |  |
| 404 | object | [Response](#ImageStore.pkg.apihandler..Response) |  |


<a name="Get Create Notification Handler"></a>

#### API: /api/store/notification/create (GET)


Get Create Notification is handler function for getting the list of notification



| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [MultiValuesResponse](#ImageStore.pkg.apihandler..MultiValuesResponse) | Create Notifications |
| 204 | object | [Response](#ImageStore.pkg.apihandler..Response) | No more Create notification |


<a name="Get Delete Notification Handler"></a>

#### API: /api/store/notification/delete (GET)


GetDeleteNotification is handler function for getting the list of notification of Images deleted



| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [MultiValuesResponse](#ImageStore.pkg.apihandler..MultiValuesResponse) | Delete Notifications |
| 204 | object | [Response](#ImageStore.pkg.apihandler..Response) | No more Delete notification |




### Models

<a name="ImageStore.pkg.apihandler..MultiValuesResponse"></a>

#### MultiValuesResponse

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| HttpStatus | int |  |
| values | array |  |

<a name="ImageStore.pkg.apihandler..Response"></a>

#### Response

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| HttpStatus | int |  |
| message | string |  |


