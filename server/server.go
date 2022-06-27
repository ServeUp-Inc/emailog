package server

import (
  "time"
  "log"

  "github.com/gofiber/fiber/v2"
  "gorm.io/gorm"

  "github.com/ServeUp-Inc/emailog/models"
  "github.com/ServeUp-Inc/emailog/handlers"
)

func createRoutes(app *fiber.App, db *gorm.DB) {
  v1 := app.Group("/v1")

  v1.Put("/", handlers.PutLead(db))

  // Catch all
  app.Use(handlers.SendStatusNotFound)
}

func Create() *fiber.App {
  db, err := models.Init()
  if err != nil {
    log.Printf("Unable to initialize database: %v", err)
    panic(err)
  }

  app := fiber.New(fiber.Config {
    ServerHeader: "ServeUp",
    AppName: "Emailog",
    ReadTimeout: 1*time.Millisecond,
    WriteTimeout: 5*time.Millisecond,
    IdleTimeout: 1*time.Millisecond,
    DisableKeepalive: true,
  })

  createRoutes(app, db)

  return app
}

