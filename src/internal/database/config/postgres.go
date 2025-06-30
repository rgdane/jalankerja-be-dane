package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

type logWriter struct{}

func (w logWriter) Printf(format string, args ...interface{}) {
	logMessage("INFO", fmt.Sprintf(format, args...))
}

func logMessage(level string, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	log.SetFlags(0)
	log.Printf("[%s] [%s] - %s", timestamp, level, message)
}

/*
getDsn constructs a DB connection URL from environment variables.

It returns an empty string if any of the required environment variables are not set.

The variables are:

- DB_HOST: the hostname of the DB server
- DB_USERNAME: the username to use for the connection
- DB_PASSWORD: the password to use for the connection
- DB_DATABASE: the name of the database to connect to
- DB_PORT: the port to use for the connection
- DB_CONNECTION: the protocol to use (PostgreSQL, MySQL, SQLite, and SQL Server are supported)
- DB_SSL_DISABLED: if set to "enable", the connection will be established

refer to the GORM official documentation: https://gorm.io/docs/connecting_to_the_database.html
*/
func getDsn() string {
	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")
	conn := os.Getenv("DB_CONNECTION")
	ssl_mode := os.Getenv("DB_SSL_DISABLED")

	if host == "" || user == "" || dbname == "" || port == "" {
		logMessage("FATAL", "❌ Missing required DB config in .env")
		os.Exit(1)
	}

	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
		conn, user, password, host, port, dbname,
		func() string {
			if ssl_mode != "enable" {
				return "disable"
			}
			return ssl_mode
		}())
	return dsn
}

func PostgresInit() bool {
	newLogger := gormLogger.New(
		logWriter{},
		gormLogger.Config{
			SlowThreshold:             0,
			LogLevel:                  gormLogger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(postgres.Open(getDsn()), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		logMessage("ERROR", fmt.Sprintf("❌ Failed to connect to database: %v", err))
		return false
	}

	DB = db
	logMessage("INFO", "✅ Connected to database")
	return true
}
