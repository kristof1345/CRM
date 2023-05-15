package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/kristof1345/CRM/database"
)

func setUpRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err := gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Println("Connection open to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setUpRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}
