package main

import (
	"fmt"
	"os"
	"strings"

	"jk-api/api/routes/v1"
	"jk-api/cmd/bootstrap"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	level := strings.ToUpper(entry.Level.String())
	message := fmt.Sprintf("[%s] [%s] - %s\n", timestamp, level, entry.Message)
	return []byte(message), nil
}

var log = logrus.New()

func main() {
	log.SetFormatter(&CustomFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)

	bootstrap.InitApp()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] [${status}] - ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 Welcome to Fiber! Use prefix '/api' for API routes.")
	})

	routes.Setup(app)

	for _, r := range app.Stack() {
		for _, route := range r {
			log.Infof("📌 Route: %-6s %s", route.Method, route.Path)
		}
	}

	log.Info("🚀 Starting Fiber server on http://localhost:5000")
	if err := app.Listen(":5000"); err != nil {
		log.WithError(err).Fatal(err)
	}
}
