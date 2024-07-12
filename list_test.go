package version

import (
	"sort"
	"testing"
)

func TestListSort(t *testing.T) {
	type TestCase struct {
		Input    List
		Expected List
	}

	testCases := []TestCase{
		{
			Input:    List{MustParse("1.1.1"), MustParse("1.0.1"), MustParse("1.1.0"), MustParse("1.0.0")},
			Expected: List{MustParse("1.0.0"), MustParse("1.0.1"), MustParse("1.1.0"), MustParse("1.1.1")},
		},
		{
			Input:    List{MustParse("4.0.0"), MustParse("3.0.0"), MustParse("2.0.0"), MustParse("1.0.0")},
			Expected: List{MustParse("1.0.0"), MustParse("2.0.0"), MustParse("3.0.0"), MustParse("4.0.0")},
		},
		{
			Input:    List{MustParse("2.0.4"), MustParse("1.2.4"), MustParse("1.2.3"), MustParse("1.3.1")},
			Expected: List{MustParse("1.2.3"), MustParse("1.2.4"), MustParse("1.3.1"), MustParse("2.0.4")},
		},
	}

	for i, testCase := range testCases {
		actual := List{}
		actual = append(actual, testCase.Input...)

		sort.Stable(actual)

		ok := true
		for j, v := range actual {
			expected := testCase.Expected[j]
			cmp := v.Compare(expected)
			if cmp != 0 {
				ok = false
				t.Errorf("test %d failed at position %d (expected %s, got %s)", i, j, expected, actual)
			}
		}

		if ok {
			t.Logf("test %d passed", i)
		}
	}
}

func TestList_Match(t *testing.T) {
	type TestCase struct {
		Input      List
		Constraint *Constraint
		Expected   List
	}

	testCases := []TestCase{
		{
			Input:    List{MustParse("1.0.0"), MustParse("1.1.0"), MustParse("2.0.2"), MustParse("3.4.5")},
			Expected: List{MustParse("1.0.0"), MustParse("1.1.0"), MustParse("2.0.2"), MustParse("3.4.5")},
		},
		{
			Input:      List{MustParse("1.0.0"), MustParse("1.1.0"), MustParse("2.0.2"), MustParse("3.4.5")},
			Constraint: &Constraint{Gt: MustParse("1.0.0")},
			Expected:   List{MustParse("1.1.0"), MustParse("2.0.2"), MustParse("3.4.5")},
		},
		{
			Input:      List{MustParse("1.0.0"), MustParse("1.1.0"), MustParse("2.0.2"), MustParse("3.4.5")},
			Constraint: &Constraint{Lt: MustParse("3.0.0")},
			Expected:   List{MustParse("1.0.0"), MustParse("1.1.0"), MustParse("2.0.2")},
		},
		{
			Input:      List{MustParse("1.0.0"), MustParse("1.1.0"), MustParse("2.0.2"), MustParse("3.4.5")},
			Constraint: &Constraint{Gte: MustParse("2.0.2")},
			Expected:   List{MustParse("2.0.2"), MustParse("3.4.5")},
		},
	}

	for i, testCase := range testCases {
		actual := testCase.Input.Match(testCase.Constraint)

		ok := true
		for j, v := range actual {
			expected := testCase.Expected[j]
			cmp := v.Compare(expected)
			if cmp != 0 {
				ok = false
				t.Errorf("test %d failed at position %d (expected %s, got %s)", i, j, expected, actual)
			}
		}

		if ok {
			t.Logf("test %d passed", i)
		}
	}
}
