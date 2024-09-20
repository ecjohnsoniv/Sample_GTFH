package main

import (
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}

func main() {
	app := fiber.New()
	app.Static("/", "./images")

	name := "'Merica x2"
	//pic := "https://upload.wikimedia.org/wikipedia/commons/thumb/d/d3/Bill_Clinton.jpg/800px-Bill_Clinton.jpg"
	pic := "/USA_orthographic.svg"

	app.Get("/", func(c *fiber.Ctx) error {
		return Render(c, Home(name, pic, time.Now()))
	})

	app.Get("/currentTime", func(c *fiber.Ctx) error {
		return Render(c, counter(time.Now()))
	})

	app.Listen(":3000")
}
