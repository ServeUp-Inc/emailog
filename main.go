package main

import (
  "log"

  "github.com/ServeUp-Inc/emailog/server"
)

func main() {
  server := server.Create()

  log.Fatal(server.Listen(":4000"))
}
