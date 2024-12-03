// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/order/checkout": {
            "post": {
                "description": "make a checkout for an order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Make an order checkout",
                "parameters": [
                    {
                        "description": "Checkout",
                        "name": "checkout",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.CheckoutRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.CheckoutDTO"
                        }
                    },
                    "400": {
                        "description": "when bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "406": {
                        "description": "when invalid params or invalid object",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/order/display": {
            "get": {
                "description": "Get order list for a display",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get order list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtos.OrderDisplayDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "when bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/order/{order_id}/preparation-status": {
            "put": {
                "description": "Update the preparation status for an order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Update order preparation status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order Identification",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Order Request Params",
                        "name": "preparation_status_update",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.PreparationStatusUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "when bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.CheckoutRequest": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "string"
                },
                "products_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "controllers.PreparationStatusUpdateRequest": {
            "type": "object",
            "properties": {
                "preparation_status": {
                    "type": "string"
                }
            }
        },
        "dtos.CheckoutDTO": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "method": {
                    "$ref": "#/definitions/dtos.MethodType"
                },
                "orderId": {
                    "type": "string"
                },
                "paymentLink": {
                    "type": "string"
                }
            }
        },
        "dtos.MethodType": {
            "type": "string",
            "enum": [
                "PIX",
                "CREDIT_CARD"
            ],
            "x-enum-varnames": [
                "PIX",
                "CREDIT_CARD"
            ]
        },
        "dtos.OrderDisplayDTO": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dtos.OrderItemDisplayDTO"
                    }
                },
                "order_id": {
                    "type": "string"
                },
                "preparation_status": {
                    "type": "string"
                }
            }
        },
        "dtos.OrderItemDisplayDTO": {
            "type": "object",
            "properties": {
                "product_name": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
