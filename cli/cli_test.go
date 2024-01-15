package cli

import (
	"os"
	"testing"
)

type testStr struct {
	name     string
	price    float64
	quantity int
	itemType string
}

func TestGetItemFromCLI(t *testing.T) {
	testRaw := testStr{
		name:     "Bread",
		price:    100,
		quantity: 2,
		itemType: "raw",
	}
	testManufactured := testStr{
		name:     "Bread",
		price:    100,
		quantity: 2,
		itemType: "manufactured",
	}
	testImported := testStr{
		name:     "Bread",
		price:    100,
		quantity: 2,
		itemType: "imported",
	}

	testInvalidCase := testStr{
		name:     "Bread",
		price:    100,
		quantity: -1,
		itemType: "exported",
	}

	tests := []struct {
		scenario string
		str      testStr
		req      *os.File
		err      error
	}{{
		scenario: "all item details provided for raw",
		str:      testRaw,
		req:      setInput("Bread raw 100 2 \n"),
		err:      nil,
	}, {
		scenario: "all item details provided for manufactured",
		str:      testManufactured,
		req:      setInput("Bread manufactured 100 2 \n"),
		err:      nil,
	}, {
		scenario: "all item details provided for imported",
		str:      testImported,
		req:      setInput("Bread imported 100 2 \n"),
		err:      nil,
	}, {
		scenario: "invalid test case",
		str:      testInvalidCase,
		req:      setInput("Bread exported 100 -1 \n"),
		err:      nil,
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
		if name != tc.str.name || itemType != tc.str.itemType || price != tc.str.price || quantity != tc.str.quantity {
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
