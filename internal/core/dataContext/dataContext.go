package dataContext

import (
	"database/sql"
	"fmt"
	"time"
	"to-do-api/internal/core/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(host string, port string, user string, password string, dbname string) *gorm.DB {
	var (
		dbContext *gorm.DB
		err       error
	)
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		host,
		port,
		user,
		password,
		dbname)

	if dbContext, err = gorm.Open(postgres.Open(connString), &gorm.Config{}); err != nil {
		logger.Logger.Error(err)
		return nil
	}

	// Migrate(DbContext)

	var sqlDB *sql.DB
	if sqlDB, err = dbContext.DB(); err != nil {
		logger.Logger.Error(err)
		return nil
	}
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return dbContext
}
