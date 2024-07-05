package entities

type Woman struct {
	ID         string
	Partner    *Man
	Preference []*Man
}
