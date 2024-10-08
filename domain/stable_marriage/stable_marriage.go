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
	engagements := newEngagement()

	// Every man courts first the most preferred woman.
	nextWomanToCourtIndex := make(map[string]int)
	for _, m := range input.Men {
		nextWomanToCourtIndex[m.ID] = 0
	}

	// Every woman starts the algorithm engaged with a very undesirable man
	for _, w := range input.Women {
		engagements.Engage("", w.ID)
	}

	for suitorIndex := 0; suitorIndex < instanceSize; {
		suitor := input.Men[suitorIndex]

		for suitor != nil {
			courted := suitor.Preference[nextWomanToCourtIndex[suitor.ID]]
			currentFiancee := getManEngagedWithWoman(engagements, input.Men, courted)
			if courted.Prefer(suitor, currentFiancee) {
				if currentFiancee != nil {
					engagements.Breakup(currentFiancee.ID)
				}
				engagements.Engage(suitor.ID, courted.ID)

				// The next suitor is the former fiancee.
				suitor = currentFiancee
			}

			// Courted refused the proposal. Court the next woman in the preference list.
			if suitor != nil {
				nextWomanToCourtIndex[suitor.ID] = nextWomanToCourtIndex[suitor.ID] + 1
			}
		}

		// Now a non-engaged man will court in the next round.
		suitorIndex = suitorIndex + 1
	}

	// Create the matching
	matching := &FindMatchingOutput{}
	for _, m := range input.Men {
		w := getWomanByID(input.Women, engagements.GetWomanID(m.ID))
		couple := &entities.Couple{
			Man:   *m,
			Woman: *w,
		}
		matching.Couples = append(matching.Couples, couple)
	}

	return matching, nil
}

func getManEngagedWithWoman(engagement engagement, men []*entities.Man, woman *entities.Woman) *entities.Man {
	if woman == nil {
		return nil
	}

	engagedManID := engagement.GetManID(woman.ID)
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
