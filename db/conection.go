package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
    // Cargar las variables de entorno desde el archivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}


func DBConection() {
	// Obtener el DSN desde la variable de entorno
    DSN := os.Getenv("DSN")
    if DSN == "" {
        log.Fatal("DSN is not set in the environment variables")
    }
	
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	}else{
		log.Println("DB connected")
	}
}