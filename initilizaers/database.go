package initilizaers

import (
	model "CrudApi/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("❌ DB_URL empty!")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Assign to global (no :=)
	if err != nil {
		log.Fatal("❌ Connect failed:", err)
	}
	log.Printf("Connected: %T", DB)

	// Ping
	sqlDB, _ := DB.DB()
	sqlDB.Ping()

	if err := DB.AutoMigrate(&model.Post{}); err != nil {
		log.Fatal("❌ Migrate failed:", err)
	}
}
