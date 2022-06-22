package models

import (
  "fmt"
  "log"

  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)

type Lead struct {
  gorm.Model
  Email      string `gorm:"type:varchar(100);uniqueIndex;not null"`
  Message    string `gorm:"not null"`
}

const dbUser string = "testuser"
const dbPass string = "testpass"
const dbHost string = "localhost"
const dbPort string = "3306"
const dbName string = "testdb"


func Init() (*gorm.DB, error) {
  dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
    dbUser, dbPass, dbHost, dbPort, dbName)

  db, dberr := gorm.Open(mysql.New(mysql.Config{
    DSN: dsn,
    DisableDatetimePrecision: true,
  }), &gorm.Config{});

  if dberr != nil {
    log.Printf("Unable to connect to database: %w", dberr)
    return nil, dberr
  }

  if err := db.AutoMigrate(&Lead{}); err != nil {
    log.Printf("Unable to migrate database: %w", err)
    return nil, err
  }

  return db, nil
}

func CreateNewLead(db *gorm.DB, lead *Lead) error {
  result := db.Create(lead)
  return result.Error
}
