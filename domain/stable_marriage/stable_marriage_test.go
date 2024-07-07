package stable_marriage_test

import (
	"github.com/rafneves/two-sided-matching/domain/entities"
	"github.com/rafneves/two-sided-matching/domain/stable_marriage"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStableMarriage_FindMatching(t *testing.T) {
	// Initialization of People
	// Men
	anatole := &entities.Man{ID: "anatole"}
	barnabe := &entities.Man{ID: "barnabe"}
	camille := &entities.Man{ID: "camille"}
	dominique := &entities.Man{ID: "dominique"}

	// Women
	antoniette := &entities.Woman{ID: "antoniette"}
	brigitte := &entities.Woman{ID: "brigitte"}
	cunegonde := &entities.Woman{ID: "cunegonde"}
	donatienne := &entities.Woman{ID: "donatienne"}

	// Men Preference
	anatole.Preference = []*entities.Woman{cunegonde, brigitte, donatienne, antoniette}
	barnabe.Preference = []*entities.Woman{brigitte, antoniette, cunegonde, donatienne}
	camille.Preference = []*entities.Woman{brigitte, donatienne, antoniette, cunegonde}
	dominique.Preference = []*entities.Woman{cunegonde, antoniette, donatienne, brigitte}

	// Women Preference
	antoniette.Preference = []*entities.Man{anatole, barnabe, dominique, camille}
	brigitte.Preference = []*entities.Man{camille, anatole, dominique, barnabe}
	cunegonde.Preference = []*entities.Man{camille, barnabe, dominique, anatole}
	donatienne.Preference = []*entities.Man{barnabe, anatole, camille, dominique}

	cases := []struct {
		name           string
		input          *stable_marriage.FindMatchingInput
		expectedResult *stable_marriage.FindMatchingOutput
		err            error
	}{
		{
			name: "sucesfull man proposing matching for Knuth, Chap 1, Example 1",
			input: &stable_marriage.FindMatchingInput{
				[]*entities.Man{anatole, barnabe, camille, dominique},
				[]*entities.Woman{antoniette, brigitte, cunegonde, donatienne},
			},
			expectedResult: &stable_marriage.FindMatchingOutput{
				Couples: []*entities.Couple{
					{
						Man:   *anatole,
						Woman: *donatienne,
					},
					{
						Man:   *barnabe,
						Woman: *antoniette,
					},
					{
						Man:   *camille,
						Woman: *brigitte,
					},
					{
						Man:   *dominique,
						Woman: *cunegonde,
					},
				},
			},
			err: nil,
		},
	}

	for _, c := range cases {
		matching := stable_marriage.StableMarriage{}
		result, err := matching.FindMatching(c.input)

		require.Equal(t, c.expectedResult, result)
		require.Equal(t, c.err, err)
	}
}
