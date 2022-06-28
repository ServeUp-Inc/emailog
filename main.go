package main

import (
  "log"
  "fmt"

  "github.com/ServeUp-Inc/emailog/configs"
  "github.com/ServeUp-Inc/emailog/server"
)

func main() {
  serverConfig, serverConfigErr := configs.ReadServerConfigFromEnv()
  if serverConfigErr != nil {
    log.Printf("Unable to read server configurations: %v", serverConfigErr)
    panic(serverConfigErr)
  }

  server := server.Create()

  log.Fatal(server.Listen(fmt.Sprintf("%s:%s", serverConfig.Host, serverConfig.Port)))
}
