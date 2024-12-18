package config

import (
	"fmt"
	"log"
	"os"

	"github.com/mromero91/go-api-contacts/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Verifica que todas las variables estén presentes
	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbPort == "" {
		log.Fatal("Error: Faltan variables de entorno para la configuración de la base de datos")
	}

	// Construye la cadena de conexión
	psqlInfo := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	fmt.Println(psqlInfo)

	db, err := gorm.Open(postgres.Open("postgres://postgres:secret@db:5432/postgres"))
	// db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	// Confirmación de la conexión
	fmt.Println("🚀 Conexión exitosa a la base de datos PostgreSQL")

	db.AutoMigrate(&models.User{})
	DB = db
}
