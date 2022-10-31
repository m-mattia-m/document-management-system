package database

import (
	"document-management-system/models"
	"document-management-system/pkg/environment"
	"fmt"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB *gorm.DB
)

var (
	host     = ""
	port     = ""
	username = ""
	password = ""
	name     = ""
)

func Client() *gorm.DB {
	config := mysqlDriver.Config{
		User:   username,
		Passwd: password,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", host, port),
		DBName: name,
		Params: map[string]string{
			"charset":   "utf8mb4",
			"parseTime": "True",
			"loc":       "Local",
		},
		AllowNativePasswords: true,
	}
	dsn := config.FormatDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(dsn)
		log.Println("[SETUP GORM]: can not open the connection to the database\t -> Error: " + err.Error())
		return nil
	}
	return db
}

func init() {
	godotenv.Load()
	log.Println("[SETUP GORM]: Status -> Grom-init is loaded")
	host, _ = environment.GetEnv("DB_HOST")
	port, _ = environment.GetEnv("DB_PORT")
	username, _ = environment.GetEnv("DB_USERNAME")
	password, _ = environment.GetEnv("DB_PASSWORD")
	name, _ = environment.GetEnv("DB_NAME")
}

func InitTables(db *gorm.DB) bool {
	err := db.AutoMigrate(&models.Docuemnt{})
	if err != nil {
		log.Println("[SETUP GORM]: create Table was failed \t\t -> Error: " + err.Error())
		return false
	}
	return true
}
