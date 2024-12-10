package config

import (
	"fmt"
	"log"

	"github.com/khodaid/Sablon/internal/models"
	"github.com/khodaid/Sablon/internal/seeders"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbEnv struct {
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
}

type repository struct {
	db *gorm.DB
}

func (env *dbEnv) InitDB() (*repository, error) {
	// Menggunakan Mysql
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.DB_USER, env.DB_PASSWORD, env.DB_HOST, env.DB_PORT, env.DB_NAME)
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Menggunakan database PostgreSql
	var dsn string
	if env.DB_PASSWORD == "" {
		dsn = fmt.Sprintf("host=%s user=%s  dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", env.DB_HOST, env.DB_USER, env.DB_NAME, env.DB_PORT)
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", env.DB_HOST, env.DB_USER, env.DB_PASSWORD, env.DB_NAME, env.DB_PORT)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	fmt.Println("Berhasil terhubung ke database")
	return &repository{db}, nil
}

func (r *repository) RunMigrate() {
	// Definisikan flag untuk CLI
	migrate := &Migrate

	if *migrate {
		// Mengambil semua model dari function GetModels
		modelsList := models.GetEntity()

		// Loop over models and run AutoMigrate
		for _, model := range modelsList {
			fmt.Print("Migrating...")
			fmt.Println(model)
			err := r.db.AutoMigrate(model)
			if err != nil {
				log.Fatalf("Migration failed: %v", err)
			}
		}
	}

	// fmt.Println("No command executed. Use --migrate for running migrations.")
}

func (r *repository) RunSeed() {
	seed := &Seed

	if *seed {
		seedList := seeders.ListSeeder()

		for _, seeders := range seedList {
			fmt.Print("Seeding......")
			fmt.Println(seeders)
			err := r.db.CreateInBatches(seeders, 100).Error
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
