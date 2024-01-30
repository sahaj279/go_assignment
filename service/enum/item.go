//go:generate enumer -type=ItemType

package enum

import (
	"database/sql/driver"

	"github.com/pkg/errors"
)

type ItemType int

const (
	Raw ItemType = iota
	Manufactured
	Imported
)

// Scan and Value function for gorm to be able to fetch itemType enum from database and put it in database as well

func (i *ItemType) Scan(value interface{}) error {
	expr, _ := value.([]byte)
	var err error
	*i, err = ItemTypeString(string(expr))
	if err != nil {
		return errors.Wrap(err, "invalid item type present in database")
	}

	return nil
}

func (i ItemType) Value() (driver.Value, error) {
	return i.String(), nil
}
