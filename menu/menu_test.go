package menu

import (
	"errors"
	"testing"

	// "errors"
	"os"

	"github.com/golang/mock/gomock"
	"github.com/sahaj279/go_assignment/repository"
	"github.com/sahaj279/go_assignment/user"
)

func TestGetDetails(t *testing.T) {
	tests := []struct {
		scenario string
		input    string
		err      error
	}{
		{
			scenario: "all item details provided",
			input:    "sahaj \n21\nd and c\n52\n4\na\nb\nc\nd\ne\n",
			err:      nil,
		},
		{
			scenario: "invalid case: course count 2",
			input:    "sahaj \n21\nd and c\n52\n2\na\nb\nc\nd\ne\n",
			err:      errors.New("course should be inbetween 4 and 6"),
		},
		{
			scenario: "invalid case: error in roll number",
			input:    "sahaj \n2a1\nd and c\n52\n2\na\nb\nc\nd\ne\n",
			err:      errors.New("error in roll number"),
		},
		{
			scenario: "invalid case: error in name",
			input:    "\n2a1\nd and c\n52\n2\na\nb\nc\nd\ne\n",
			err:      errors.New("error in name"),
		},
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	for _, tc := range tests {
		os.Stdin = setInput(tc.input)
		_, _, _, _, _, err := GetDetails()
		if err != nil && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if err == nil && tc.err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}

func TestAddUser(t *testing.T) {
	tests := []struct {
		scenario string
		input    string
		mockErr  bool
		err      error
	}{
		{
			scenario: "all item details provided",
			input:    "sahaj \n21\nd and c\n52\n4\na\nb\nc\nd\ne\n",
			err:      nil,
		},
		{
			scenario: "invalid case: course count 2",
			input:    "sahaj \n21\nd and c\n52\n5\na\nb\nc\nd\ne\n",
			mockErr:  true,
			err:      errors.New("error in adding user to repo"),
		},
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	for _, tc := range tests {

		mockRepo := repository.NewMockSvc(gomock.NewController(t))
		menu := NewMenu(mockRepo)
		os.Stdin = setInput(tc.input)
		if tc.mockErr {
			mockRepo.EXPECT().Add(gomock.Any()).Return(errors.New("error in adding user to repo"))
		} else {
			mockRepo.EXPECT().Add(gomock.Any()).Return(nil)
		}
		err := menu.addUser()
		if err != nil && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if err == nil && tc.err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}

func TestDisplayUsers(t *testing.T) {
	tests := []struct {
		scenario string
		input    string
		mockList bool
		err      error
	}{
		{
			scenario: "sort by name in asc",
			input:    "name\n1\n",
			err:      nil,
			mockList: true,
		},
		{
			scenario: "sort in desc failed",
			input:    "name\n3\n",
			err:      errors.New("error in entering order,3 is not acceptable"),
		},
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	for _, tc := range tests {
		mockRepo := repository.NewMockSvc(gomock.NewController(t))
		menu := NewMenu(mockRepo)
		os.Stdin = setInput(tc.input)
		if tc.mockList {
			mockRepo.EXPECT().List(gomock.Any(), gomock.Any()).Return([]user.User{})
		}
		err := menu.displayUsers()
		if err != nil && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if err == nil && tc.err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		scenario   string
		input      string
		mockDelete bool
		mockErr    bool
		err        error
	}{
		{
			scenario:   "delete a valid roll number",
			input:      "2\n",
			err:        nil,
			mockDelete: true,
			mockErr:    false,
		},
		{
			scenario:   "delete an invalid roll number",
			input:      "2\n",
			err:        errors.New("roll number invalid"),
			mockDelete: true,
			mockErr:    true,
		},
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	for _, tc := range tests {
		mockRepo := repository.NewMockSvc(gomock.NewController(t))
		menu := NewMenu(mockRepo)
		os.Stdin = setInput(tc.input)
		if tc.mockDelete {
			if tc.mockErr {
				mockRepo.EXPECT().Delete(gomock.Any()).Return(errors.New("roll number invalid"))
			} else {
				mockRepo.EXPECT().Delete(gomock.Any()).Return(nil)
			}
		}
		err := menu.deleteUser()
		if err != nil && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if err == nil && tc.err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}

func TestSaveUser(t *testing.T) {
	tests := []struct {
		scenario string
		mockErr  bool
		err      error
	}{
		{
			scenario: "save successful without error",

			err: nil,

			mockErr: false,
		},
		{
			scenario: "error in saving",

			err: errors.New("error in saving"),

			mockErr: true,
		},
	}

	for _, tc := range tests {
		mockRepo := repository.NewMockSvc(gomock.NewController(t))
		menu := NewMenu(mockRepo)
		mockRepo.EXPECT().List(gomock.Any(), gomock.Any()).Return([]user.User{})

		if tc.mockErr {
			mockRepo.EXPECT().Save(gomock.Any()).Return(errors.New("error while saving"))
		} else {
			mockRepo.EXPECT().Save(gomock.Any()).Return(nil)
		}

		err := menu.saveUser()
		if err != nil && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if err == nil && tc.err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}

func TestInit(t *testing.T) {
	tests := []struct {
		scenario string
		input    string
		mock     bool
		mockErr  bool
		err      error
	}{
		{
			scenario: "initialization properly error due to scaner",
			input:    " 5\nn\n",
			err:      errors.New("error in new scanner"),
			mockErr:  false,
			mock:     true,
		},
		{
			scenario: "initialization failed error in load",
			input:    "5\nn\n",
			err:      errors.New("error in load"),
			mockErr:  true,
			mock:     true,
		},
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	for _, tc := range tests {
		mockRepo := repository.NewMockSvc(gomock.NewController(t))
		menu := NewMenu(mockRepo)
		os.Stdin = setInput(tc.input)
		mockRepo.EXPECT().Close().Return(nil)

		if tc.mock {
			if tc.mockErr {
				mockRepo.EXPECT().Load(gomock.Any()).Return(errors.New("load failed"))
			} else {
				mockRepo.EXPECT().Load(gomock.Any()).Return(nil)
			}
		}
		err := menu.Init()
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
