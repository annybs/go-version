package version

import (
	"testing"
)

func TestVersion_String(t *testing.T) {
	type TestCase struct {
		Expected string  // Expected
		Input    Version // Input
	}

	testCases := []TestCase{
		{Expected: "0.0.0", Input: Version{}},
		{Expected: "1.0.0", Input: Version{Major: 1}},
		{Expected: "0.1.0", Input: Version{Minor: 1}},
		{Expected: "0.0.1", Input: Version{Patch: 1}},
		{Expected: "10.20.30", Input: Version{Major: 10, Minor: 20, Patch: 30}},
		{Expected: "27.31.15", Input: Version{Major: 27, Minor: 31, Patch: 15}},
		{Expected: "v1.2.3", Input: Version{Major: 1, Minor: 2, Patch: 3, Text: "v1.2.3"}},
		{Expected: "v1.2.3", Input: Version{Text: "v1.2.3"}},
		{Expected: "v1.2.0a", Input: Version{Major: 1, Minor: 2, Extension: "a", Text: "v1.2.0a"}},
		{Expected: "v1-alpha2", Input: Version{Major: 1, Extension: "-alpha2", Text: "v1-alpha2"}},
	}

	for i, testCase := range testCases {
		actual := testCase.Input.String()

		if actual != testCase.Expected {
			t.Errorf("test %d failed (expected %s, actual %s)", i, testCase.Expected, actual)
		} else {
			t.Logf("test %d passed with %s\n", i, actual)
		}
	}
}

func TestVersion_Compare(t *testing.T) {
	type TestCase struct {
		A        *Version
		B        *Version
		Expected int
	}

	testCases := []TestCase{
		{A: MustParse("0.0.1"), B: MustParse("0.0.1"), Expected: 0},
		{A: MustParse("0.1.0"), B: MustParse("0.1.0"), Expected: 0},
		{A: MustParse("1.0.0"), B: MustParse("1.0.0"), Expected: 0},
		{A: MustParse("1.0.0"), B: MustParse("0.1.0"), Expected: 1},
		{A: MustParse("1.0.0"), B: MustParse("0.0.1"), Expected: 1},
		{A: MustParse("1.0.0"), B: MustParse("1.1.0"), Expected: -1},
		{A: MustParse("1.0.0"), B: MustParse("1.0.1"), Expected: -1},
		{A: MustParse("1.0.0"), B: MustParse("2.0.0"), Expected: -1},
		{A: MustParse("1.1.0"), B: MustParse("1.2.0"), Expected: -1},
		{A: MustParse("1.1.1"), B: MustParse("1.2.0"), Expected: -1},
		{A: MustParse("1.20.0"), B: MustParse("1.2.0"), Expected: 1},
		{A: MustParse("1.20.0"), B: MustParse("1.2.20"), Expected: 1},
		{A: MustParse("1.20.0"), B: MustParse("1.20.1"), Expected: -1},
	}

	for i, testCase := range testCases {
		actual := testCase.A.Compare(testCase.B)
		if actual != testCase.Expected {
			t.Errorf("test %d failed (expected %d, actual %d)", i, testCase.Expected, actual)
		} else {
			t.Logf("test %d passed with %d", i, actual)
		}
	}
}

func TestVersion_Less(t *testing.T) {
	type TestCase struct {
		A        *Version
		B        *Version
		Expected bool
	}

	testCases := []TestCase{
		{A: MustParse("0.0.1"), B: MustParse("0.0.1"), Expected: false},
		{A: MustParse("0.1.0"), B: MustParse("0.1.0"), Expected: false},
		{A: MustParse("1.0.0"), B: MustParse("1.0.0"), Expected: false},
		{A: MustParse("1.0.0"), B: MustParse("0.1.0"), Expected: false},
		{A: MustParse("1.0.0"), B: MustParse("0.0.1"), Expected: false},
		{A: MustParse("1.0.0"), B: MustParse("1.1.0"), Expected: true},
		{A: MustParse("1.0.0"), B: MustParse("1.0.1"), Expected: true},
		{A: MustParse("1.0.0"), B: MustParse("2.0.0"), Expected: true},
		{A: MustParse("1.1.0"), B: MustParse("1.2.0"), Expected: true},
		{A: MustParse("1.1.1"), B: MustParse("1.2.0"), Expected: true},
		{A: MustParse("1.20.0"), B: MustParse("1.2.0"), Expected: false},
		{A: MustParse("1.20.0"), B: MustParse("1.2.20"), Expected: false},
		{A: MustParse("1.20.0"), B: MustParse("1.20.1"), Expected: true},
	}

	for i, testCase := range testCases {
		actual := testCase.A.Less(testCase.B)
		if actual != testCase.Expected {
			t.Errorf("test %d failed (expected %v, actual %v)", i, testCase.Expected, actual)
		} else {
			t.Logf("test %d passed with %v", i, actual)
		}
	}
}
