//go:generate enumer -type=ItemType

package enum

type ItemType int

const (
	Raw ItemType = iota
	Manufactured
	Imported
)
