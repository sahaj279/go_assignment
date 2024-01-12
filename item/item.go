package item

import (
	"assignment1/item/enum"
)

type Item struct {
	Name     string
	Type     enum.ItemType
	Price    float64
	Quantity int
	SalesTax float64
}

func CreateItem(name string, itemType string, price float64, quantity int) Item {
	itemTypeEnum := enum.MapItemTypeToEnum(itemType)
	return Item{Name: name, Type: itemTypeEnum, Price: price, Quantity: quantity}
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
