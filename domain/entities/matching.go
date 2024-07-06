package entities

import "fmt"

type Couple struct {
	Man   Man
	Woman Woman
}

type Matching struct {
	Couples []*Couple
}

func (m *Matching) Print() {
	for _, couple := range m.Couples {
		fmt.Printf("Man %s is matched with woman %s.", couple.Man.ID, couple.Woman.ID)
	}
}
