package cli

import (
	"assignment1/item"
	"fmt"
	"log"
)

const (
	Yes = "y"
	No  = "n"
)

func Init() {
	// To store all the items
	var items = []item.Item{}

	// Accept item details and store them in items slice
	err := acceptItemDetails(&items)
	if err != nil {
		printError(err)
		return
	}

	// Calculate sales tax for each item
	calculateSalesTax(items)

	// Output in command line item name, item price, sales tax liability per item, final price (sales tax + item prize)
	printItems(&items)

}

func calculateSalesTax(items []item.Item) {
	for i := range items {
		items[i].CalculateSalesTax()
	}
}

func printItems(items *[]item.Item) {
	for _, currentItem := range *items {
		fmt.Println("\nItem Name:", currentItem.Name)
		fmt.Println("Item Price:", currentItem.Price)
		fmt.Println("Sales Tax Liability:", currentItem.SalesTax)
		fmt.Println("Final Price:", currentItem.SalesTax+currentItem.Price)
		fmt.Println("----------------------------")
	}
}

func acceptItemDetails(items *[]item.Item) error {
	// While enterItem is true, we keep on entering item details
	enterItem := true

	for enterItem {
		// Get item details from cli
		name, itemType, price, quantity, err := getItemFromCLI()
		if err != nil {
			return err
		}

		// Create item from the entered details
		item, err := item.CreateItem(name, itemType, price, quantity)
		if err != nil {
			return err
		}

		// Store them in items storage
		*items = append(*items, item)

		//Ask for more items
		enterItem, err = enterMore()
		if err != nil {
			return err
		}

	}
	return nil
}

func getItemFromCLI() (name string, itemType string, price float64, quantity int, err error) {

	fmt.Println("\nEnter Item Details")
	fmt.Println("Enter item name, item type, price and quantity with spaces :")

	if _, err = fmt.Scanf("%s", &name); err != nil {
		err = fmt.Errorf("error occurred while entering name : %v", err)
		return
	}

	if _, err = fmt.Scanf("%s", &itemType); err != nil {
		err = fmt.Errorf("error occurred while entering input type : %v", err)
		return
	}

	if _, err = fmt.Scanf("%f", &price); err != nil {
		err = fmt.Errorf("error occurred while entering price : %v", err)
		return
	}

	if _, err = fmt.Scanf("%d\n", &quantity); err != nil {
		err = fmt.Errorf("error occurred while entering quantity : %v", err)
		return
	}

	return
}

func enterMore() (_ bool, err error) {
	var res string

	// take input until we get a y/n
	for res != Yes && res != No {
		fmt.Println("Do you want to enter details of any other item (y/n):")

		if _, err = fmt.Scanf("%s\n", &res); err != nil {
			err = fmt.Errorf("error occurred while entering confirmation: %v", err)
			return
		}

		if res != Yes && res != No {
			fmt.Println("answer in only y or n")
		}

	}
	return res == Yes, nil
}

func printError(err error) {
	log.Println(err)
}
