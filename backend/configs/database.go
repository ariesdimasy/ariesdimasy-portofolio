package configs

import (
	"fmt"
	"log"

	"ariesdimasy-portofolio/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPgDB() {
	dbUser := GetEnv("DB_USER", "postgres")
	dbPassword := GetEnv("DB_PASSWORD", "postgres")
	dbHost := GetEnv("DB_HOST", "localhost")
	dbPort := GetEnv("DB_PORT", "5432")
	dbName := GetEnv("DB_NAME", "postgres")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	fmt.Println("Connected to database successfully")

	err = DB.AutoMigrate(
		&models.User{},
		&models.Biodata{},
		&models.Skill{},
		&models.Sosmed{},
		&models.Education{},
		&models.Experience{},
		&models.Certificate{},
		&models.Project{},
		&models.ProjectImage{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database : ", err)
	}

	fmt.Println("Database migrated successfully")
}
