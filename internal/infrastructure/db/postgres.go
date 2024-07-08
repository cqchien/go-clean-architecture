package db

import (
	"fmt"
	"log"
	"todo/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgresInstance(config *config.Configuration, migration bool) *gorm.DB {
	dsn := getDbConnectionString(config)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Can not connect to DB. Error: ", err)
		panic(err)
	}

	if migration {
		errorMigrate := db.AutoMigrate()

		if errorMigrate != nil {
			log.Println("Error when run migrations", errorMigrate)

			panic(errorMigrate)
		}
	}

	return db
}

func ShutdownDBConnection(dbInstance *gorm.DB) {
	sqlDB, err := dbInstance.DB()

	if err != nil {
		log.Fatal("Failed to obtain database object:", err)
		return
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatal("Got error when closing the DB connection:", err)
	}
}

func getDbConnectionString(config *config.Configuration) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		config.DbHost, config.DbUser, config.DbPassword, config.DbName, config.DbPort)
}
