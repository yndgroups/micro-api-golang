// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "杨大琼",
            "url": "http://localhost",
            "email": "826422592@qq.com"
        },
        "license": {
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/file/delete": {
            "post": {
                "description": "主要用于文件删除",
                "tags": [
                    "文件上传管理"
                ],
                "summary": "文件删除",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer token",
                        "description": "授权令牌",
                        "name": "accessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "应用参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.File"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "400": {
                        "description": "请求失败",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "404": {
                        "description": "请求失败",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "请求失败",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/file/getList": {
            "post": {
                "description": "主要用于获取文件列表",
                "tags": [
                    "文件上传管理"
                ],
                "summary": "获取文件列表",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer token",
                        "description": "授权令牌",
                        "name": "accessToken",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "400": {
                        "description": "请求失败",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "404": {
                        "description": "请求失败",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "请求失败",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/file/multipleUpload/{sizeType}": {
            "post": {
                "description": "主要用于多尺寸单个文件上传",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "文件上传管理"
                ],
                "summary": "多尺寸多个文件上传",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer token",
                        "description": "授权令牌",
                        "name": "accessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "上传类型: 1：单尺寸上传 2：多尺寸上传",
                        "name": "sizeType",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "300": {
                        "description": "请求失败",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "400": {
                        "description": "请求失败",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "请求失败",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/file/singleUpload/{sizeType}": {
            "post": {
                "description": "主要用于多尺寸单个文件上传",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "文件上传管理"
                ],
                "summary": "多尺寸单个文件上传",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer token",
                        "description": "授权令牌",
                        "name": "accessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "上传类型: 1：单尺寸上传 2：多尺寸上传",
                        "name": "sizeType",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "400": {
                        "description": "请求失败",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "404": {
                        "description": "请求失败",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "请求失败",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "结果状态码",
                    "type": "integer",
                    "default": 0
                },
                "data": {
                    "description": "结果"
                },
                "msg": {
                    "description": "状态信息",
                    "type": "string",
                    "default": "请求失败"
                }
            }
        },
        "model.File": {
            "type": "object",
            "required": [
                "filePath",
                "sizeType"
            ],
            "properties": {
                "filePath": {
                    "description": "文件路径",
                    "type": "string"
                },
                "sizeType": {
                    "description": "尺寸类型",
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/api/resources",
	Schemes:          []string{},
	Title:            "资源管理服务提供者文档",
	Description:      "swagger go API文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
