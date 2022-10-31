package main

import (
	"document-management-system/pkg/api"
	"document-management-system/pkg/database"
	"log"
	"os"
)

func main() {
	status := database.InitTables(database.DB)
	if !status {
		log.Println("[SETUP GORM]: an error has occurred with a table.")
		os.Exit(1)
	}

	api.Router()

}

func init() {
	database.DB = database.Client()
}
