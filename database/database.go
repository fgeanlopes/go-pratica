package database

import (
	"fmt"
	"log"
	"os"

	"go-pratica/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	godotenv.Load()

	// String de conexão MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados:", err)
	}

	fmt.Println("✅ Conexão com banco de dados estabelecida!")

	// Auto-migrate: cria/atualiza as tabelas automaticamente
	database.AutoMigrate(&models.Client{})
	fmt.Println("✅ Migração concluída!")

	DB = database
}
