package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/khodaid/Sablon/internal/handler"
	"github.com/khodaid/Sablon/internal/repositories"
	"github.com/khodaid/Sablon/internal/route"
	"github.com/khodaid/Sablon/internal/service"
)

// type RouteConfig struct {
// 	G *gin.Engine
// }

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		if value == "" {
			return fallback
		}
		return value
	}

	return fallback
}

// func NewRoute(g *gin.Engine) *RouteConfig {
// 	return &RouteConfig{g}
// }

func runCommand(g *repository, db_conect string) bool {
	if Migrate || Seed {
		if Migrate {
			g.RunMigrate(db_conect)
		}
		if Seed {
			fmt.Println("Run Seeder")
			g.RunSeed()
		}
		return true
	}
	return false
}

func Run() {
	app_config := appEnv{}
	db_config := dbEnv{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	app_config.APP_NAME = getEnv("APP_NAME", "Khoda App")
	app_config.APP_ENV = getEnv("APP_ENV", "dev")
	app_config.APP_HOST = getEnv("APP_HOST", "localhost")
	app_config.APP_PORT = getEnv("APP_PORT", "3636")

	db_config.DB_CONNECTION = getEnv("DB_CONNECTION", "pgsql")
	db_config.DB_HOST = getEnv("DB_HOST", "127.0.0.1")
	db_config.DB_PORT = getEnv("DB_PORT", "3306")
	db_config.DB_NAME = getEnv("DB_NAME", "khoda")
	db_config.DB_USER = getEnv("DB_USERNAME", "root")
	db_config.DB_PASSWORD = getEnv("DB_PASSWORD", "")

	flag.Parse()
	arg := flag.Arg(0)

	if arg != "" {

	} else {
		g, err := db_config.InitDB()
		if err != nil {
			return
		}

		command := runCommand(g, db_config.DB_CONNECTION)
		if command {
			return
		}

		userRepository := repositories.NewUserRepository(g.db)
		userService := service.NewUserService(userRepository)
		userHandler := handler.NewUserHandler(userService)

		supplierRepository := repositories.NewSupplierRepository(g.db)

		storeRepository := repositories.NewStoreRepository(g.db)
		storeService := service.NewStoreService(storeRepository, supplierRepository)
		storehandler := handler.NewStoreHandler(g.db, storeService)
		routing := route.NewRoute(userHandler, storehandler)

		r := routing.InitRoute()
		fmt.Println(app_config)
		app_config.InitApp(r)
	}
}
