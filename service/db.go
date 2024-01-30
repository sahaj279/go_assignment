package service

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/sahaj279/go_assignment/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//go:generate mockgen -source=db.go -destination=db_mock.go -package=service
type DB interface {
	GetItems() (items []Item, err error)
}

type Repository struct {
	db *gorm.DB
}

func Open(cfg config.AppConfig) (*gorm.DB, func(), error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
	if err != nil {
		err = errors.Wrap(err, "failed to setup mysql connection")
		return nil, nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not set sql.DB params")
	}

	sqlDB.SetConnMaxIdleTime(cfg.Database.MaxConnectionIdleTime)
	sqlDB.SetConnMaxLifetime(cfg.Database.MaxConnectionLifeTime)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConnections)

	closeDB := func() {
		if err := sqlDB.Close(); err != nil {
			log.Printf("failed to close db connections %v", err)
		}
	}

	return db, closeDB, nil
}

func NewRepo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) GetItems() (items *[]Item, err error) {
	if err := r.db.Find(&items).Error; err != nil {
		err = errors.Wrap(err, "fetching items from database failed")
		return nil, err
	}

	return items, nil
}
