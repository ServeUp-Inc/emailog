package main

import (
  "time"
  "github.com/gofiber/fiber/v2"
  //"gorm.io/gorm"
  //"gorm.io/driver/mysql"
)

func main() {
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
