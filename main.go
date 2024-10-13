package main

import (
	"log"
	"my-fiber-app/database"
	"my-fiber-app/middleware"
	"my-fiber-app/routes"
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/gofiber/fiber/v2"
)



func main() {
	
	// Connect to database
	database.Connect()
	   
	// config.LoadConfig()
	app := fiber.New(fiber.Config{
		// Prefork: true,
		JSONEncoder: sonic.Marshal,    // Use sonic's Marshal for encoding
		JSONDecoder: sonic.Unmarshal,  // Use sonic's Unmarshal for decoding
	})

	// Middleware
	app.Use(recover.New())				 // Recover middleware to catch panics
	app.Use(cors.New(cors.Config{          // CORS configuration
		AllowOrigins:    "http://localhost:3000", // Allow all origins or specify
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))
	// Helmet adds security headers
	app.Use(helmet.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,  // Compression level for faster performance
	}))

	// Request ID to track requests
	app.Use(requestid.New())

	
	// Rate limiter to prevent abuse
	app.Use(limiter.New(limiter.Config{
		Max:        50,                // Allow 50 requests
		Expiration: 30 * time.Second,   // Within 30 seconds
	}))

	// Cache middleware
	app.Use(cache.New(cache.Config{
		Expiration:   1 * time.Minute,
		CacheControl: true,
	}))

	// CSRF protection middleware
	// app.Use(csrf.New(csrf.Config{
	// 	CookieName:     "_csrf",
	// 	CookieSameSite: "Strict",
	// 	Expiration:     1 * time.Hour,
	// }))

	// Idempotency middleware
	app.Use(idempotency.New())

	// Add pprof middleware to expose pprof routes
	app.Use(pprof.New())
	
	// Use the logging middleware
    app.Use(middleware.LoggingMiddleware)

	// Setup routes
	routes.SetupUserRoutes(app)
	routes.SetupProductRoutes(app)

	// app.Get("/csrf-token", func(c *fiber.Ctx) error {
	// 	token := c.Locals("csrf")
	// 	return c.JSON(fiber.Map{"csrf_token": token})
	//  })

	// Monitoring middleware
	app.Get("/metrics", monitor.New(monitor.Config{Title: "App Metrics"}))


	// Health check route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		log.Default().Println("Hello, World!")
		return c.SendString("Go Fiber App")
	})

	
	// Start the server
    if err := app.Listen(":3000"); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }

}

