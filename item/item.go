package item

import (
	"assignment1/item/enum"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

type ItemHandler interface {
	CreateItem(string, string, float64, int) (Item, error)
	CalculateSalesTax()
}

type Item struct {
	Name     string
	Type     enum.ItemType
	Price    float64
	Quantity int
	SalesTax float64
}

func CreateItem(name string, itemType string, price float64, quantity int) (item Item, err error) {

	item = Item{Name: name, Price: price, Quantity: quantity}
	item.Type, err = enum.ItemTypeString(itemType)
	// Returns error if item type is invalid
	if err != nil {
		return
	}

	// Returns error if item details validation failed
	err = validateItem(item)
	if err != nil {
		err = errors.Wrap(err, "item invalid")
		return
	}

	return
}

func validateItem(item Item) error {
	return validation.ValidateStruct(&item,
		validation.Field(&item.Name, validation.Length(2, 0)),
		validation.Field(&item.Quantity, validation.By(checkNegativeValue)),
		validation.Field(&item.Price, validation.By(checkNegativeValue)))
}

func checkNegativeValue(value interface{}) error {
	err := errors.New("negative value")
	switch t := value.(type) {
	case int:
		if t < 0 {
			return err
		}
	case float64:
		if t < 0.0 {
			return err
		}
	}

	return nil
}

func (item *Item) CalculateSalesTax() {
	rawTax := 0.125 * item.Price
	switch item.Type {
	case enum.Raw:
		item.SalesTax = rawTax
	case enum.Manufactured:
		item.SalesTax = rawTax + (0.02 * (rawTax + item.Price))
	case enum.Imported:
		importDuty := item.Price * 0.1
		finalCost := item.Price + importDuty
		if finalCost <= 100 {
			item.SalesTax = importDuty + 5
		} else if finalCost <= 200 {
			item.SalesTax = importDuty + 10
		} else {
			item.SalesTax = importDuty + finalCost*0.05
		}
	}
}
