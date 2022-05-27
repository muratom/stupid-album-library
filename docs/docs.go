// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "muratom",
            "email": "muratom73@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/albums": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves all albums from DB",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Album"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Adds an album from JSON",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Adds an album from JSON",
                "parameters": [
                    {
                        "description": "Album info",
                        "name": "album",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AlbumRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "created",
                        "schema": {
                            "$ref": "#/definitions/models.Album"
                        }
                    },
                    "500": {
                        "description": "fail"
                    }
                }
            }
        },
        "/albums/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves album based on given ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Album ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Album"
                        }
                    },
                    "500": {
                        "description": "Fail"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Album": {
            "type": "object",
            "properties": {
                "artist": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.AlbumRequest": {
            "type": "object",
            "properties": {
                "artist": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "My first Swagger API",
	Description:      "My first Swagger API for Golang/Gin",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}