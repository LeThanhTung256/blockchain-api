package db

import (
	"fmt"
	"tayson/blockchain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "p@ssW0rd"
	dbname   = "blockchain"
)

var DB *gorm.DB

func Connect() {
	psqldsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s  port=%s sslmode=disable", host, user, password, dbname, port)
	connection, err := gorm.Open(postgres.Open(psqldsn), &gorm.Config{})
	CheckError(err)

	DB = connection

	DB.AutoMigrate(models.User{})
}

func CheckError(err error) {
	if err != nil {
		panic("Could not connect to database")
	}
}
