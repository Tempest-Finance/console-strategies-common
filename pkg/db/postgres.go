package db

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var pg *gorm.DB

func InitPostgres(config Config) error {
	if pg != nil {
		return nil
	}

	gormDB, err := gorm.Open(
		postgres.Open(PConn(config)),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.LogLevel(config.LogLevel)),
		},
	)
	if err != nil {
		return err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.ConnLifeTime) * time.Second)

	if err = sqlDB.Ping(); err != nil {
		return err
	}

	pg = gormDB

	return nil
}

func Instance() *gorm.DB {
	return pg
}
