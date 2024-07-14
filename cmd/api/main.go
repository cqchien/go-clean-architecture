package main

import (
	"log"
	"todo/config"
	driverDB "todo/internal/driver/db"
	driverServer "todo/internal/driver/server"

	"github.com/sirupsen/logrus"
)

func main() {
	log.Println("Starting System")

	log.Println("Load env...")
	config := config.GetConfig()

	log.Println("Connect to DB...")
	dbInstance := driverDB.GetPostgresInstance(config, false)

	defer driverDB.ShutdownDBConnection(dbInstance)

	server := driverServer.NewServer(config, dbInstance, logrus.New(), nil)
	if err := server.Bootstrap(); err != nil {
		log.Fatal(err)
	}
}
