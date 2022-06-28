package server

import (
  "time"
  "log"

  "github.com/gofiber/fiber/v2"
  "gorm.io/gorm"

  "github.com/ServeUp-Inc/emailog/models"
  "github.com/ServeUp-Inc/emailog/configs"
  "github.com/ServeUp-Inc/emailog/handlers"
)

func createRoutes(app *fiber.App, db *gorm.DB) {
  v1 := app.Group("/v1")

  v1.Put("/", handlers.PutLead(db))

  // Catch all
  app.Use(handlers.SendStatusNotFound)
}

func Create() *fiber.App {
  dbConfig, dbConfigErr := config.ReadDBConfigFromEnv()
  if dbConfigErr != nil {
    log.Printf("Unable to read database configurations: %v", dbConfigErr)
    panic(dbConfigErr)
  }

  db, dbErr := models.Init(
    dbConfig.User,
    dbConfig.Pass,
    dbConfig.Host,
    dbConfig.Port,
    dbConfig.Name)
  if dbErr != nil {
    log.Printf("Unable to initialize database: %v", dbErr)
    panic(dbErr)
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

