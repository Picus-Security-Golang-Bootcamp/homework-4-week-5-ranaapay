package postgres

import (
	"PicusHomework4/src/pkg/httpErrors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func NewPsqlDB() (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("PATIKA_DB_HOST"),
		os.Getenv("PATIKA_DB_PORT"),
		os.Getenv("PATIKA_DB_USERNAME"),
		os.Getenv("PATIKA_DB_NAME"),
		os.Getenv("PATIKA_DB_PASSWORD"),
	)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, httpErrors.OpenDBError
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, httpErrors.GetSQLDBError
	}
	if err = sqlDB.Ping(); err != nil {
		return nil, httpErrors.PingSQLDBError
	}
	return db, nil
}
