package main

import (
  "time"
  "log"

  "github.com/gofiber/fiber/v2"

  "github.com/ServeUp-Inc/mailog/models"
)

func main() {
  db, err := models.Init()
  if err != nil {
    log.Printf("Unable to initialize database: %w", err)
    panic(err)
  }

  //TODO Create lead data 
  // validate email address
  // Escape message string
  newLead := models.Lead{Email: "test@gmail.com", Message: "this is a msg"}
  models.CreateNewLead(db, &newLead)

  app := fiber.New(fiber.Config {
    ServerHeader: "ServeUp",
    AppName: "Emailog",
    ReadTimeout: 1*time.Millisecond,
    WriteTimeout: 5*time.Millisecond,
    IdleTimeout: 1*time.Millisecond,
    DisableKeepalive: true,
  })

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("hello")
  })

  app.Listen(":4000")
}
