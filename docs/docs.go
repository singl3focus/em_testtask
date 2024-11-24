// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/song/add": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Add song",
                "parameters": [
                    {
                        "description": "Song",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Song was succesfully added",
                        "schema": {
                            "$ref": "#/definitions/handler.Result"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        },
        "/song/info/get": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Get songs info by matched params",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group Name",
                        "name": "groupName",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Song Title",
                        "name": "songTitle",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Songs info",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.SongInfo"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        },
        "/song/remove": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Remove song by groupName, songTitle",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group Name",
                        "name": "groupName",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Song Title",
                        "name": "songTitle",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Song was successfully removed",
                        "schema": {
                            "$ref": "#/definitions/handler.Result"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        },
        "/song/text/by-verses": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Get song's text by groupName,songTitle with verses pagination",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group Name",
                        "name": "groupName",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Song Title",
                        "name": "songTitle",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Song text",
                        "schema": {
                            "$ref": "#/definitions/handler.Result"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        },
        "/song/update": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Update song by groupName, songTitle",
                "parameters": [
                    {
                        "description": "Old and new song info",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.UpdateSongInfoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Song's info was successfully update",
                        "schema": {
                            "$ref": "#/definitions/handler.Result"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "handler.Result": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.UpdateSongInfoRequest": {
            "type": "object",
            "properties": {
                "newGroupName": {
                    "type": "string"
                },
                "newSongTitle": {
                    "type": "string"
                },
                "oldGroupName": {
                    "type": "string"
                },
                "oldSongTitle": {
                    "type": "string"
                }
            }
        },
        "models.Song": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "verses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Verse"
                    }
                }
            }
        },
        "models.SongInfo": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.Verse": {
            "type": "object",
            "properties": {
                "number": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Songs API",
	Description:      "This is a song library API as a test assignment for the company Effective mobile",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
