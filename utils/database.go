// utils/database.go
package utils

import (
	. "api-go/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitializeDatabase() {
	var err error

	// Configuration de GORM pour SQLite en mémoire
	//dsn := "test.db"
	dsn := "host=localhost user=best password=best dbname=best port=5432 sslmode=disable "
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Printf("Erreur lors de la connexion à la base de données: %v\n", err)
		panic("Impossible de se connecter à la base de données")
	}

	// Migration automatique des modèles
	DB.AutoMigrate(&Product{})
}
