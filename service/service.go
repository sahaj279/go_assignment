package service

import (
	"github.com/pkg/errors"
	"github.com/sahaj279/go_assignment/config"
)

func Init() error {
	config, err := config.LoadAppConfig()
	if err != nil {
		return errors.Wrap(err, "loadAppConfig")
	}

	db, closeDB, err := Open(config)
	if err != nil {
		return errors.Wrap(err, "open db connection from config")
	}

	defer closeDB()

	repo := NewRepo(db)
	items, err := repo.GetItems()
	if err != nil {
		return errors.Wrap(err, "failed to fetch items from db")
	}

	execute(items)

	return nil
}
