package config

import (
  "os"
  "fmt"
)

const envDBName string = "DB_NAME"
const envDBUser string = "DB_USER"
const envDBPass string = "DB_PASS"
const envDBHost string = "DB_HOST"
const envDBPort string = "DB_PORT"

const envServerHost string = "SERVER_HOST"
const envServerPort string = "SERVER_PORT"

type ServerConfig struct {
  Host string
  Port string
}

type DBConfig struct {
  Name string
  User string
  Pass string
  Host string
  Port string
}

func checkEnvVar(envVarName string, envVar string) error {
  if len(envVar) == 0 {
    return fmt.Errorf("Environment variable: %s not found", envVarName)
  }
  return nil
}

func ReadDBConfigFromEnv() (DBConfig, error) {
  dbName := os.Getenv(envDBName)
  if err := checkEnvVar(envDBName, dbName); err != nil { return DBConfig{}, err }

  dbUser:= os.Getenv(envDBUser)  
  if err := checkEnvVar(envDBUser, dbUser); err != nil { return DBConfig{}, err }

  dbPass := os.Getenv(envDBPass)  
  if err := checkEnvVar(envDBPass, dbPass); err != nil { return DBConfig{}, err }

  dbHost := os.Getenv(envDBHost)  
  if err := checkEnvVar(envDBHost, dbHost); err != nil { return DBConfig{}, err }

  dbPort := os.Getenv(envDBPort)  
  if err := checkEnvVar(envDBPort, dbPort); err != nil { return DBConfig{}, err }

  return DBConfig{
    Name: dbName,
    User: dbUser,
    Pass: dbPass,
    Host: dbHost,
    Port: dbPort,
  }, nil
}

func ReadServerConfigFromEnv() (ServerConfig, error) {
  host := os.Getenv(envServerHost)
  if err := checkEnvVar(envServerHost, host); err != nil { return ServerConfig{}, err }

  port := os.Getenv(envServerPort)  
  if err := checkEnvVar(envServerPort, port); err != nil { return ServerConfig{}, err }

  return ServerConfig{
    Host: host,
    Port: port,
  }, nil
}
