package database

import (
	"fmt"
	log "github.com/andrepriyanto10/server_favaa/configs/logger"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

type ConfigConn struct {
	env *viper.Viper
	log *log.Log
}

func NewConnection(env *viper.Viper, log *log.Log) *ConfigConn {
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

	// make a pool connection
	var pool sync.Pool
	pool.New = func() any {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Error),
			SkipDefaultTransaction: true,
		})
		if err != nil {
			c.log.ErrorLog.Println("Error connecting to database: ", err)
			return nil
		}
		return db
	}

	db := pool.Get().(*gorm.DB)

	c.log.InfoLog.Println("Database connection established successfully!")

	c.log.InfoLog.Println("Migrating database...")

	err := db.AutoMigrate()
	if err != nil {
		c.log.ErrorLog.Fatalf("Error migrating database: %v", err)
	}

	c.log.InfoLog.Println("Database migrated successfully!")

	return db
}
