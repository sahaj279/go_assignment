package cli

import (
	"assignment1/item"
	"errors"
	"fmt"
	"strings"
)

func AcceptItemDetails(items *[]item.Item) error {
	// Get item details from cli
	name, itemType, price, quantity, err := getItemFromCLI()
	if err != nil {
		return err
	}

	//Validate the entered details
	err = validateEnteredDetails(name, itemType, price, quantity)
	if err != nil {
		return err
	}

	// Store them in items storage
	item := item.CreateItem(name, itemType, price, quantity)
	*items = append(*items, item)

	//Ask for more items
	enterOtherItem, err := enterOtherItemDetail()
	if err != nil {
		return err
	}
	if enterOtherItem {
		err = AcceptItemDetails(items)
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

	return name, itemType, price, quantity, nil

}

func validateEnteredDetails(name string, itemType string, price float64, quantity int) error {
	if len(name) < 1 {
		return errors.New("item name should have more than 1 characters")
	}
	if price <= 0 {
		return errors.New("price can not be negative or 0")

	}
	if quantity <= 0 {
		return errors.New("quantity can not be negative or 0")

	}
	itemType = strings.ToLower(itemType)
	if itemType != "raw" && itemType != "manufactured" && itemType != "imported" {
		return errors.New("entered item type is invalid")

	}
	return nil
}

func enterOtherItemDetail() (enterOtherItem bool, err error) {
	fmt.Println("Do you want to enter details of any other item (y/n):")
	var response string
	if _, err = fmt.Scanf("%s\n", &response); err != nil {
		err = fmt.Errorf("error occurred while entering confirmation: %v", err)
		return
	}
	switch response {
	case "y":
		enterOtherItem = true
	case "n":
		enterOtherItem = false
	default:
		err = errors.New("answer in only y or n")
	}
	return
}

func PrintItemDetails(name string, price float64, salesTax float64, finalPrice float64) {
	fmt.Println("\nItem Name:", name)
	fmt.Println("Item Price:", price)
	fmt.Println("Sales Tax Liability:", salesTax)
	fmt.Println("Final Price:", finalPrice)
	fmt.Println("----------------------------")
}

func PrintError(err error) {
	fmt.Println("Error", err)
}
