package lead

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kristof1345/CRM/database"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func NewLead(c *fiber.Ctx) {
	lead := new(Lead)
	db := database.DBConn

	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead

	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead found")
		return
	}
	db.Delete(&lead)
	c.JSON("Lead deleted")
}
