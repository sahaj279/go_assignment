package main

import (
	"assignment1/cli"
	"assignment1/item"
)

func main() {
	// To store all the items
	var items = []item.Item{}

	// Accept item details with validations and store them in items slice
	err := cli.AcceptItemDetails(&items)
	if err != nil {
		cli.PrintError(err)
		return
	}

	// Calculate sales tax for each item
	calculateSalesTax(items)

	// Output in command line item name, item price, sales tax liability per item, final price (sales tax + item prize)
	printItemDetails(&items)

}

func calculateSalesTax(items []item.Item) {
	for i := range items {
		items[i].CalculateSalesTax()
	}
}

func printItemDetails(items *[]item.Item) {
	for _, currentItem := range *items {
		cli.PrintItemDetails(currentItem.Name, currentItem.Price, currentItem.SalesTax, currentItem.Price+currentItem.SalesTax)
	}
}
