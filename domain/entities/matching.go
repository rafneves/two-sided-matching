package entities

type Couple struct {
	Man   Man
	Woman Woman
}

type Matching struct {
	Men     []*Man
	Women   []*Woman
	Couples []*Couple
}
