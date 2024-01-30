package service

import (
	"fmt"

	"github.com/sahaj279/go_assignment/service/enum"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

type Item struct {
	Name     string        `gorm:"column:Name;type:varchar;size:255;" json:"name"`
	Price    float64       `gorm:"column:Price;type:float" json:"price"`
	Quantity int           `gorm:"column:Quantity:type:int;" json:"quantity"`
	Type     enum.ItemType `gorm:"column:Type;type:enum('raw','manufactured','imported');" json:"type"`
}

func NewItem(name string, itemType string, price float64, quantity int) (item Item, err error) {
	item = Item{Name: name, Price: price, Quantity: quantity}
	item.Type, err = enum.ItemTypeString(itemType)
	if err != nil {
		err = errors.Wrap(err, "newItem")
		return
	}

	err = validate(item)
	if err != nil {
		err = errors.Wrap(err, "newItem")
		return
	}

	return
}

func validate(item Item) error {
	return validation.ValidateStruct(&item,
		validation.Field(&item.Name, validation.Length(1, 0)),
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

func (item Item) Invoice() string {
	return fmt.Sprintf("Name: %s Price: %.2f Quantity: %d Type: %s Tax: %.2f Cost: %.2f", item.Name, item.Price, item.Quantity, item.Type.String(), item.getTax(), item.getEffectivePrice())
}

const (
	RAWItmTaxRate                       = 0.125
	ImportDuty                          = 0.100
	FirstImportDuty                     = 100
	SecondImportDuty                    = 200
	FirstImportDutySurchargeAmt         = 5
	SecondImportDutySurchargeAmt        = 10
	ExceedSecondImportDutySurchargeRate = 0.05
	ManufacturedItmTaxRate              = 0.125
	ManufacturedItmExtraTaxRate         = 0.02 // Extra =ItemCost +12.5% Item Cost
)

func (item Item) getEffectivePrice() float64 {
	var effectivePrice float64
	surcharge := 0.0
	tax := item.getTax()

	switch item.Type {
	case enum.Raw:
		effectivePrice = item.Price + tax + surcharge
	case enum.Manufactured:
		effectivePrice = item.Price + tax + surcharge
	case enum.Imported:
		priceTemp := item.Price + tax
		surcharge = item.importSurcharge(priceTemp)
		effectivePrice = priceTemp + surcharge
	}

	return effectivePrice
}

func (item Item) getTax() float64 {
	var tax float64
	switch item.Type {
	case enum.Raw:
		// raw: 12.5% of the item cost
		tax = RAWItmTaxRate * item.Price
	case enum.Manufactured:
		// manufactured: 12.5% of the item cost + 2% of (item cost + 12.5% of the item cost)
		tax = ManufacturedItmTaxRate*item.Price + ManufacturedItmExtraTaxRate*(item.Price+ManufacturedItmTaxRate*item.Price)
	case enum.Imported:
		// imported: 10% import duty on item cost
		tax = ImportDuty * item.Price
	}

	return tax
}

func (item Item) importSurcharge(price float64) float64 {
	if price <= FirstImportDuty {
		return FirstImportDutySurchargeAmt
	} else if price <= SecondImportDuty {
		return SecondImportDutySurchargeAmt
	}
	return price * ExceedSecondImportDutySurchargeRate
}
