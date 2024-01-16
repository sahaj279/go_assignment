package cli

import (
	"errors"
	"os"
	"testing"
)

type expectedRes struct {
	name     string
	price    float64
	quantity int
	itemType string
}

func TestGetItemFromCLI(t *testing.T) {

	tests := []struct {
		scenario string
		res      expectedRes
		req      *os.File
		err      error
	}{{
		scenario: "all item details provided",
		res: expectedRes{
			name:     "Bread",
			itemType: "raw",
			quantity: 2,
			price:    100,
		},
		req: setInput("Bread raw 100 2 \n"),
		err: nil,
	}, {
		scenario: "someone just pressed enter after bread",
		req:      setInput("bread \n"),
		err:      errors.New("unexpected newline"),
	},
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	for _, tc := range tests {
		os.Stdin = tc.req
		name, itemType, price, quantity, err := getItemFromCLI()
		if err != nil && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if err == nil && tc.err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}

		if (name != tc.res.name || itemType != tc.res.itemType || price != tc.res.price || quantity != tc.res.quantity) && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}

}

func setInput(test string) *os.File {
	contentB := []byte(test)

	tempFile, err := os.CreateTemp("", "temp_file")
	if err != nil {
		return nil
	}

	// Write the content to the file
	err = os.WriteFile(tempFile.Name(), contentB, 0600)
	if err != nil {
		// Close and remove the file if writing fails
		tempFile.Close()
		os.Remove(tempFile.Name())
		return nil
	}

	return tempFile
}
