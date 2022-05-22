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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/interests/{filter}": {
            "get": {
                "description": "Get interests by performance",
                "tags": [
                    "Interests"
                ],
                "summary": "Get interests",
                "parameters": [
                    {
                        "enum": [
                            "bad",
                            "good",
                            "excellent",
                            "three"
                        ],
                        "type": "string",
                        "description": "Filter",
                        "name": "filter",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/client_models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/client_models.BadResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/client_models.BadResponse"
                        }
                    }
                }
            }
        },
        "/students/{filter}": {
            "get": {
                "description": "Get students by filter\nCurrently only supporting vk group id",
                "tags": [
                    "Students"
                ],
                "summary": "Get students by filter",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filter",
                        "name": "filter",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/client_models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/client_models.BadResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/client_models.BadResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "client_models.BadResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "NOT OK"
                }
            }
        },
        "client_models.Response": {
            "type": "object",
            "properties": {
                "statistics": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "status": {
                    "type": "string",
                    "enum": [
                        "OK"
                    ]
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "euandvkanalysis.herokuapp.com",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "EU and VK Analytics API documentation",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
