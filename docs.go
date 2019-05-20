
package 
//This file is generated automatically. Do not try to edit it manually.

var ResourceListingJson = `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://127.0.0.1:8081/api/store",
    "apis": [
        {
            "path": "/api/store",
            "description": "Create album handler creates the album by name"
        },
        {
            "path": "/create/album/",
            "description": "Create Album API"
        },
        {
            "path": "/create/image/",
            "description": "Create Image API"
        },
        {
            "path": "/delete/image/",
            "description": "Delete Image API"
        },
        {
            "path": "/albums",
            "description": "Get Albums list API"
        },
        {
            "path": "/images/",
            "description": "GetImagesByName API"
        },
        {
            "path": "/api/store/notification/create",
            "description": "Get Create Notification API"
        },
        {
            "path": "/notification/delete",
            "description": "GetDeleteNotification API"
        }
    ],
    "info": {
        "title": "IMAGE STORE SWAGGER API",
        "description": "Image store for albums and images"
    }
}`
var ApiDescriptionsJson = map[string]string{"api/store":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://127.0.0.1:8081/api/store",
    "resourcePath": "/api/store",
    "apis": [
        {
            "path": "/api/store/create/album/{albumname}",
            "description": "Create album handler creates the album by name",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "CreateAlbumHandler",
                    "type": "ImageStore.pkg.apihandler..Response",
                    "items": {},
                    "summary": "Create album handler creates the album by name",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "albumname",
                            "description": "Album Name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "ImageStore.pkg.apihandler..Response"
                        },
                        {
                            "code": 209,
                            "message": "Album Name already present",
                            "responseType": "object",
                            "responseModel": "ImageStore.pkg.apihandler..Response"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/api/store/create/image/{albumname}/{imagename}",
            "description": "Create Image handler creates the image by name",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "Create Image Handler",
                    "type": "ImageStore.pkg.apihandler..Response",
                    "items": {},
                    "summary": "Create Image handler creates the image by name",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "albumname",
                            "description": "Album Name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "imagename",
                            "description": "Image Name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "ImageStore.pkg.apihandler..Response"
                        },
                        {
                            "code": 404,
                            "message": "Album Name already present",
                            "responseType": "object",
                            "responseModel": "ImageStore.pkg.apihandler..Response"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/api/store/delete/album/{albumname}",
            "description": "DeleteAlbumHandler is handler function for deleting an album",
            "operations": [
                {
                    "httpMethod": "DELETE",
                    "nickname": "Delete Album Handler",
                    "type": "ImageStore.pkg.apihandler..Response",
                    "items": {},
                    "summary": "DeleteAlbumHandler is handler function for deleting an album",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "albumname",
                            "description": "Album Name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "ImageStore.pkg.apihandler..Response"
                        },
                        {
                            "code": 404,
                            "message": "Album Name already present",
                            "responseType": "object",
                            "responseModel": "ImageStore.pkg.apihandler..Response"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/api/store/delete/image/{albumname}/{imagename}",
            "description": "Create Image handler creates the image by name",
            "operations": [
                {
                    "httpMethod": "DELETE",
                    "nickname": "Delete Image Handler",
                    "type": "ImageStore.pkg.apihandler..Response",
                    "items": {},
                    "summary": "Create Image handler creates the image by name",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "albumname",
                            "description": "Album Name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "imagename",
                            "description": "Image Name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "ImageStore.pkg.apihandler..Response"
                        },
                        {
                            "code": 404,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "ImageStore.pkg.apihandler..Response"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/api/store/albums",
            "description": "GetAlbumsList is handler function for getting list of albums",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "Get Albums List Handler",
                    "type": "ImageStore.pkg.apihandler..MultiValuesResponse",
                    "items": {},
                    "summary": "GetAlbumsList is handler function for getting list of albums",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "ImageStore.pkg.apihandler..MultiValuesResponse"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/api/store/images/{albumname}/",
            "description": "GetImages is handler function for getting list of image in an album",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "Get Images List Handler",
                    "type": "ImageStore.pkg.apihandler..Response",
                    "items": {},
                    "summary": "GetImages is handler function for getting list of image in an album",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "albumname",
                            "description": "Album Name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "ImageStore.pkg.apihandler..Response"
                        },
                        {
                            "code": 404,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "ImageStore.pkg.apihandler..Response"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/api/store/images/{albumname}/{imagename}",
            "description": "GetImagesByName is handler function for getting an image in an album",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "GetImagesByName Handler",
                    "type": "ImageStore.pkg.apihandler..Response",
                    "items": {},
                    "summary": "GetImagesByName is handler function for getting an image in an album",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "albumname",
                            "description": "Album Name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "imagename",
                            "description": "Image Name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "ImageStore.pkg.apihandler..Response"
                        },
                        {
                            "code": 404,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "ImageStore.pkg.apihandler..Response"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/api/store/notification/create",
            "description": "Get Create Notification is handler function for getting the list of notification",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "Get Create Notification Handler",
                    "type": "ImageStore.pkg.apihandler..MultiValuesResponse",
                    "items": {},
                    "summary": "Get Create Notification is handler function for getting the list of notification",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Create Notifications",
                            "responseType": "object",
                            "responseModel": "ImageStore.pkg.apihandler..MultiValuesResponse"
                        },
                        {
                            "code": 204,
                            "message": "No more Create notification",
                            "responseType": "object",
                            "responseModel": "ImageStore.pkg.apihandler..Response"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/api/store/notification/delete",
            "description": "GetDeleteNotification is handler function for getting the list of notification of Images deleted",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "Get Delete Notification Handler",
                    "type": "ImageStore.pkg.apihandler..MultiValuesResponse",
                    "items": {},
                    "summary": "GetDeleteNotification is handler function for getting the list of notification of Images deleted",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Delete Notifications",
                            "responseType": "object",
                            "responseModel": "ImageStore.pkg.apihandler..MultiValuesResponse"
                        },
                        {
                            "code": 204,
                            "message": "No more Delete notification",
                            "responseType": "object",
                            "responseModel": "ImageStore.pkg.apihandler..Response"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "ImageStore.pkg.apihandler..MultiValuesResponse": {
            "id": "ImageStore.pkg.apihandler..MultiValuesResponse",
            "properties": {
                "HttpStatus": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "values": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "type": "string"
                    },
                    "format": ""
                }
            }
        },
        "ImageStore.pkg.apihandler..Response": {
            "id": "ImageStore.pkg.apihandler..Response",
            "properties": {
                "HttpStatus": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "message": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`,}
