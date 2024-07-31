// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/message/get": {
            "get": {
                "description": "Handles request to get messages by filter.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "messages"
                ],
                "summary": "Get messages",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Message ID",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Message content",
                        "name": "message",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Status",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Creation timestamp (DateTime format)",
                        "name": "created_at",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Update timestamp (DateTime format)",
                        "name": "updated_at",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully got messages",
                        "schema": {
                            "$ref": "#/definitions/models.FilterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/message/new": {
            "post": {
                "description": "Handles request to create a new message and returns the message information in JSON.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "messages"
                ],
                "summary": "Creating a new message",
                "parameters": [
                    {
                        "description": "Passport",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Created message",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreateRequest": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.FilterResponse": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Message"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "models.Message": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/models.Status"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Status": {
            "type": "string",
            "enum": [
                "COMPLETED",
                "PROCESSING",
                "FAILED"
            ],
            "x-enum-varnames": [
                "StatusCompleted",
                "StatusProcessing",
                "StatusFailed"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "89.169.136.165:8800",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Messages server API",
	Description:      "This is messages_server server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
