package enum

//go:generate enumer -type=Course -json

type Course int

const (
	A Course = iota
	B
	C
	D
	E
	F
)
