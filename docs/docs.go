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
        "/common": {
            "get": {
                "description": "Pinging for server and app",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Common"
                ],
                "summary": "Router for ping",
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/pwa": {
            "post": {
                "description": "Создает новый PWA на основе переданных данных",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "PWA"
                ],
                "summary": "Создание нового PWA",
                "parameters": [
                    {
                        "description": "Данные для создания PWA",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dao.PwaCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное создание PWA"
                    },
                    "400": {
                        "description": "Неверный формат запроса"
                    },
                    "500": {
                        "description": "Ошибка сервера при создании PWA"
                    }
                }
            }
        }
    },
    "definitions": {
        "dao.PwaCreateRequest": {
            "type": "object",
            "properties": {
                "iconUrl": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/enums.TypeCampaign"
                }
            }
        },
        "enums.TypeCampaign": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-varnames": [
                "GoogleBlue",
                "GoogleGreen",
                "AppleStore"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Gin Swagger Landing Constructor",
	Description:      "App for landing constructor.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
