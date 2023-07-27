package routes

import (
	"github.com/Almazatun/qrcode-iod/pkg/http/handler"
	"github.com/gofiber/fiber/v2"
)

// Public routes.
func PubRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/")

	// Serve the static index.html file when the root URL is requested
	route.Static("", "pkg/http/static/index.html")

	//QR code handler
	qrcodeHandler := handler.QRCodeHandlerInstance{}

	route.Post("qrcode", qrcodeHandler.Create)
}
