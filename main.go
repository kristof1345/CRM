package main

import (
	"github.com/gofiber/fiber"
)

func setUpRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

func initDatabase() {

}

func main() {
	app := fiber.New()
	setUpRoutes(app)
	app.Listen(3000)

}
