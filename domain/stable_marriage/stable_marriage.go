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

	nextWomanToCourtIndex := make(map[string]int)
	for _, m := range input.Men {
		nextWomanToCourtIndex[m.ID] = 0
	}

	suitorIndex := 0
	for len(engagements) < instanceSize {
		suitor := input.Men[suitorIndex]

		for suitor != nil {
			courted := suitor.Preference[nextWomanToCourtIndex[suitor.ID]]
			currentFiancee := getManEngagedWithWoman(engagements, input.Men, courted)
			if courted.Prefer(suitor, currentFiancee) {
				// Break up the current couple with the courted woman and make the former fiancee the next suitor.
				if currentFiancee != nil {
					delete(engagements, currentFiancee.ID)
				}
				engagements[suitor.ID] = courted.ID
				suitor = currentFiancee
			}

			// Courted refused the proposal. Try the next.
			if suitor != nil {
				nextWomanToCourtIndex[suitor.ID] = nextWomanToCourtIndex[suitor.ID] + 1
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

func getManEngagedWithWoman(engagements map[string]string, men []*entities.Man, woman *entities.Woman) *entities.Man {
	if woman == nil {
		return nil
	}

	engagedManID := ""
	for manID, womanID := range engagements {
		if woman.ID == womanID {
			engagedManID = manID
			break
		}
	}

	for _, m := range men {
		if m.ID == engagedManID {
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
