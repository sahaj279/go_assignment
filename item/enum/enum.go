package enum

type ItemType int

const (
	Raw ItemType = iota
	Manufactured
	Imported
)

func MapItemTypeToEnum(itemType string) ItemType {
	switch itemType {
	case "raw":
		return Raw
	case "manufactured":
		return Manufactured
	default:
		return Imported
	}
}
