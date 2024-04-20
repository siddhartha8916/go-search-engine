package routes

import (
	"fmt"
	"sidd6916/search-engine/views"
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}

type loginForm struct {
	Email    string `form:email`
	Password string `form:password`
}

type settingsform struct {
	Amount   int  `form:"amount"`
	SearchOn bool `form:"searchOn"`
	AddNew   bool `form:"addNew"`
}

func SetRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return Render(c, views.Home())
	})

	app.Post("/", func(c *fiber.Ctx) error {
		time.Sleep(2 * time.Second)
		input := settingsform{}
		if err := c.BodyParser(&input); err != nil {
			return c.SendString("<h2>Error: Something went wrong</h2>")
		}
		fmt.Println(input)
		return c.SendStatus(200)
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return Render(c, views.Login())
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		time.Sleep(2 * time.Second)
		input := loginForm{}
		if err := c.BodyParser(&input); err != nil {
			return c.SendString("<h2>Error : Something went wrong !</h2>")
		}
		return c.SendStatus(200)

	})
}
