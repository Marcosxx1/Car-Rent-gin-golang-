# Car Rental API(WORK IN PROGRESS - WIP)
This is a WIP repository that I'm using to study REST API with golang, I'm using gin for routing
This is the third iteration of this project, first that was [this repo](https://github.com/Marcosxx1/third-challenge/blob/main/README.md) that was the third and last challenge for my internship. Then, I implemented this [this other repo](https://github.com/Marcosxx1/Car-Rent-express-node-typescript-/blob/main/README.md) that I used the latter to implement again following the haxagonal architecture

And, we arive in this one, a new adventure with a new language, Golang
Even though this (WIP) repository have similarities with the other two, is/will be more robust, more functionalities, using Golang's capabilities for microservices, concurrency and paralelism
Here is a quick overview of ERD (Entity-Relationship Diagram) also avaliable [here](https://dbdiagram.io/d/car-rent-go-65ad8309ac844320ae6349d4))

### Db diagram
<img width="1980" alt="Db diagram" src="https://github.com/Marcosxx1/Car-Rent-gin-golang-/assets/37447545/968e8525-f5a7-4372-8b2b-3a82f02b9a55">




## Car Rental App Architecture Explanation
For this project I chose to follow Clean Architecture, which is a software architectural pattern introduced by Robert C. Martin. Clean Architecture emphasizes the separation of concerns and the independence of the business logic (domain) from the external concerns such as databases, frameworks, and delivery mechanisms.

- **Domain Layer:** Contains the business logic and entities of the application. It represents the core functionality and rules of the system.

- **Application Layer**: Contains use cases or application-specific business rules. It orchestrates the flow of data between the external layers (like the UI and infrastructure) and the domain layer.

- **Infrastructure Layer**: Contains the implementation details and external concerns such as databases, frameworks, and third-party libraries. It's responsible for interacting with external systems.

As well in with the other projects this architecture enables enhanced modularity, testability, and maintainability.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
  - [Authentication](#authentication)
  - [Category Routes](#categories)
  - [Car Routes](#car_routes)
  - [Specifications Routes](#specifications)
  - [User Routes](#user_routes)
- [Folder Structure](#folder-structure)
- [Contributing](#contributing)
- [Functional Requirements (FR)](#functional-requirements-fr)
- [Non-Functional Requirements (NFR)](#non-functional-requirements-nfr)
- [Business Rules (BR)](#business-rules-br)

## Installation
### WIP

## Usage locally
### WIP

## Usage with docker
### WIP

## API Endpoints

The application offers various API endpoints catering to different functionalities:

### Authentication

- `POST /sessions`: WIP.

## Category Routes
### Create Category

- **Endpoint:** `POST /api/v1/category/create`
- **Description:** Creates a new category.
- **Handler Function:** `category_endpoints.PostCategoryController`

### List Categories
- **Endpoint:** `GET /api/v1/category/list`
- **Description:** Retrieves a list of categories.
- **Handler Function:** `category_endpoints.ListCategoriesController`
- **HTTP Method:** GET

## Car Routes

### Get All Cars

- **Endpoint:** `GET /api/v1/cars`
- **Description:** Retrieves a list of all cars.
- **Handler Function:** `endpoints.GetAllCarsController`
- **HTTP Method:** GET

### Find Car by ID

- **Endpoint:** `GET /api/v1/cars/:id`
- **Description:** Retrieves details of a specific car by ID.
- **Handler Function:** `endpoints.FindCarByIdController`
- **HTTP Method:** GET

### Register Car

- **Endpoint:** `POST /api/v1/cars/create`
- **Description:** Registers a new car.
- **Handler Function:** `endpoints.RegisterCarController`
- **HTTP Method:** POST

### Delete Car

- **Endpoint:** `DELETE /api/v1/cars/delete/:id`
- **Description:** Deletes a car by ID.
- **Handler Function:** `endpoints.DeleteCarController`
- **HTTP Method:** DELETE

### Update Car

- **Endpoint:** `PUT /api/v1/cars/update/:id`
- **Description:** Updates information for a specific car.
- **Handler Function:** `endpoints.UpdateCarController`
- **HTTP Method:** PUT


## Specification Routes
### Create Specification
- **Endpoint:** `POST /api/v1/specification/create`
- **Description:** Creates a new specification.
- **Handler Function:** `specificationendpoints.PostSpecificationController`

## User Routes
### Create User
- **Endpoint:** `POST /api/v1/user/create`
- **Description:** Registers a new user.
- **Handler Function:** `userendpoints.RegisterUserController`
- **HTTP Method:** POST


### Get User by ID
- **Endpoint:** `GET /api/v1/user/:id`
- **Description:** Retrieves user details by ID.
- **Handler Function:** `userendpoints.GetUserByIdController`
- **HTTP Method:** GET

### Update User
- **Endpoint:** `PATCH /api/v1/user/:id`
- **Description:** Updates user information.
- **Handler Function:** `userendpoints.PatchUserController`
- **HTTP Method:** PATCH

### Change Password
- **Endpoint:** `PATCH /api/v1/user/:id/change-password`
- **Description:** Changes the password for a user.
- **Handler Function:** `userendpoints.ChangePasswordController`
- **HTTP Method:** PATCH


## Folder Structure
<pre><details>
<summary>Folder structure, click here</summary>
Car-Rent-gin-golang-/
┣ .vscode/
┣ api/
┃ ┣ application/
┃ ┃ ┣ repositories/
┃ ┃ ┃ ┣ car-repository.go
┃ ┃ ┃ ┣ category-repository.go
┃ ┃ ┃ ┣ maintenance-repository.go
┃ ┃ ┃ ┣ specification-repository.go
┃ ┃ ┃ ┗ user-repository.go
┃ ┃ ┗ use-cases/
┃ ┃   ┣ car-use-cases/
┃ ┃ ┃ ┃ ┣ car-use-case-tests/
┃ ┃ ┃ ┃ ┃ ┣ mocks-and-structs/
┃ ┃ ┃ ┃ ┃ ┃ ┣ mock-car-input.go
┃ ┃ ┃ ┃ ┃ ┃ ┣ mock-car-output.go
┃ ┃ ┃ ┃ ┃ ┃ ┗ mock-car-repository.go
┃ ┃ ┃ ┃ ┃ ┣ delete-car-use-case_test.go
┃ ┃ ┃ ┃ ┃ ┣ get-all-cars-use-case_test.go
┃ ┃ ┃ ┃ ┃ ┗ get-car-by-id-use-case_test.go
┃ ┃ ┃ ┃ ┣ delete-car-use-case.go
┃ ┃ ┃ ┃ ┣ get-all-cars-use-case.go
┃ ┃ ┃ ┃ ┣ get-car-by-id-use-case.go
┃ ┃ ┃ ┃ ┣ post-car-use-case.go
┃ ┃ ┃ ┃ ┗ put-car-use-case.go
┃ ┃   ┣ category-use-cases/
┃ ┃ ┃ ┃ ┣ get-all-categories-use-case.go
┃ ┃ ┃ ┃ ┣ post-category-use-case.go
┃ ┃ ┃ ┃ ┗ post-category-use-case_test.go
┃ ┃   ┣ maintenance-use-cases/
┃ ┃ ┃ ┃ ┗ post-maintenance-use-case.go
┃ ┃   ┣ repo-utils/
┃ ┃ ┃ ┃ ┗ convert-specification-to-dto.go
┃ ┃   ┣ specification-use-cases/
┃ ┃ ┃ ┃ ┣ post-specification-use-case-test.go
┃ ┃ ┃ ┃ ┗ post-specification-use-case.go
┃ ┃   ┗ user-use-cases/
┃ ┃ ┃   ┣ change-password-use-case.go
┃ ┃ ┃   ┣ get-user-by-id-use-case.go
┃ ┃ ┃   ┣ patch-user-use-case.go
┃ ┃ ┃   ┗ post-user-use-case.go
┃ ┣ domain/
┃ ┃ ┣ car-maintenance.go
┃ ┃ ┣ car.entity.go
┃ ┃ ┣ category.go
┃ ┃ ┣ maintenance.go
┃ ┃ ┣ specification.go
┃ ┃ ┣ user.go
┃ ┃ ┗ user_car.go
┃ ┗ infra/
┃   ┣ database/
┃ ┃ ┃ ┣ postgres/
┃ ┃ ┃ ┃ ┗ db-config/
┃ ┃ ┃ ┃   ┗ connnection.go
┃ ┃ ┃ ┣ pg-car-repository.go
┃ ┃ ┃ ┣ pg-category-repository.go
┃ ┃ ┃ ┣ pg-maintenance-repository.go
┃ ┃ ┃ ┣ pg-specification-repository.go
┃ ┃ ┃ ┗ pg-user-repository.go
┃   ┣ http/
┃ ┃ ┃ ┣ controllers/
┃ ┃ ┃ ┃ ┣ car-controller/
┃ ┃ ┃ ┃ ┃ ┣ car-dtos/
┃ ┃ ┃ ┃ ┃ ┃ ┗ car-dto.go
┃ ┃ ┃ ┃ ┃ ┗ car-endpoints/
┃ ┃ ┃ ┃ ┃   ┣ delete-car.go
┃ ┃ ┃ ┃ ┃   ┣ get-all-cars.go
┃ ┃ ┃ ┃ ┃   ┣ get-car-by-id.go
┃ ┃ ┃ ┃ ┃   ┣ post-car.go
┃ ┃ ┃ ┃ ┃   ┗ update-car.go
┃ ┃ ┃ ┃ ┣ category-controller/
┃ ┃ ┃ ┃ ┃ ┣ category-dtos/
┃ ┃ ┃ ┃ ┃ ┃ ┗ category-dto.go
┃ ┃ ┃ ┃ ┃ ┗ category-endpoints/
┃ ┃ ┃ ┃ ┃   ┣ get-all-categories.go
┃ ┃ ┃ ┃ ┃   ┗ post-category.go
┃ ┃ ┃ ┃ ┣ maintenance-controller/
┃ ┃ ┃ ┃ ┃ ┣ maintenance-dtos.go/
┃ ┃ ┃ ┃ ┃ ┃ ┗ maintenance-dto.go
┃ ┃ ┃ ┃ ┃ ┗ maintenance-endpoints/
┃ ┃ ┃ ┃ ┃   ┗ post-maintenance.go
┃ ┃ ┃ ┃ ┣ specification-controller/
┃ ┃ ┃ ┃ ┃ ┣ specification-dtos/
┃ ┃ ┃ ┃ ┃ ┃ ┗ specification-dto.go
┃ ┃ ┃ ┃ ┃ ┗ specification-endpoints/
┃ ┃ ┃ ┃ ┃   ┗ post-specification.go
┃ ┃ ┃ ┃ ┗ user-controller/
┃ ┃ ┃ ┃   ┣ hash-password/
┃ ┃ ┃ ┃ ┃ ┃ ┣ hash-password.go
┃ ┃ ┃ ┃ ┃ ┃ ┗ verify-password.go
┃ ┃ ┃ ┃   ┣ user-dtos/
┃ ┃ ┃ ┃ ┃ ┃ ┗ user-dto.go
┃ ┃ ┃ ┃   ┗ user-endpoints/
┃ ┃ ┃ ┃ ┃   ┣ change-passwor-controller.go
┃ ┃ ┃ ┃ ┃   ┣ get-user-by-id-controller.go
┃ ┃ ┃ ┃ ┃   ┣ patch-user-controller.go
┃ ┃ ┃ ┃ ┃   ┗ post-user-controller.go
┃ ┃ ┃ ┣ docs/
┃ ┃ ┃ ┃ ┣ docs.go
┃ ┃ ┃ ┃ ┣ swagger.json
┃ ┃ ┃ ┃ ┗ swagger.yaml
┃ ┃ ┃ ┣ routes/
┃ ┃ ┃ ┃ ┣ car-rotues.go
┃ ┃ ┃ ┃ ┣ category-routes.go
┃ ┃ ┃ ┃ ┣ maintenance-routes.go
┃ ┃ ┃ ┃ ┣ specification-routes.go
┃ ┃ ┃ ┃ ┗ user-routes.go
┃ ┃ ┃ ┣ tmp/
┃ ┃ ┃ ┃ ┗ main.exe
┃ ┃ ┃ ┗ main.go
┃   ┗ validation_errors/
┃ ┃   ┣ error.go
┃ ┃   ┗ validator.go
┣ tmp/
┃ ┣ build-errors.log
┃ ┗ main.exe
┣ .air.toml
┣ .env
┣ .gitignore
┣ go.mod
┣ go.sum
┗ README.md
</details>
</pre>



## Functional Requirements (FR)

### Car Registration
- Ability to register a new car.
- Car registration requires attributes like make, model, and year.
- Cars are registered with an availability status by default.
- Car registration limited to administrators.

### Car Listing
- Ability to list all available cars.
- Users can filter available cars by category, manufacturer, and name.
- User authentication not required for listing all cars.
### Car Specification Registration
- It should be possible to register specifications for a car.
- Specifications could include engine type, fuel efficiency, etc.
- Specifications must not be registered for non-existing cars.
- Users should be administrators to register car specifications.
- It should be possible to list all specifications.
- It should be possible to list all cars along with their specifications.

### Car Image Registration
- It must be possible to register images for cars.
- Users can upload more than one image for the same car.
- Users responsible for registration should be administrators.
- Multer should be used to handle image uploads.
- It should be possible to list all cars along with their images.

### Car Rent
- It should be possible to register a car rental.
- Car rental details should include start date, end date, user ID, etc.
- Rentals must be for a minimum of 24 hours.
- Users should not be able to register a new rental if they already have one.
- Users should not be able to register a new rental for a car that is already rented.


## Non-Functional Requirements (NFR)

- Validation mechanisms ensure accurate user inputs.
- Robust error handling enhances user experience.
- User authentication and authorization ensure secure operations.
- Comprehensive documentation covers codebase, APIs, and setup.
- Focus on performance, security, and scalability enhances application quality.

## Business Rules (BR)

### Car Registration
- Prevent car registration with an existing license plate.
- Disallow changing license plates of already registered cars.
- Register cars with an availability status by default.
- Only administrators can perform car registration.

### Car Specification Registration
- Specifications must not be registered for non-existing cars.
- Registering the same specification for the same car should not be allowed.
- Only administrators can register car specifications.

### Car Image Registration
- Users can upload multiple images for the same car.
- Users responsible for registration should be administrators.

### Car Rent
- Rentals must be for a minimum of 24 hours.
- Users must not be able to register a new rental if they already have one.
- Users must not be able to register a new rental for a car that is already rented.
## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests for improvements or bug fixes.

The Car Rental App's architecture provides clarity, modularity, and scalability. This documentation offers insights into installation, usage, functionalities, and requirements, fostering a better understanding of the application's ecosystem.
## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.


