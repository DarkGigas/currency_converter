package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	t.Run("returns code 404 - route not found", func(t *testing.T) {
		app := fiber.New()

		req := httptest.NewRequest(http.MethodGet, "/api/hello", nil)
		resp, err := app.Test(req)

		assert.Nil(t, err, "Error creating request")

		assert.Equal(t, http.StatusOK, resp.StatusCode, "Status code should be 200 OK")

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		assert.Nil(t, err, "Error reading response body")

		assert.Equal(t, "Hello, World!", string(body), "Unexpected response body")
	})
}
