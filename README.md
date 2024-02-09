# Car-Rent-gin-golang-

## (Work in Progress)
# Car Rental Service API Documentation

This document provides information about the Car Rental Service API. The API is built using the Gin framework and allows users to perform various operations related to car rental, maintenance, categories, and specifications. (WIP)


Up to this point (2024/02/09)
## Car Endpoints

### GET /api/v1/cars
List all cars.

### PUT /api/v1/cars/create
Update a car.

### POST /api/v1/cars/create
Create a new car.

### DELETE /api/v1/cars/delete/{id}
Delete a car.

### GET /api/v1/cars/{id}
Find car by ID.

## Maintenance Endpoints

### POST /api/v1/cars/{carID}/maintenance/create
Create a new maintenance record for a car.

## Category Endpoints

### POST /api/v1/category/create
Create a new category.

### GET /api/v1/category/list
Retrieve a list of categories.

## Specification Endpoints

### POST /api/v1/specification/create
Create a new specification.

