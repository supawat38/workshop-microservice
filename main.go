package main

import (
	controllers "app/app/controllers/cornjob"

	"app/pkg/configs"
	"app/pkg/middleware"
	"app/pkg/routes"
	"app/pkg/utils"
	"app/platform/database"
	"app/platform/logger"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

func main() {
	// session.InitSession()
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("ALLOW_ACCESS"),
		AllowHeaders:     "*",
		AllowCredentials: true,
		AllowMethods:     "*",
	}))

	// Define zap
	logger.InitLogger()
	logger.SugarLogger.Infof("Initial Logger")

	// Define gorm
	DBTableName := os.Getenv("DB_SEVER_HOST_DBNAME_WRITE")
	err := database.PostgreSQLConnection(DBTableName)
	MIGRATE_DATABASE, _ := strconv.Atoi(os.Getenv("MIGRATE_DATABASE"))
	if MIGRATE_DATABASE == 1 {
		database.Init()
	}
	if err != nil {
		panic(err)
	}

	// Middlewares. - ดัก ตรวจสอบ
	middleware.FiberMiddleware(app)

	// Routes - Private.
	routes.PrivateRoutes(app)

	// Routes - Public.
	routes.PublicRoutes(app)

	// Cronjob (ไม่ได้ใช้)
	go controllers.CronJob(app)

	// routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	// Start server
	utils.StartServer(app)
}
