package main

import (
	"fmt"
	"log"
	"os"

	"os/signal"
	"sidd6916/search-engine/db"
	"sidd6916/search-engine/routes"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":4300"
	} else {
		port = ":" + port
	}

	app := fiber.New(fiber.Config{
		AppName:     "Search Engine v1.0.1",
		IdleTimeout: 5 * time.Second,
	})
	app.Use(compress.New())
	db.InitDB()
	routes.SetRoutes(app)

	// Start our server
	func() {
		err := app.Listen(port)
		if err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c // Block the main thread to continue execution until interrupted

	app.Shutdown()
	fmt.Println("Shutting down...")

}
