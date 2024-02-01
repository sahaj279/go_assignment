package consumer

import (
	"fmt"
	"sync"

	"github.com/sahaj279/go_assignment/service/item"
)

func GetInvoice(itemDB chan item.Item, wg *sync.WaitGroup) {
	for item := range itemDB {
		fmt.Println(item.Invoice())
	}
	wg.Done()
}
