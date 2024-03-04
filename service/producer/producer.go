package producer

import (
	"sync"

	"github.com/pkg/errors"
	"github.com/sahaj279/go_assignment/repo"
	"github.com/sahaj279/go_assignment/service/item"
)

func FetchItem(itemChan chan item.Item, wg *sync.WaitGroup, repo *repo.Repository, errChan chan error) {
	defer wg.Done()
	defer close(itemChan)
	defer close(errChan)
	items, err := repo.GetItems()
	if err != nil {
		errChan <- errors.Wrap(err, "failed to fetch items from db")
		return
	}
	for _, item := range *items {
		itemChan <- item
	}
}
