package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"store/config"
	"store/database"
	"store/query"
	"store/routes"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Define the IP address flag
	ip := flag.String("ip", ":8007", "IP address and port to listen on")

	// Parse the command-line flags
	flag.Parse()

	// Load configuration
	config.MustLoad()

	// Connect to the database
	database.ConnectDb()

	// Initialize query
	query.SetDefault(database.Database.DB)
	// cron.ScheduleDeviceDataCheck(database.Database.DB)

	// Initialize Fiber
	app := fiber.New()

	// Use the custom rate limiter middleware
	// app.Use(middleware.RateLimiter())

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}))

	// Initialize additional routes
	routes.InitRoutes(app)
	// consumer.InitConsumers()

	// Start the server
	go func() {
		if err := app.Listen(*ip); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

}
