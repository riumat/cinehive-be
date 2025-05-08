package migrations

import (
	"log"

	"github.com/riumat/cinehive-be/database"
	"github.com/riumat/cinehive-be/models"
	"gorm.io/gorm"
)

func createEnums(db *gorm.DB) {
	// Crea il tipo enumerativo content_type_enum
	err := db.Exec(`
			DO $$ BEGIN
					IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'content_type_enum') THEN
							CREATE TYPE content_type_enum AS ENUM ('movie', 'tv');
					END IF;
			END $$;
	`).Error
	if err != nil {
		log.Fatalf("Errore durante la creazione del tipo enumerativo content_type_enum: %v", err)
	}

	// Crea il tipo enumerativo request_status_enum
	err = db.Exec(`
			DO $$ BEGIN
					IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'request_status_enum') THEN
							CREATE TYPE request_status_enum AS ENUM ('pending', 'accepted', 'rejected');
					END IF;
			END $$;
	`).Error
	if err != nil {
		log.Fatalf("Errore durante la creazione del tipo enumerativo request_status_enum: %v", err)
	}
}

func MigrateDB(db *gorm.DB) {
	createEnums(db)
	err := db.AutoMigrate(
		&models.User{},
		&models.Content{},
		&models.Person{},
		&models.Relationship{},
		&models.Watchlist{},
	)

	if err != nil {
		log.Fatalf("Errore durante la migrazione: %v", err)
	}
}

func Migrate() {
	database.ConnectDB()

	MigrateDB(database.DB)

	log.Println("Migrazioni completate con successo!")
}
