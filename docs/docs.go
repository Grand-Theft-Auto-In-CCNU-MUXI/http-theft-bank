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
        "/bank/gate": {
            "put": {
                "description": "站点3，用 put 方法传输 error code",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bank"
                ],
                "summary": "user use error code to hack this gate",
                "parameters": [
                    {
                        "description": "request",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.Request"
                        }
                    },
                    {
                        "type": "string",
                        "description": "代号名",
                        "name": "code",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "通行证",
                        "name": "passport",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/bank/iris_recognition_gate": {
            "get": {
                "description": "站点4，获取游戏场景文本",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bank"
                ],
                "summary": "站点4获取文本",
                "parameters": [
                    {
                        "type": "string",
                        "description": "代号名",
                        "name": "code",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "通行证",
                        "name": "passport",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/muxi/backend/computer/examination": {
            "get": {
                "description": "站点5，获取文本",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "end"
                ],
                "summary": "站点5 获取游戏场景文本",
                "parameters": [
                    {
                        "type": "string",
                        "description": "代号名",
                        "name": "code",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "通行证",
                        "name": "passport",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "站点5，上传文件",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "end"
                ],
                "summary": "站点5 上传文件 ac",
                "parameters": [
                    {
                        "description": "file,这个使用表单！！！",
                        "name": "file",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "primitive"
                        }
                    },
                    {
                        "type": "string",
                        "description": "代号名",
                        "name": "code",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "通行证",
                        "name": "passport",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/organization/code": {
            "get": {
                "description": "站点1，返回token在头部",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "organization"
                ],
                "summary": "send user token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "代号名",
                        "name": "code",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/organization/iris_sample": {
            "get": {
                "description": "站点4，下载图片",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "organization"
                ],
                "summary": "下载图片",
                "parameters": [
                    {
                        "type": "string",
                        "description": "代号名",
                        "name": "code",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "通行证",
                        "name": "passport",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "站点4，上传图片",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bank"
                ],
                "summary": "站点4 上传图片",
                "parameters": [
                    {
                        "description": "file,这个是用表单上传！！！！",
                        "name": "file",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "primitive"
                        }
                    },
                    {
                        "type": "string",
                        "description": "代号名",
                        "name": "code",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "通行证",
                        "name": "passport",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/organization/secret_key": {
            "get": {
                "description": "站点2，返回加密的密钥和错误代码片段，在 body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "organization"
                ],
                "summary": "send user secretKey and error function",
                "parameters": [
                    {
                        "type": "string",
                        "description": "代号名",
                        "name": "code",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "通行证",
                        "name": "passport",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.Request": {
            "type": "object",
            "properties": {
                "content": {
                    "description": "目前只有 checkpoint3 需要这个请求，都放这",
                    "type": "string"
                },
                "extra_info": {
                    "description": "此字段暂时不用",
                    "type": "string"
                }
            }
        },
        "handler.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "tags": [
        {
            "description": "组织服务",
            "name": "organization"
        },
        {
            "description": "银行服务",
            "name": "bank"
        },
        {
            "description": "终点服务",
            "name": "end"
        }
    ]
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
	Version:     "1.0",
	Host:        "localhost",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "http-theft-bank",
	Description: "a game to study http with go",
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
	swag.Register(swag.Name, &s{})
}
