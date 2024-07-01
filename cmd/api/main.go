package main

import (
	"log"
	"todo/config"
	"todo/db"
	"todo/internal/server"

	"github.com/sirupsen/logrus"
)

func main() {
	log.Println("Starting System")

	log.Println("Load env...")
	config := config.GetConfig()
	log.Println("Load env successfully!!")

	log.Println("Connect to DB...")
	db := db.GetPostgresInstance(config, false)
	log.Println("Connect to DB successfully")

	server := server.NewServer(config, db, logrus.New(), nil)
	if err := server.Bootstrap(); err != nil {
		log.Fatal(err)
	}
}
