package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

type datos struct {
	Carne  string 
	Nombre string 
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Escuchando peticiones")
	})

	datos := datos{
		Carne:  "201700521",
		Nombre: "Ludwing Gabriel Paz Hernandez",
	}

	app.Get("/data", func(c *fiber.Ctx) error {
		return c.JSON(datos)
	})
	log.Fatal(app.Listen(":3000"))
}