package migration

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"latihan-hris/config"
	"latihan-hris/database/seeders"
)

func RunMigration() {
	dbURL := os.Getenv("DB_URL")

	m, err := migrate.New("file://database/migrations", dbURL)
	if err != nil {
		log.Fatal("Gagal membuat migrasi:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Gagal menjalankan migrasi:", err)
	}

	db := config.DB

	// seeder role dan user
	seeders.RoleSeeds(db)
	seeders.UserSeeds(db)

	log.Println("Migrasi & seeder selesai dijalankan")
}

func RunDropAll() {
	dbURL := os.Getenv("DB_URL")

	m, err := migrate.New("file://database/migrations", dbURL)
	if err != nil {
		log.Fatal("Gagal membuat migrasi:", err)
	}

	// Drop semua tabel
	if err := m.Drop(); err != nil {
		log.Fatal("Gagal drop semua tabel:", err)
	}

	log.Println("Semua tabel berhasil dihapus")
}

func RunDown() {
	dbURL := os.Getenv("DB_URL")

	m, err := migrate.New("file://database/migrations", dbURL)
	if err != nil {
		log.Fatal("Gagal membuat migrasi:", err)
	}

	// Rollback satu langkah migrasi
	if err := m.Steps(-1); err != nil {
		log.Fatal("Gagal menurunkan migrasi:", err)
	}

	log.Println("Satu langkah migrasi berhasil diturunkan")
}
