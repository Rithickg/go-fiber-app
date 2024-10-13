package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// LoggingMiddleware logs details of incoming requests.
func LoggingMiddleware(c *fiber.Ctx) error {
    start := time.Now()

    // Process the request
    err := c.Next()

    // Log the request details
    log.Printf("Method: %s, Path: %s, Status: %d, Time: %s\n", c.Method(), c.Path(), c.Response().StatusCode(), time.Since(start))

    return err
}
