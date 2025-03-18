package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/khodaid/Sablon/internal/config/jwt"
	"github.com/khodaid/Sablon/internal/handler"
	"github.com/khodaid/Sablon/internal/middleware"
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

func getEnvAsInt(key string, fallback int) int {

	if value, ok := os.LookupEnv(key); ok {
		if value == "" {
			return fallback
		}

		result, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return result
	}

	return fallback
}

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
	jwt_config := jwt.JwtEnv{}

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

	jwt_config.JwtSecret = getEnv("JWT_SECRET", "rahasia")
	jwt_config.JwtExpiration = getEnvAsInt("JWT_EXPIRATION", 3600)

	csrfSecret := getEnv("CSRF_SECRET", "rahasia")
	csrfExpired := getEnvAsInt("CSRF_EXPIRATION", 3600)

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

		jwtService := jwt.NewJWTService(jwt_config.JwtSecret, jwt_config.JwtExpiration)
		csrfService := jwt.NewServiceCsrfToken(csrfSecret, csrfExpired)
		csrfHandler := handler.NewCsrfHandler(csrfService)

		roleRepository := repositories.NewRoleRepository(g.db)
		userRoleAdminRepository := repositories.NewUserRoleAdminRepository(g.db)
		userStoreRepository := repositories.NewUserStoreRepository(g.db)

		userRepository := repositories.NewUserRepository(g.db)
		userService := service.NewUserService(g.db, userRepository, roleRepository, userStoreRepository, userRoleAdminRepository)
		userHandler := handler.NewUserHandler(userService, jwtService)

		supplierRepository := repositories.NewSupplierRepository(g.db)

		storeRepository := repositories.NewStoreRepository(g.db)
		storeService := service.NewStoreService(storeRepository, supplierRepository)
		storehandler := handler.NewStoreHandler(g.db, storeService)

		authMiddleware := middleware.NewAuthMiddleware(jwtService, userService)
		csrfMiddleware := middleware.NewCSRFMiddleware(csrfService)
		corsMiddleware := middleware.NewCorsMiddleware()

		routingHandler := route.NewRouteHandler(csrfHandler, userHandler, storehandler)
		routingMiddleware := route.NewRouteMiddleware(authMiddleware, csrfMiddleware, corsMiddleware)
		routing := route.NewRoute(routingHandler, routingMiddleware)

		r := routing.InitRoute()
		fmt.Println(app_config)
		app_config.InitApp(r)
	}
}
