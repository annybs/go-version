package version

import (
	"errors"
	"testing"
)

func TestVersion_Parse(t *testing.T) {
	type TestCase struct {
		Input    string
		Expected Version
		Err      error
	}

	testCases := []TestCase{
		{Input: "0.0.0", Expected: Version{Text: "0.0.0"}},
		{Input: "1.0.0", Expected: Version{Major: 1, Text: "1.0.0"}},
		{Input: "0.1.0", Expected: Version{Minor: 1, Text: "0.1.0"}},
		{Input: "0.0.1", Expected: Version{Patch: 1, Text: "0.0.1"}},
		{Input: "10.20.30", Expected: Version{Major: 10, Minor: 20, Patch: 30, Text: "10.20.30"}},
		{Input: "27.31.15", Expected: Version{Major: 27, Minor: 31, Patch: 15, Text: "27.31.15"}},
		{Input: "v1.2.3", Expected: Version{Major: 1, Minor: 2, Patch: 3, Text: "v1.2.3"}},
		{Input: "v1", Expected: Version{Major: 1, Text: "v1"}},
		{Input: "v2.31", Expected: Version{Major: 2, Minor: 31, Text: "v2.31"}},
		{Input: "v1.2.0a", Expected: Version{Major: 1, Minor: 2, Extension: "a", Text: "v1.2.0a"}},
		{Input: "v1.2a", Expected: Version{Major: 1, Minor: 2, Extension: "a", Text: "v1.2a"}},
		{Input: "v1-alpha2", Expected: Version{Major: 1, Extension: "-alpha2", Text: "v1-alpha2"}},
		{Input: "invalid version", Err: ErrInvalidVersion},
		{Input: "v.01", Err: ErrInvalidVersion},
		{Input: "v-any", Err: ErrInvalidVersion},
	}

	for i, testCase := range testCases {
		actual, err := Parse(testCase.Input)

		if testCase.Err != nil {
			if err == nil {
				t.Errorf("test %d failed (expected error %s, actual nil)", i, testCase.Err)
			} else if !errors.Is(err, testCase.Err) {
				t.Errorf("test %d failed (expected error %s, actual error %s)", i, testCase.Err, err)
			} else {
				t.Logf("test %d passed with error %s for %q\n", i, err, testCase.Input)
			}
		} else if err != nil {
			t.Errorf("test %d failed (expected error nil, actual error %s)", i, err)
		} else if actual.Major != testCase.Expected.Major || actual.Minor != testCase.Expected.Minor || actual.Patch != testCase.Expected.Patch || actual.Text != testCase.Expected.Text {
			t.Errorf("test %d failed (expected %v, actual %v)", i, testCase.Expected, actual)
		} else {
			t.Logf("test %d passed with %v\n", i, actual)
		}
	}
}
