package persistence

import (
	"fmt"
	"go-poc/configs"
	"log"

	"go-poc/internal/domain/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDbConnection(configs *configs.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configs.Postgres.Host,
		configs.Postgres.Port,
		configs.Postgres.UserName,
		configs.Postgres.Password,
		configs.Postgres.DatabaseName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	// Tabloyu oluştur (eğer yoksa)
	if err := db.AutoMigrate(&entities.Order{}, &entities.OrderProduct{}, &entities.OutboxMessage{}, &entities.InboxMessage{}); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	return db
}
