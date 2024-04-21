package routes

import (
	// "sidd6916/search-engine/views"

	"sidd6916/search-engine/db"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}

func SetRoutes(app *fiber.App) {
	app.Get("/", AuthMiddleware, DashboardHandler)
	app.Post("/", AuthMiddleware, DashboardPostHandler)
	app.Post("/logout", LogoutHandler)

	app.Get("/login", LoginHandler)
	app.Post("/login", LoginPostHandler)

	app.Get("/create", func(c *fiber.Ctx) error {
		u := &db.User{}
		u.CreateAdmin()
		return c.SendString("Admin Created")
	})
}
