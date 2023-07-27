package main

import (
	"fmt"
	"log"
	"os"

	routes "github.com/Almazatun/qrcode-iod/pkg/http"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	loadENVs()
	app := fiber.New()
	routes.PubRoutes(app)

	log.Fatal(app.Listen(os.Getenv("PORT")))
}

func loadENVs() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("🔴 Error loading .env variables")
	}
}
