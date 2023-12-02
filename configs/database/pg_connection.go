package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type ConfigConn struct {
	env *viper.Viper
	log *log.Logger
}

func NewConnection(env *viper.Viper, log *log.Logger) *ConfigConn {
	return &ConfigConn{
		env, log,
	}
}

func (c *ConfigConn) Open() *gorm.DB {
	const dsnFormat = "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta"

	sslMode := "disable"

	if c.env.GetString("APP_ENV") == "production" {
		sslMode = "require"
	}

	dsn := fmt.Sprintf(dsnFormat,
		c.env.GetString("DB_HOST"),
		c.env.GetString("DB_USERNAME"),
		c.env.GetString("DB_PASSWORD"),
		c.env.GetString("DB_DATABASE"),
		c.env.GetString("DB_PORT"),
		sslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Error),
		SkipDefaultTransaction: true,
	})

	if err != nil {
		c.log.Fatalf("Error connecting to database: %v", err)
	}

	c.log.Println("Database connection established successfully!")

	c.log.Println("Migrating database...")

	err = db.AutoMigrate()
	if err != nil {
		c.log.Fatalf("Error migrating database: %v", err)
	}

	c.log.Println("Database migrated successfully!")

	return db
}
