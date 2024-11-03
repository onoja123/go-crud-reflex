# Go CRUD System with SQLite - Using Reflex for Auto-Reload

This project is a simple CRUD system built with Go, Gin, and SQLite. It demonstrates setting up RESTful routes, and it includes configuration for **Reflex**, a tool that automatically reruns tasks when code changes are detected. This setup helps streamline development by automatically running commands, such as tests or compiling code, whenever you make changes.

## Features

- Retrieve all books
- Retrieve a specific book
- Add a new book
- Delete a book

## What is Reflex?

**Reflex** is a lightweight tool that watches your directory for file changes and reruns a specified command whenever it detects modifications. This is useful for automatically running tasks like compiling, linting, or testing, which saves time and keeps your development workflow efficient.

### Setting Up Reflex

To use Reflex in this project:

1. Install Reflex if you haven’t already:

    ```bash
    go install github.com/cespare/reflex@latest
    ```

2. Configure Reflex to watch your Go files. Create a `.reflex` configuration file in your project root with the following contents:

    ```text
    # Rerun main.go whenever a .go file changes
    -r '\.go$' -s -- sh -c "go run main.go"
    ```

3. Run Reflex:

    ```bash
    reflex -c reflex.conf
    ```

Now, Reflex will automatically restart the server whenever you modify a `.go` file, making it easy to see changes in real-time.

## Getting Started

### Prerequisites

- Go 1.16 or higher
- SQLite installed on your system
- Reflex installed (optional, but recommended for development)

### Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/onoja123/go-crud-reflex.git
    cd go-crud-reflex
    ```

2. **Install dependencies:**

    ```bash
    go mod download
    ```

3. **Run the application** (if using Reflex, see above):

    ```bash
    go run main.go
    ```

## API Endpoints

- `GET /books` - Retrieve all books
- `GET /books/:id` - Retrieve a book by ID
- `POST /books` - Add a new book
- `DELETE /books/:id` - Delete a book by ID

## Code Structure

The API routes are defined in the `routes` package:

```go
import (
	"github.com/gin-gonic/gin"
	"go-crud/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetOneBook)
	router.POST("/books", controllers.CreateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
}
