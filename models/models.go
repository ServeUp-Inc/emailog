package models

import (
  "fmt"
  "log"
  "errors"

  "gorm.io/gorm"
  "gorm.io/driver/mysql"
  "github.com/go-playground/validator/v10"
  mysqlDefs "github.com/go-sql-driver/mysql"
)

type Lead struct {
  gorm.Model
  Email string `gorm:"type:varchar(100);uniqueIndex;not null" validate:"required,email"`
  Msg string `gorm:"not null"`
}

const mysqlDuplicateEntryCode = 1062


func Init(
  dbUser string,
  dbPass string,
  dbHost string,
  dbPort string,
  dbName string,
) (*gorm.DB, error) {
  dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
    dbUser, dbPass, dbHost, dbPort, dbName)

  db, dberr := gorm.Open(mysql.New(mysql.Config{
    DSN: dsn,
    DisableDatetimePrecision: true,
  }), &gorm.Config{});

  if dberr != nil {
    log.Printf("Unable to connect to database: %v", dberr)
    return nil, dberr
  }

  if err := db.AutoMigrate(&Lead{}); err != nil {
    log.Printf("Unable to migrate database: %v", err)
    return nil, err
  }

  return db, nil
}

func CreateNewLeadIfNotExists(db *gorm.DB, lead *Lead) error {
  validate := validator.New()

  if err := validate.Struct(lead); err != nil {
    return errors.New("Invalid email") 
  }

  result := db.Create(lead)

  // Check for duplicate entry
  var mysqlErr *mysqlDefs.MySQLError
  if errors.As(result.Error, &mysqlErr) && mysqlErr.Number == mysqlDuplicateEntryCode {
    return nil
  }

  return result.Error
}
