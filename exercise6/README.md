# Product Management with Golang and Gin-Gonic

![Golang Logo](https://golang.org/doc/gopher/frontpage.png)

This repository demonstrates a simple implementation of Product Management API using Golang and Gin-Gonic framework. The API allows users to perform basic CRUD (Create, Read, Update, Delete) operations on products, with support for authentication using BasicAuth and JWT (JSON Web Tokens).

## Features

- Create a new product by providing product details as JSON input.
- Update an existing product's details using a JSON input.
- Delete a product by its ID.
- Retrieve a list of all products.
- Basic Authentication to protect certain endpoints.
- JWT Authentication for secure access to protected endpoints.

## Installation and Usage

Follow the steps below to set up and run the Product Management API on your local machine.

### Prerequisites

- Go (Golang) installed on your system.
- Git installed on your system.

### Step 1: Clone the Repository

```bash
git clone https://github.com/your-username/your-repo.git
cd your-repo
```

### Step 2: Set Up the Configuration

Edit the configuration files `development.yaml` and `production.yaml` in the `configs` directory to set your desired database credentials and JWT secret.

### Step 3: Install Dependencies

```bash
go mod download
```

### Step 4: Run the API

```bash
go run cmd/main.go
```

The API will run on `http://localhost:8080` by default.

## File Structure

The project follows a modular and organized file structure to improve maintainability and scalability.

```
my-web-api/
  |- cmd/
  |   |- main.go
  |
  |- internal/
  |   |- api/
  |   |   |- handlers/
  |   |   |   |- handler1.go
  |   |   |   |- handler2.go
  |   |   |   |- ...
  |   |   |
  |   |   |- middlewares/
  |   |   |   |- middleware1.go
  |   |   |   |- middleware2.go
  |   |   |   |- ...
  |   |   |
  |   |   |- models/
  |   |   |   |- model1.go
  |   |   |   |- model2.go
  |   |   |   |- ...
  |   |   |
  |   |   |- routes/
  |   |   |   |- routes.go
  |   |   |
  |   |   |- server.go
  |   |
  |   |- config/
  |   |   |- config.go
  |   |
  |   |- data/
  |   |   |- database.go
  |   |
  |   |- utils/
  |   |   |- utility1.go
  |   |   |- utility2.go
  |   |   |- ...
  |   |
  |   |- constants/
  |   |   |- constants.go
  |
  |- pkg/
  |   |- mypackage1/
  |   |   |- mypackage1.go
  |   |
  |   |- mypackage2/
  |   |   |- mypackage2.go
  |   |   |- ...
  |
  |- migrations/
  |   |- migration1.sql
  |   |- migration2.sql
  |   |- ...
  |
  |- configs/
  |   |- development.yaml
  |   |- production.yaml
  |
  |- tests/
  |   |- handler_test.go
  |   |- middleware_test.go
  |   |- model_test.go
  |   |- ...
  |
  |- go.mod
  |- go.sum
  |- README.md
```

## API Endpoints

### POST /products

Create a new product. It receives product details as JSON input.

### PUT /products/{product_id}

Update a product's details. It receives updated product details as JSON input.

### DELETE /products/{product_id}

Delete a product by its ID.

### GET /products

Retrieve a list of all products.

---

Feel free to explore the code and modify it according to your requirements. If you have any questions or feedback, please create an issue or submit a pull request.

Happy coding! 🚀

---

Please make sure to replace `your-username` and `your-repo` in the clone command with your actual GitHub username and repository name. Additionally, modify the content of the `configs/development.yaml` and `configs/production.yaml` files with your desired configuration settings.

[This](README.md) provides an overview of the project, the file structure, the installation steps, and the available API endpoints. It serves as a quick guide for users who want to understand and use your Product Management API implementation.

Feel free to customize [this](README.md) to suit your specific project or repository needs. Happy scraping!

Or navigate [back](../README.md) to see the overview of the Golang Learning Repository!
