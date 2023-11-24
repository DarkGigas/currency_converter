# Golang API with Fiber

This is a simple Golang API using the Fiber web framework, the purpose of this app is to convert a specific currency to another.

## Getting Started

1. Install Go on your machine: <https://golang.org/doc/install>
2. Install Fiber: `go get -u github.com/gofiber/fiber/v2`
3. Clone this repository: `git clone <repository-url>`
4. Navigate to the project directory: `cd <project-directory>`
5. Run the application: `go run main.go`

The server will start at <http://localhost:3000>.

## Sample Route

- Endpoint: `GET /api/hello`
  - Description: Returns a greeting message.

## Chain of Responsibility Middleware

The API includes a sample implementation of the chain of responsibility pattern with two middleware functions.
