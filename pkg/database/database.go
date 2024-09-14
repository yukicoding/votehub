package database

import (
	"fmt"
	"sync"
	"time"
	configs "yukicoding/voteHub/configs"
	"yukicoding/voteHub/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

func GetDB() *gorm.DB {
	return db
}

func Init(config *configs.Config) error {
	var initErr error

	dbOnce.Do(func() {
		dsn := config.GetPostgreSQLDSN()
		logger.Warn(dsn)
		// dsn := "host=124.71.213.188 port=5432 user=postgres password=root dbname=votehub sslmode=disable"
		// logger.Info(dsn)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			initErr = fmt.Errorf("failed to connect to database: %w", err)
			return
		}

		sqlDB, err := db.DB()
		if err != nil {
			initErr = fmt.Errorf("failed to get database instance: %w", err)
			return
		}

		sqlDB.SetMaxOpenConns(config.PostgreSQL.MaxOpenConns)
		sqlDB.SetMaxIdleConns(config.PostgreSQL.MaxIdleConns)

		if config.PostgreSQL.ConnMaxLifetime != "" {
			duration, err := time.ParseDuration(config.PostgreSQL.ConnMaxLifetime)
			if err != nil {
				initErr = fmt.Errorf("invalid ConnMaxLifetime: %w", err)
				return
			}
			sqlDB.SetConnMaxLifetime(duration)

		}
	})
	return initErr
}
