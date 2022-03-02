// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/container-by-receipt": {
            "post": {
                "description": "Receipt By Container",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "All example",
                "parameters": [
                    {
                        "description": "ReceiptByContainer",
                        "name": "ReceiptByContainer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ReceiptNumber"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.ContainerByReceipt"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/count": {
            "get": {
                "description": "Count Receipt Sea",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "All example",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token. Format: Bearer access_token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.CountReceiptSea"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/detail": {
            "get": {
                "description": "Receipt Detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "All example",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ReceiptId",
                        "name": "receiptId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "ContainerId",
                        "name": "containerId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.ReceiptListResult"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/list": {
            "get": {
                "description": "Receipt List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "All example",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token. Format: Bearer access_token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Pages",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Status",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.ReceiptListResult"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/receipt/list": {
            "get": {
                "description": "Receipt List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "All example",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token. Format: Bearer access_token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Pages",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "receiptType",
                        "name": "receiptType",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.ReceiptListByTypeResult"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/tracking": {
            "get": {
                "description": "Receipt Tracking",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Tracking example",
                "parameters": [
                    {
                        "type": "string",
                        "description": "receiptNumber",
                        "name": "receiptNumber",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "markingCode",
                        "name": "markingCode",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.ReceiptListResult"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ContainerByReceipt": {
            "type": "object",
            "properties": {
                "containerID": {
                    "type": "integer"
                },
                "receiptID": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "dto.CountReceiptSea": {
            "type": "object",
            "properties": {
                "arrivedSoon": {
                    "type": "integer"
                },
                "delay": {
                    "type": "integer"
                },
                "onTheWay": {
                    "type": "integer"
                }
            }
        },
        "dto.ReceiptListByTypeResult": {
            "type": "object",
            "properties": {
                "content": {},
                "pagination": {}
            }
        },
        "dto.ReceiptListResult": {
            "type": "object",
            "properties": {
                "content": {},
                "pagination": {}
            }
        },
        "dto.ReceiptNumber": {
            "type": "object",
            "properties": {
                "receiptSeaNumber": {
                    "type": "string"
                }
            }
        },
        "helper.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "errors": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
