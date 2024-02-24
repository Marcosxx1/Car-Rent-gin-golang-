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
        "/api/v1/cars": {
            "get": {
                "description": "Retrieve a list of cars with pagination support.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car"
                ],
                "summary": "List all cars",
                "operationId": "list-car",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number (default is 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page (default is 10)",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of cars",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtos.CarOutputDTO"
                            }
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/validation_errors.HTTPError"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/cars/create": {
            "post": {
                "description": "Create a new car with the provided information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car"
                ],
                "summary": "Create a new car",
                "operationId": "post-car",
                "parameters": [
                    {
                        "description": "Car information to be created",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CarInputDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully created car",
                        "schema": {
                            "$ref": "#/definitions/dtos.CarOutputDTO"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/validation_errors.HTTPErrorCar"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/cars/delete/{id}": {
            "delete": {
                "description": "Delete a car with the provided ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car"
                ],
                "summary": "Delete a car",
                "operationId": "delete-car",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Car ID to be deleted",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Car deleted successfully"
                    },
                    "400": {
                        "description": "Error details",
                        "schema": {
                            "$ref": "#/definitions/validation_errors.HTTPErrorCar"
                        }
                    }
                }
            }
        },
        "/api/v1/cars/update/:id": {
            "put": {
                "description": "Update a car with the provided ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car"
                ],
                "summary": "Update a car",
                "operationId": "put-car",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Car ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Car information to be updated",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CarInputDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully updated car",
                        "schema": {
                            "$ref": "#/definitions/dtos.CarOutputDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/validation_errors.HTTPErrorCar"
                        }
                    }
                }
            }
        },
        "/api/v1/cars/{id}": {
            "get": {
                "description": "Find a car with the provided id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car"
                ],
                "summary": "Find car",
                "operationId": "get-car",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Car ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "car",
                        "schema": {
                            "$ref": "#/definitions/dtos.CarOutputDTO"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/validation_errors.HTTPError"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/category/create": {
            "post": {
                "description": "Create a new category with the provided information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Create a new category",
                "operationId": "post-category",
                "parameters": [
                    {
                        "description": "Category information to be created",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/categorydtos.CategoryInputDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully created category",
                        "schema": {
                            "$ref": "#/definitions/categorydtos.CategoryOutputDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/category/list": {
            "get": {
                "description": "Retrieve a list of all categories.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Retrieve a list of categories",
                "operationId": "list-categories",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit the number of categories to be retrieved",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset for pagination of categories",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of categories",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/categorydtos.CategoryOutputDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/validation_errors.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/maintenance/:carID/maintenance/create": {
            "post": {
                "description": "Create a new maintenance with the provided information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Maintenance"
                ],
                "summary": "Create a new maintenance",
                "operationId": "post-maintenance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CarID",
                        "name": "carID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Maintenance information to be created",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/maintenancedtos.MaintenanceInputDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully created maintenance",
                        "schema": {
                            "$ref": "#/definitions/maintenancedtos.MaintenanceOutputDTO"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/validation_errors.HTTPErrorCar"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/maintenance/maintenances": {
            "get": {
                "description": "Retrieve a list of maintenance records for a specific car",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Maintenance"
                ],
                "summary": "Retrieve a list of maintenance records",
                "operationId": "get-maintenance-list",
                "responses": {
                    "200": {
                        "description": "List of maintenance records",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/maintenancedtos.MaintenanceOutputDTO"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/validation_errors.HTTPErrorCar"
                        }
                    }
                }
            }
        },
        "/api/v1/maintenance/{maintenanceID}": {
            "delete": {
                "description": "Delete a maintenance with the provided ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Maintenance"
                ],
                "summary": "Delete a maintenance",
                "operationId": "delete-maintenance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Maintenance ID to be deleted",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Maintenance deleted successfully"
                    },
                    "400": {
                        "description": "Error details",
                        "schema": {
                            "$ref": "#/definitions/validation_errors.HTTPErrorCar"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update an existing maintenance with the provided information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Maintenance"
                ],
                "summary": "Update an existing maintenance",
                "operationId": "patch-maintenance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Maintenance ID",
                        "name": "maintenanceID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Maintenance information to be updated",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/maintenancedtos.MaintenanceInputDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully updated maintenance",
                        "schema": {
                            "$ref": "#/definitions/maintenancedtos.MaintenanceOutputDTO"
                        }
                    }
                }
            }
        },
        "/api/v1/specification/create": {
            "post": {
                "description": "Create a new specification with the provided information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Specification"
                ],
                "summary": "Create a new specification",
                "operationId": "post-specification",
                "parameters": [
                    {
                        "description": "Specification information to be created",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/specificationdtos.SpecificationInputDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully created specification",
                        "schema": {
                            "$ref": "#/definitions/specificationdtos.SpecificationOutputDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/validation_errors.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "categorydtos.CategoryInputDTO": {
            "type": "object",
            "required": [
                "description",
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "example": "'description is required'"
                },
                "name": {
                    "type": "string",
                    "example": "'name is required'"
                }
            }
        },
        "categorydtos.CategoryOutputDTO": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dtos.CarInputDTO": {
            "type": "object",
            "required": [
                "available",
                "brand",
                "category_id",
                "daily_rate",
                "description",
                "fine_amount",
                "license_plate",
                "name"
            ],
            "properties": {
                "available": {
                    "type": "boolean"
                },
                "brand": {
                    "type": "string"
                },
                "category_id": {
                    "type": "string"
                },
                "daily_rate": {
                    "type": "number"
                },
                "description": {
                    "type": "string"
                },
                "fine_amount": {
                    "type": "number"
                },
                "license_plate": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "specification": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/specificationdtos.SpecificationInputDto"
                    }
                }
            }
        },
        "dtos.CarOutputDTO": {
            "type": "object",
            "properties": {
                "available": {
                    "type": "boolean"
                },
                "brand": {
                    "type": "string"
                },
                "category_id": {
                    "type": "string"
                },
                "daily_rate": {
                    "type": "number"
                },
                "description": {
                    "type": "string"
                },
                "fine_amount": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                },
                "license_plate": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "specification": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/specificationdtos.SpecificationOutputDto"
                    }
                }
            }
        },
        "maintenancedtos.MaintenanceInputDTO": {
            "type": "object",
            "required": [
                "last_maintenance_date",
                "maintenance_completion_date",
                "maintenance_status",
                "maintenance_type",
                "next_maintenance_due_date",
                "odometer_reading"
            ],
            "properties": {
                "car_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "labor_cost": {
                    "type": "integer",
                    "minimum": 1
                },
                "last_maintenance_date": {
                    "type": "string"
                },
                "maintenance_completion_date": {
                    "type": "string"
                },
                "maintenance_duration": {
                    "type": "string"
                },
                "maintenance_notes": {
                    "type": "string"
                },
                "maintenance_status": {
                    "type": "string"
                },
                "maintenance_type": {
                    "type": "string"
                },
                "next_maintenance_due_date": {
                    "type": "string"
                },
                "odometer_reading": {
                    "type": "integer",
                    "minimum": 0
                },
                "parts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/maintenancedtos.PartInputDTO"
                    }
                },
                "parts_cost": {
                    "type": "integer",
                    "minimum": 1
                },
                "scheduled_maintenance": {
                    "type": "boolean"
                }
            }
        },
        "maintenancedtos.MaintenanceOutputDTO": {
            "type": "object",
            "required": [
                "last_maintenance_date",
                "maintenance_completion_date",
                "next_maintenance_due_date"
            ],
            "properties": {
                "car_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "labor_cost": {
                    "type": "integer"
                },
                "last_maintenance_date": {
                    "type": "string"
                },
                "maintenance_completion_date": {
                    "type": "string"
                },
                "maintenance_duration": {
                    "type": "string"
                },
                "maintenance_notes": {
                    "type": "string"
                },
                "maintenance_status": {
                    "type": "string"
                },
                "maintenance_type": {
                    "type": "string"
                },
                "next_maintenance_due_date": {
                    "type": "string"
                },
                "odometer_reading": {
                    "type": "integer"
                },
                "parts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/maintenancedtos.PartOutputDTO"
                    }
                },
                "parts_cost": {
                    "type": "integer"
                },
                "scheduled_maintenance": {
                    "type": "boolean"
                }
            }
        },
        "maintenancedtos.PartInputDTO": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "cost": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "replacement_date": {
                    "type": "string"
                }
            }
        },
        "maintenancedtos.PartOutputDTO": {
            "type": "object",
            "properties": {
                "cost": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "maintenance_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "replacement_date": {
                    "type": "string"
                }
            }
        },
        "specificationdtos.SpecificationInputDto": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "car_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "specificationdtos.SpecificationOutputDto": {
            "type": "object",
            "properties": {
                "car_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "validation_errors.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "validation_errors.HTTPErrorCar": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 422
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "['name is required'",
                        " 'description is required'",
                        " 'licenseplate is required'",
                        " 'brand is required'",
                        " 'A car needs a category to be registered']"
                    ]
                },
                "message": {
                    "type": "string",
                    "example": ""
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Serviço de locação de carros",
	Description:      "Serviço utilizando o framework gin",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
