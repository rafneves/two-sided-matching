package stable_marriage_test

import (
	"github.com/rafneves/two-sided-matching/domain/entities"
	"github.com/rafneves/two-sided-matching/domain/stable_marriage"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStableMarriage_FindMatching_KnuthExample1(t *testing.T) {
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
				Men:   []*entities.Man{anatole, barnabe, camille, dominique},
				Women: []*entities.Woman{antoniette, brigitte, cunegonde, donatienne},
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

func TestStableMarriage_FindMatching_KnuthExample2(t *testing.T) {
	// Initialization of People
	// Men
	manA := &entities.Man{ID: "manA"}
	manB := &entities.Man{ID: "manB"}
	manC := &entities.Man{ID: "manC"}

	// Women
	womanA := &entities.Woman{ID: "womanA"}
	womanB := &entities.Woman{ID: "womanB"}
	womanC := &entities.Woman{ID: "womanC"}

	// Men Preference
	manA.Preference = []*entities.Woman{womanB, womanA, womanC}
	// Arbitrary list
	manB.Preference = []*entities.Woman{womanB, womanA, womanC}
	manC.Preference = []*entities.Woman{womanA, womanB, womanC}

	// Women Preference
	womanA.Preference = []*entities.Man{manA, manC, manB}
	womanB.Preference = []*entities.Man{manC, manA, manB}
	// Arbitraty List
	womanC.Preference = []*entities.Man{manC, manB, manA}

	cases := []struct {
		name           string
		input          *stable_marriage.FindMatchingInput
		expectedResult *stable_marriage.FindMatchingOutput
		err            error
	}{
		{
			name: "sucesfull man proposing matching for Knuth, Chap 1, Example 2",
			input: &stable_marriage.FindMatchingInput{
				Men:   []*entities.Man{manA, manB, manC},
				Women: []*entities.Woman{womanA, womanB, womanC},
			},
			expectedResult: &stable_marriage.FindMatchingOutput{
				Couples: []*entities.Couple{
					{
						Man:   *manA,
						Woman: *womanB,
					},
					{
						Man:   *manB,
						Woman: *womanC,
					},
					{
						Man:   *manC,
						Woman: *womanA,
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

func TestStableMarriage_FindMatching_KnuthExample3(t *testing.T) {
	// Initialization of People
	// Men
	manA := &entities.Man{ID: "manA"}
	manB := &entities.Man{ID: "manB"}
	manC := &entities.Man{ID: "manC"}
	manD := &entities.Man{ID: "manD"}
	manE := &entities.Man{ID: "manE"}

	// Women
	womanA := &entities.Woman{ID: "womanA"}
	womanB := &entities.Woman{ID: "womanB"}
	womanC := &entities.Woman{ID: "womanC"}
	womanD := &entities.Woman{ID: "womanD"}
	womanE := &entities.Woman{ID: "womanE"}

	// Men Preference (Circular permutation)
	manA.Preference = []*entities.Woman{womanA, womanB, womanC, womanD, womanE}
	manB.Preference = []*entities.Woman{womanB, womanC, womanD, womanE, womanA}
	manC.Preference = []*entities.Woman{womanC, womanD, womanE, womanA, womanB}
	manD.Preference = []*entities.Woman{womanD, womanE, womanA, womanB, womanC}
	manE.Preference = []*entities.Woman{womanE, womanA, womanB, womanC, womanD}

	// Women Preference
	womanA.Preference = []*entities.Man{manB, manC, manD, manE, manA}
	womanB.Preference = []*entities.Man{manC, manD, manE, manA, manB}
	womanC.Preference = []*entities.Man{manD, manE, manA, manB, manC}
	womanD.Preference = []*entities.Man{manE, manA, manB, manC, manD}
	womanE.Preference = []*entities.Man{manA, manB, manC, manD, manE}

	cases := []struct {
		name           string
		input          *stable_marriage.FindMatchingInput
		expectedResult *stable_marriage.FindMatchingOutput
		err            error
	}{
		{
			name: "sucesfull man proposing matching for Knuth, Chap 1, Example 3",
			input: &stable_marriage.FindMatchingInput{
				Men:   []*entities.Man{manA, manB, manC, manD, manE},
				Women: []*entities.Woman{womanA, womanB, womanC, womanD, womanE},
			},
			expectedResult: &stable_marriage.FindMatchingOutput{
				Couples: []*entities.Couple{
					{
						Man:   *manA,
						Woman: *womanA,
					},
					{
						Man:   *manB,
						Woman: *womanB,
					},
					{
						Man:   *manC,
						Woman: *womanC,
					},
					{
						Man:   *manD,
						Woman: *womanD,
					},
					{
						Man:   *manE,
						Woman: *womanE,
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
