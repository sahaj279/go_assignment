package enum

type DataField int

//go:generate enumer -type=DataField -json
const (
	Name DataField = iota
	Age
	RollNo
	Address
)
