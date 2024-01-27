package user

import (
	"errors"
	"testing"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		scenario string
		name     string
		age      int
		course   []string
		rollNo   int
		address  string
		err      error
	}{
		{
			scenario: "a valid test case",
			name:     "sahaj",
			age:      21,
			course:   []string{"a", "b", "c", "d", "e"},
			rollNo:   53,
			address:  "d and c",
			err:      nil,
		},
		{
			scenario: "an invalid test case with repeated courses",
			name:     "sahaj",
			age:      21,
			course:   []string{"a", "c", "c", "d", "e"},
			rollNo:   53,
			address:  "d and c",
			err:      errors.New("course should not be same"),
		},
		{
			scenario: "an invalid test case with invalid course name",
			name:     "sahaj",
			age:      21,
			course:   []string{"ab", "c", "c", "d", "e"},
			rollNo:   53,
			address:  "d and c",
			err:      errors.New("course should valid"),
		},
		{
			scenario: "an invalid test case with negative rollnumber",
			name:     "sahaj",
			age:      -21,
			course:   []string{"a", "c", "c", "d", "e"},
			rollNo:   -53,
			address:  "d and c",
			err:      errors.New("roll number should br positive"),
		},
	}

	for _, tc := range tests {
		_, err := NewUser(tc.name, tc.age, tc.address, tc.rollNo, tc.course)
		if err != nil && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if err == nil && tc.err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}

	}
}
