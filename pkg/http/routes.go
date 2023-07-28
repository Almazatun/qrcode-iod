package routes

import (
	"path/filepath"

	"github.com/Almazatun/qrcode-iod/pkg/http/handler"
	"github.com/gofiber/fiber/v2"
)

// Public routes.
func PubRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/")

	absPath, err := filepath.Abs("pkg/http/static/index.html")

	if err != nil {
		panic(err)
	}

	// Serve the static index.html file when the root URL is requested
	route.Static("", absPath)

	//QR code handler
	qrcodeHandler := handler.QRCodeHandlerInstance{}

	route.Post("qrcode", qrcodeHandler.Create)
}
