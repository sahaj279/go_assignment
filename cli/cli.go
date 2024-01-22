package cli

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/sahaj279/go-assignment/item"
)

const (
	Yes = "y"
	No  = "n"
)

type Cli struct {
	itemSvc item.ItemHandler
}

func NewCli(itemSvc item.ItemHandler) Cli {
	return Cli{
		itemSvc: itemSvc,
	}
}

func (c Cli) Init() error {
	// While moreItem is true, we keep on entering item details
	moreItem := true
	for moreItem {
		name, itemType, price, quantity, err := enterItem()
		if err != nil {
			return errors.Wrap(err, "init")
		}

		item, err := c.itemSvc.CreateItem(name, itemType, price, quantity)
		if err != nil {
			return errors.Wrap(err, "init")
		}

		tax := c.itemSvc.CalculateTax(&item)
		printItem(&item, tax)

		moreItem, err = enterMore()
		if err != nil {
			return errors.Wrap(err, "init")
		}
	}

	return nil
}

func printItem(item *item.Item, tax float64) {
	fmt.Println("----------------------------")
	fmt.Println("Item Name:", item.Name)
	fmt.Println("Item Price:", item.Price)
	fmt.Println("Sales Tax Liability:", tax)
	fmt.Println("Final Price:", tax+item.Price)
	fmt.Println("----------------------------")
}

func enterItem() (name string, itemType string, price float64, quantity int, err error) {
	fmt.Println("\nEnter Item Details")
	fmt.Println("Enter item name, item type, price and quantity with spaces :")

	if _, err = fmt.Scanf("%s", &name); err != nil {
		err = errors.Wrap(err, "enter item : name error")
		return
	}

	if _, err = fmt.Scanf("%s", &itemType); err != nil {
		err = errors.Wrap(err, "enter item : type error")
		return
	}

	if _, err = fmt.Scanf("%f", &price); err != nil {
		err = errors.Wrap(err, "enter item : price error")
		return
	}

	if _, err = fmt.Scanf("%d\n", &quantity); err != nil {
		err = errors.Wrap(err, "enter item : quantity error")
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
			err = errors.Wrap(err, "enterMore")
			return
		}

		if res != Yes && res != No {
			fmt.Println("answer in only y or n")
		}

	}
	return res == Yes, nil
}

func LogError(err error) {
	log.Println(err)
}
