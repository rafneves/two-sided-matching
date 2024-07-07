package stable_marriage

import (
	"errors"
	"github.com/rafneves/two-sided-matching/domain/entities"
)

type StableMarriage struct {
}

type FindMatchingInput struct {
	Men   []*entities.Man
	Women []*entities.Woman
}

type FindMatchingOutput entities.Matching

func (m *StableMarriage) FindMatching(input *FindMatchingInput) (*FindMatchingOutput, error) {
	if len(input.Men) != len(input.Women) {
		return nil, errors.New("the number of men don't match the numbe of women")
	}
	instanceSize := len(input.Men)

	engagements := make(map[string]string)

	nextCourtshipIndex := make(map[string]int)
	for _, m := range input.Men {
		nextCourtshipIndex[m.ID] = 0
	}

	suitorIndex := 0
	for len(engagements) < instanceSize {
		suitor := input.Men[suitorIndex]

		for suitor != nil {
			courted := suitor.Preference[nextCourtshipIndex[suitor.ID]]
			if suitorIsPreferredToCurrentFiancee(engagements, *courted, *suitor) {
				// Break up the current couple with the courted woman and make the former fiancee the next suitor.
				nextSuitor := getManByID(input.Men, getManIDByFiancee(engagements, *courted))
				if nextSuitor != nil {
					delete(engagements, nextSuitor.ID)
				}
				engagements[suitor.ID] = courted.ID
				suitor = nextSuitor
			} else {
			}

			// Remove the courted woman from suitor list as she already has someone preferred to him.
			if suitor != nil {
				nextCourtshipIndex[suitor.ID] = nextCourtshipIndex[suitor.ID] + 1
			}
		}

		// The next man will be the suitor
		suitorIndex = suitorIndex + 1
	}

	// Create the matching
	matching := &FindMatchingOutput{}
	for _, m := range input.Men {
		couple := &entities.Couple{
			Man:   *m,
			Woman: *getWomanByID(input.Women, engagements[m.ID]),
		}
		matching.Couples = append(matching.Couples, couple)
	}

	return matching, nil
}

func suitorIsPreferredToCurrentFiancee(engagements map[string]string, woman entities.Woman, man entities.Man) bool {

	for _, m := range woman.Preference {
		if m.ID == getManIDByFiancee(engagements, woman) {
			return false
		}

		if man.ID == m.ID {
			return true
		}
	}

	// Any partner is always preferred to being single
	return true
}

func getManIDByFiancee(engagements map[string]string, woman entities.Woman) string {
	for manID, womanID := range engagements {
		if woman.ID == womanID {
			return manID
		}
	}

	return ""
}

func getManByID(men []*entities.Man, ID string) *entities.Man {
	for _, m := range men {
		if m.ID == ID {
			return m
		}
	}
	return nil
}

func getWomanByID(women []*entities.Woman, ID string) *entities.Woman {
	for _, w := range women {
		if w.ID == ID {
			return w
		}
	}
	return nil
}
