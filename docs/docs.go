// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-02-28 16:12:54.1913979 +0800 CST m=+0.202999101

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "IRIS-Community API document",
        "title": "Swagger IRIS-Community API",
        "contact": {},
        "license": {},
        "version": "1.0"
    },
    "paths": {
        "/auth": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "获取Authorization（用户登录）",
                "operationId": "user-auth",
                "parameters": [
                    {
                        "description": "RequestAuthUser",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.RequestAuthUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.UserAuth"
                        }
                    }
                }
            }
        },
        "/auth/rest": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "重置密码后通过code获取Authorization",
                "operationId": "auth-rest",
                "parameters": [
                    {
                        "description": "RequestAuthRest",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.RequestAuthRest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.UserAuth"
                        }
                    }
                }
            }
        },
        "/auth/user": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "获取当前Authorization对应的用户信息",
                "operationId": "user-auth",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Users"
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "create user",
                "operationId": "create-user",
                "parameters": [
                    {
                        "description": "RequestCreateUser",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.RequestCreateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Users"
                        }
                    }
                }
            }
        },
        "/user/activate": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "activate user",
                "operationId": "activate-user",
                "parameters": [
                    {
                        "description": "RequestActivateUser",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.RequestActivateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Users"
                        }
                    }
                }
            }
        },
        "/user/resendAct": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "resend email to activate user",
                "operationId": "resend-email-activate",
                "parameters": [
                    {
                        "description": "RequestUser",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.RequestUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Users"
                        }
                    }
                }
            }
        },
        "/user/reset": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "reset user password by email",
                "operationId": "reset-password-email",
                "parameters": [
                    {
                        "description": "RequestUser",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.RequestUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Users"
                        }
                    }
                }
            }
        },
        "/user/updatePwd": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "update password",
                "operationId": "update-password",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "RequestUpateUser",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.RequestUpateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Users"
                        }
                    }
                }
            }
        },
        "/verify": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "verify"
                ],
                "summary": "获取验证码",
                "operationId": "create-verifyCode",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.VerifyResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.UserAuth": {
            "type": "object",
            "properties": {
                "AuthCode": {
                    "type": "string"
                },
                "Createtime": {
                    "type": "string"
                },
                "ExpiresIn": {
                    "type": "integer"
                },
                "Id": {
                    "type": "integer"
                },
                "Updatetime": {
                    "type": "string"
                },
                "UserId": {
                    "type": "integer"
                }
            }
        },
        "models.Users": {
            "type": "object",
            "properties": {
                "Createtime": {
                    "type": "string"
                },
                "Email": {
                    "type": "string"
                },
                "Id": {
                    "type": "integer"
                },
                "IsActived": {
                    "type": "boolean"
                },
                "IsBlocked": {
                    "type": "boolean"
                },
                "Password": {
                    "type": "string"
                },
                "Salt": {
                    "type": "string"
                },
                "Updatetime": {
                    "type": "string"
                }
            }
        },
        "rest.RequestActivateUser": {
            "type": "object",
            "properties": {
                "Code": {
                    "type": "string"
                },
                "Id": {
                    "type": "string"
                }
            }
        },
        "rest.RequestAuthRest": {
            "type": "object",
            "properties": {
                "Code": {
                    "type": "string"
                },
                "Id": {
                    "type": "string"
                }
            }
        },
        "rest.RequestAuthUser": {
            "type": "object",
            "properties": {
                "Email": {
                    "type": "string"
                },
                "Password": {
                    "type": "string"
                },
                "VerifyCode": {
                    "type": "string"
                }
            }
        },
        "rest.RequestCreateUser": {
            "type": "object",
            "properties": {
                "Email": {
                    "type": "string"
                },
                "InvitationCode": {
                    "type": "string"
                },
                "Password": {
                    "type": "string"
                },
                "VerifyCode": {
                    "type": "string"
                }
            }
        },
        "rest.RequestUpateUser": {
            "type": "object",
            "properties": {
                "Email": {
                    "type": "string"
                },
                "Password": {
                    "type": "string"
                }
            }
        },
        "rest.RequestUser": {
            "type": "object",
            "properties": {
                "Email": {
                    "type": "string"
                }
            }
        },
        "rest.VerifyResponse": {
            "type": "object",
            "properties": {
                "Code": {
                    "type": "string"
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}