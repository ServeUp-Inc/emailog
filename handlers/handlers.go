package handlers

import (
  "gorm.io/gorm"
  "github.com/gofiber/fiber/v2"

  "github.com/ServeUp-Inc/mailog/models"
)

func SendStatusNotFound(c *fiber.Ctx) error {
  return c.SendStatus(fiber.StatusNotFound)
}

func PutLead(db *gorm.DB) func(*fiber.Ctx) error {
  return func(c *fiber.Ctx) error {
    lead := new(models.Lead) 
    if err := c.BodyParser(lead); err != nil {
      return c.SendStatus(fiber.StatusBadRequest)
    }

    if err := models.CreateNewLeadIfNotExists(db, lead); err != nil {
      return c.SendStatus(fiber.StatusBadRequest)
    }

    return c.SendStatus(fiber.StatusOK)
  }
}


