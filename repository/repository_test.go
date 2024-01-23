package repository

import (
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		scenario string
		err      error
		fileName string
		repo     Repository
	}{
		{
			scenario: "file to load data from do not exists",
			repo:     Repository{},
			fileName: "empty.json",
		},
		{
			scenario: "file to load data from exists but is empty",
			repo:     Repository{},
			fileName: "empty.json",
		},
	}
	for _, tc := range tests {
		defer tc.repo.Close()
		err := tc.repo.Load(tc.fileName)
		if err != nil {
			t.Errorf("for scenario %s, got %v, expected %v", tc.scenario, err, tc.err)
		}
	}
}
