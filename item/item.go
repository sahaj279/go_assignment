package item

import (
	"github.com/sahaj279/go-assignment/item/enum"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

//go:generate mockgen -source=item.go -destination=mock_item/mock_item.go
type ItemHandler interface {
	CreateItem(string, string, float64, int) (Item, error)
	CalculateSalesTax() float64
}

type Item struct {
	Name     string
	Type     enum.ItemType
	Price    float64
	Quantity int
}

func CreateItem(name string, itemType string, price float64, quantity int) (item Item, err error) {

	item = Item{Name: name, Price: price, Quantity: quantity}
	item.Type, err = enum.ItemTypeString(itemType)
	// Returns error if item type is invalid
	if err != nil {
		err = errors.Wrap(err, "CreateItem")
		return
	}

	err = validate(item)
	if err != nil {
		err = errors.Wrap(err, "CreateItem")
		return
	}

	return
}

func validate(item Item) error {
	return validation.ValidateStruct(&item,
		validation.Field(&item.Name, validation.Length(2, 0)),
		validation.Field(&item.Quantity, validation.By(positiveValue)),
		validation.Field(&item.Price, validation.By(positiveValue)))
}

func positiveValue(value interface{}) error {
	err := errors.New("negative value")
	switch t := value.(type) {
	case int:
		if t < 0 {
			return errors.Wrap(err, "validate")
		}
	case float64:
		if t < 0.0 {
			return errors.Wrap(err, "validate")
		}
	}

	return nil
}

func (item *Item) CalculateTax() float64 {
	var tax float64
	rawTax := 0.125 * item.Price

	switch item.Type {
	case enum.Raw:
		tax = rawTax
	case enum.Manufactured:
		tax = rawTax + (0.02 * (rawTax + item.Price))
	case enum.Imported:
		importDuty := item.Price * 0.1
		finalCost := item.Price + importDuty
		if finalCost <= 100 {
			tax = importDuty + 5
		} else if finalCost <= 200 {
			tax = importDuty + 10
		} else {
			tax = importDuty + finalCost*0.05
		}
	}
	return tax
}
