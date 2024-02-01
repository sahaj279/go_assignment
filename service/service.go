package service

import (
	"sync"

	"github.com/pkg/errors"
	"github.com/sahaj279/go_assignment/config"
	"github.com/sahaj279/go_assignment/repo"
	"github.com/sahaj279/go_assignment/service/consumer"
	"github.com/sahaj279/go_assignment/service/item"
	"github.com/sahaj279/go_assignment/service/producer"
)

func Init() error {
	config, err := config.LoadAppConfig()
	if err != nil {
		return errors.Wrap(err, "loadAppConfig")
	}

	db, closeDB, err := repo.Open(config)
	if err != nil {
		return errors.Wrap(err, "open db connection from config")
	}

	defer closeDB()

	repo := repo.NewRepo(db)

	itemChan := make(chan item.Item)
	errChan := make(chan error)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go producer.FetchItem(itemChan, &wg, repo, errChan)

	go consumer.GetInvoice(itemChan, &wg)

	for err := range errChan {
		return errors.Wrap(err, "fetchItem")
	}

	wg.Wait()

	return nil
}
