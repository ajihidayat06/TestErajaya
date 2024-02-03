package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DBConnection *gorm.DB

type DBConnectionStruct struct {
	Username string
	Password string
	Port     string
	Hostname string
	DBName   string
}

func StartConnectionDb(connectionStruct DBConnectionStruct) (*gorm.DB, error) {
	log.Println("Start connecting DB...")
	dbUrl := fmt.Sprintf(`postgres://%s:%s@%s:%s/%s`, connectionStruct.Username, connectionStruct.Password, connectionStruct.Hostname, connectionStruct.Port, connectionStruct.DBName)

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Success connecting DB")

	return db, nil
}

func MigrateModel(db *gorm.DB, model ...interface{}) error {
	log.Println("Start migrate model...")

	err := db.AutoMigrate(model...)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Success migrate model")

	return err
}
