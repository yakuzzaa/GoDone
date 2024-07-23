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
        "/api/items/{list_id}": {
            "get": {
                "description": "Get List Items",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Get List Items",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List ID",
                        "name": "list_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.ListItemResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create item for to-do list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Create Item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List ID",
                        "name": "list_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "CreateItemRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/serializer.CreateItemRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.CreateItemResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/items/{list_id}/{item_id}": {
            "get": {
                "description": "Get List Items",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Get Item by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List ID",
                        "name": "list_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "item_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.GetItemResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Get List Items",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Update Item by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List ID",
                        "name": "list_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "item_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "UpdateItemRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/serializer.UpdateItemRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.UpdateItemResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Item by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Delete Item by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List ID",
                        "name": "list_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "item_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.UpdateItemResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/lists": {
            "get": {
                "description": "Get user to-do lists",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "list"
                ],
                "summary": "Get Lists",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.ListsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create user to-do list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "list"
                ],
                "summary": "Create List",
                "parameters": [
                    {
                        "description": "CreateListRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/serializer.CreateListRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.CreateListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/lists/{id}": {
            "get": {
                "description": "Get user to-do list by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "list"
                ],
                "summary": "Get List by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.CreateListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update user to-do list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "list"
                ],
                "summary": "Update List",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "UpdateListRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/serializer.UpdateListRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.UpdateListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete user to-do list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "list"
                ],
                "summary": "Delete List",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.DeleteListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "Sign in a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User sign-in",
                "parameters": [
                    {
                        "description": "SignInRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/serializer.SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.SignInResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "Sign up a user and get token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User sign-up",
                "parameters": [
                    {
                        "description": "SignInRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/serializer.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.SignUpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serializer.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "serializer.CreateItemRequest": {
            "type": "object",
            "properties": {
                "info": {
                    "$ref": "#/definitions/serializer.ItemInfo"
                }
            }
        },
        "serializer.CreateItemResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "serializer.CreateListRequest": {
            "type": "object",
            "properties": {
                "info": {
                    "$ref": "#/definitions/serializer.ListInfo"
                }
            }
        },
        "serializer.CreateListResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "serializer.DeleteListResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "serializer.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "serializer.GetItemResponse": {
            "type": "object",
            "properties": {
                "item": {
                    "$ref": "#/definitions/serializer.Item"
                }
            }
        },
        "serializer.Item": {
            "type": "object",
            "properties": {
                "create_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "info": {
                    "$ref": "#/definitions/serializer.ItemInfo"
                },
                "update_at": {
                    "type": "string"
                }
            }
        },
        "serializer.ItemInfo": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "done": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "serializer.List": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "info": {
                    "$ref": "#/definitions/serializer.ListInfo"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "serializer.ListInfo": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "serializer.ListItemResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/serializer.Item"
                    }
                }
            }
        },
        "serializer.ListsResponse": {
            "type": "object",
            "properties": {
                "lists": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/serializer.List"
                    }
                }
            }
        },
        "serializer.SignInInfo": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "serializer.SignInRequest": {
            "type": "object",
            "properties": {
                "info": {
                    "$ref": "#/definitions/serializer.SignInInfo"
                }
            }
        },
        "serializer.SignInResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "serializer.SignUpInfo": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "serializer.SignUpRequest": {
            "type": "object",
            "properties": {
                "info": {
                    "$ref": "#/definitions/serializer.SignUpInfo"
                }
            }
        },
        "serializer.SignUpResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "serializer.UpdateItemInfo": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "done": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "serializer.UpdateItemRequest": {
            "type": "object",
            "properties": {
                "info": {
                    "$ref": "#/definitions/serializer.UpdateItemInfo"
                }
            }
        },
        "serializer.UpdateItemResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "serializer.UpdateListInfo": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "serializer.UpdateListRequest": {
            "type": "object",
            "properties": {
                "info": {
                    "$ref": "#/definitions/serializer.UpdateListInfo"
                }
            }
        },
        "serializer.UpdateListResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "GoDone API",
	Description:      "Simple To-Do list backend",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
