package stable_marriage

import "github.com/rafneves/two-sided-matching/domain/entities"

type StableMarriage struct {
}

type FindMatchingInput struct {
	Men   []*entities.Man
	Women []*entities.Woman
}

type FindMatchingOutput entities.Matching

func (m *StableMarriage) FindMatching(*FindMatchingInput) (*FindMatchingOutput, error) {
	return &FindMatchingOutput{}, nil
}
