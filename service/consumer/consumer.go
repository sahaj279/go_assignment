package consumer

import (
	"fmt"
	"sync"
	"time"

	"github.com/sahaj279/go_assignment/service/item"
)

func GetInvoice(itemDB chan item.Item, wg *sync.WaitGroup) {
	for item := range itemDB {
		// to make this process take time 1sec
		time.Sleep(1000000000)
		fmt.Println(item.Invoice())
	}
	wg.Done()
}
