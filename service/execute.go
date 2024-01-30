package service

import (
	"fmt"
	"sync"
)

func execute(items *[]Item) {
	itemChan := make(chan Item)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Produce
	go fetchItem(items, itemChan, &wg)

	// Consume
	go getInvoice(itemChan, &wg)

	wg.Wait()
}

func fetchItem(items *[]Item, itemChan chan Item, wg *sync.WaitGroup) {
	for _, item := range *items {
		itemChan <- item
	}
	close(itemChan)
	wg.Done()
}

func getInvoice(itemDB chan Item, wg *sync.WaitGroup) {
	for item := range itemDB {
		fmt.Println(item.Invoice())
	}
	wg.Done()
}
