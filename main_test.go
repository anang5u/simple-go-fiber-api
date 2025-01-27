package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetRoot(t *testing.T) {
	// Membuat instance aplikasi Fiber
	app := fiber.New()

	// Setup route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World",
		})
	})

	// Membuat request GET ke endpoint "/"
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// Membuat recorder untuk merekam response
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	// Membaca body response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}

	// Verifikasi status code
	assert.Equal(t, 200, resp.StatusCode)

	// Verifikasi isi response body
	assert.Contains(t, string(body), `"message":"Hello, World"`)
}
