package main

import (
	"log"
	"todo/config"
	"todo/internal/infrastructure/db"
	"todo/internal/infrastructure/server"

	"github.com/sirupsen/logrus"
)

func main() {
	log.Println("Starting System")

	log.Println("Load env...")
	config := config.GetConfig()

	log.Println("Connect to DB...")
	dbInstance := db.GetPostgresInstance(config, false)

	defer db.ShutdownDBConnection(dbInstance)

	server := server.NewServer(config, dbInstance, logrus.New(), nil)
	if err := server.Bootstrap(); err != nil {
		log.Fatal(err)
	}
}
