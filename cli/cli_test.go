package cli

import (
	"errors"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sahaj279/go-assignment/item"
	"github.com/sahaj279/go-assignment/item/mock_item"
)

type expectedRes struct {
	name     string
	price    float64
	quantity int
	itemType string
}

func TestEnterItem(t *testing.T) {
	tests := []struct {
		scenario string
		res      expectedRes
		req      *os.File
		err      error
	}{
		{
			scenario: "all item details provided",
			res: expectedRes{
				name:     "Bread",
				itemType: "raw",
				quantity: 2,
				price:    100,
			},
			req: setInput("Bread raw 100 2 \n"),
			err: nil,
		},
		{
			scenario: "all item details provided though invalid",
			res: expectedRes{
				name:     "Bread",
				itemType: "ram",
				quantity: -2,
				price:    -100,
			},
			req: setInput("Bread ram -100 -2 \n"),
			err: nil,
		},
		{
			scenario: "alphabets entered as price",
			req:      setInput("Bread ram one -2 \n"),
			err:      errors.New("enter item : price error: strconv.ParseFloat: parsing \"\": invalid syntax"),
		},
		{
			scenario: "someone just pressed enter after bread",
			req:      setInput("bread \n"),
			err:      errors.New("unexpected newline"),
		},
		{
			scenario: "someone just pressed enter and no input",
			req:      setInput(" \n"),
			err:      errors.New("unexpected newline"),
		},
		{
			scenario: "alphabets entered as quantity",
			req:      setInput(" bread exported -2 zero \n"),
			err:      errors.New("unexpected newline"),
		},
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	for _, tc := range tests {
		os.Stdin = tc.req
		name, itemType, price, quantity, err := enterItem()
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

func TestInit(t *testing.T) {
	mockItemSvc := mock_item.NewMockItemHandler(gomock.NewController(t))
	newCli := NewCli(mockItemSvc)
	tests := []struct {
		scenario            string
		cliInput            string
		mockEnterItemError  bool
		mockCreateItemError bool
		err                 error
	}{
		{
			scenario: "valid input with 1 item details",
			cliInput: "bread raw 40 4\nn\n",
			err:      nil,
		},
		{
			scenario: "valid input with 2 item details",
			cliInput: "bread raw 40 4\nn\nbread ram -3 6\ny\n",
			err:      nil,
		},
		{
			scenario: "valid input with 2 item details",
			cliInput: "bread ram 40 4\nn\nbread ram -3 6\ny\n",
			err:      nil,
		},
		{
			scenario:            "error in create item",
			cliInput:            " bread exported -2 -1 \nn\n",
			mockCreateItemError: true,
			err:                 errors.New("error in create item"),
		},
		{
			scenario: "error in enter more",
			cliInput: " bread raw 2 1 \nno\n",
			err:      errors.New("error in enter more"),
		},
		{
			scenario:           "error in enter item",
			cliInput:           " bread raw two 1 \nno\n",
			err:                errors.New("error in enter item"),
			mockEnterItemError: true,
		},
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	for _, tc := range tests {
		os.Stdin = setInput(tc.cliInput)
		if !tc.mockEnterItemError {
			if tc.mockCreateItemError {
				mockItemSvc.EXPECT().CreateItem(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(item.Item{}, errors.New("error in type"))
			} else {
				mockItemSvc.EXPECT().CreateItem(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(item.Item{}, nil)
			}
			if !tc.mockCreateItemError {
				mockItemSvc.EXPECT().CalculateTax(gomock.Any()).Return(0.0)
			}
		}

		err := newCli.Init()
		if err != nil && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if err == nil && tc.err != nil {
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
	err = os.WriteFile(tempFile.Name(), contentB, 0o600)
	if err != nil {
		// Close and remove the file if writing fails
		tempFile.Close()
		os.Remove(tempFile.Name())
		return nil
	}

	return tempFile
}
