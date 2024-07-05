package entities

type Man struct {
	ID         string
	Partner    *Woman
	Preference []*Woman
}
