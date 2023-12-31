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
        "/book/addbook": {
            "get": {
                "tags": [
                    "图书模块"
                ],
                "summary": "添加图书",
                "parameters": [
                    {
                        "type": "string",
                        "description": "书名",
                        "name": "bookname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "作者",
                        "name": "author",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "封面图片",
                        "name": "img",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "描述",
                        "name": "desc",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "tag",
                        "name": "tag",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "图书状态",
                        "name": "isreturn",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/book/booklists": {
            "get": {
                "tags": [
                    "图书模块"
                ],
                "summary": "图书列表",
                "responses": {
                    "200": {
                        "description": "code\":\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/book/updatebook": {
            "get": {
                "tags": [
                    "图书模块"
                ],
                "summary": "修改图书借阅状态",
                "parameters": [
                    {
                        "type": "string",
                        "description": "图书ID",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "图书状态",
                        "name": "isreturn",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "tags": [
                    "用户模块"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "302": {
                        "description": "message\",\"data\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/signin": {
            "post": {
                "tags": [
                    "用户模块"
                ],
                "summary": "注册用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "重新确认密码",
                        "name": "repassword",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "QQ号",
                        "name": "qq",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "图书借阅系统",
	Description:      "Go练手项目",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
