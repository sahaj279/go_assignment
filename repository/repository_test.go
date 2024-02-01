package repository

import (
	"testing"

	"github.com/pkg/errors"
	field "github.com/sahaj279/go_assignment/repository/data_field"
	"github.com/sahaj279/go_assignment/user"
	enum "github.com/sahaj279/go_assignment/user/course"
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
		{
			scenario: "file to load data from exists and is not empty",
			repo:     Repository{},
			fileName: "user_data_exists.json",
		},
		{
			scenario: "file to load data from exists and is not empty but data is corrupted",
			repo:     Repository{},
			fileName: "user_data_corrupted.json",
			err:      errors.New("error in decoding corrupted json file"),
		},
	}
	for _, tc := range tests {
		defer tc.repo.Close()
		err := tc.repo.Load(tc.fileName)
		if err != nil && tc.err == nil {
			t.Errorf("for scenario %s, got %v, expected %v", tc.scenario, err, tc.err)
		} else if err == nil && tc.err != nil {
			t.Errorf("for scenario %s, got %v, expected %v", tc.scenario, err, tc.err)
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		scenario string
		user     user.User
		err      error
		repo     Repository
	}{
		{
			scenario: "add user which already exists",
			user: user.User{
				Name:    "sahaj",
				Age:     21,
				Address: "d and c unico",
				RollNo:  43,
				Courses: []enum.Course{enum.A, enum.B, enum.C, enum.D},
			},
			repo: Repository{
				users: map[int]user.User{
					43: {
						Name:    "sahaj",
						Age:     21,
						Address: "d and c unico",
						RollNo:  43,
						Courses: []enum.Course{enum.A, enum.B, enum.C, enum.D},
					},
				},
			},
			err: errors.New("same roll number already exists"),
		}, {
			scenario: "add user with proper input",
			user: user.User{
				Name:    "sahaj",
				Age:     21,
				Address: "d and c unico",
				RollNo:  43,
				Courses: []enum.Course{enum.A, enum.B, enum.C, enum.D},
			},
			repo: Repository{
				users: map[int]user.User{
					42: {
						Name:    "sahaj",
						Age:     21,
						Address: "d and c unico",
						RollNo:  43,
						Courses: []enum.Course{enum.A, enum.B, enum.C, enum.D},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		err := tc.repo.Add(tc.user)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}

func TestList(t *testing.T) {
	tests := []struct {
		scenario  string
		dataField field.DataField
		ascOrder  bool
		repo      Repository
		exp       []user.User
	}{
		{
			scenario:  "list in asc order by name",
			dataField: field.Name,
			repo: Repository{
				users: map[int]user.User{
					43: {
						Name:    "sahaj",
						Age:     21,
						Address: "d and c unico",
						RollNo:  43,
						Courses: []enum.Course{enum.A, enum.B, enum.C, enum.D},
					},
					41: {
						Name:    "rohan",
						Age:     21,
						Address: "a and c naranpura",
						RollNo:  41,
						Courses: []enum.Course{enum.A, enum.B, enum.C, enum.D},
					},
				},
			},
			ascOrder: true,
			exp: []user.User{
				{
					Name:    "rohan",
					Age:     21,
					Address: "a and c naranpura",
					RollNo:  41,
					Courses: []enum.Course{enum.A, enum.B, enum.C, enum.D},
				},
				{
					Name:    "sahaj",
					Age:     21,
					Address: "d and c unico",
					RollNo:  43,
					Courses: []enum.Course{enum.A, enum.B, enum.C, enum.D},
				},
			},
		},
		{
			scenario:  "list in desc order by roll",
			dataField: field.RollNo,
			repo: Repository{
				users: map[int]user.User{
					43: {
						Name:    "sahaj",
						Age:     21,
						Address: "d and c unico",
						RollNo:  43,
						Courses: []enum.Course{enum.A, enum.B, enum.C, enum.D},
					},
					41: {
						Name:    "rohan",
						Age:     21,
						Address: "a and c naranpura",
						RollNo:  41,
						Courses: []enum.Course{enum.A, enum.B, enum.C, enum.D},
					},
				},
			},
			ascOrder: false,
			exp: []user.User{
				{
					Name:    "sahaj",
					Age:     21,
					Address: "d and c unico",
					RollNo:  43,
					Courses: []enum.Course{enum.A, enum.B, enum.C, enum.D},
				},
				{
					Name:    "rohan",
					Age:     21,
					Address: "a and c naranpura",
					RollNo:  41,
					Courses: []enum.Course{enum.A, enum.B, enum.C, enum.D},
				},
			},
		},
	}

	for _, tc := range tests {
		users := tc.repo.List(tc.dataField, tc.ascOrder)
		for i, user := range users {
			if user.RollNo != tc.exp[i].RollNo {
				t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, users, tc.exp)
				break
			}
		}

	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		scenario string
		rollNo   int
		err      error
		repo     Repository
	}{
		{
			scenario: "delete user which does not exist",
			rollNo:   21,
			repo: Repository{
				users: map[int]user.User{
					43: {
						Name:    "sahaj",
						Age:     21,
						Address: "d and c unico",
						RollNo:  43,
						Courses: []enum.Course{enum.A, enum.B, enum.C, enum.D},
					},
				},
			},
			err: errors.New("roll number does not exists"),
		}, {
			scenario: "delete user which exists",
			rollNo:   42,
			repo: Repository{
				users: map[int]user.User{
					42: {
						Name:    "sahaj",
						Age:     21,
						Address: "d and c unico",
						RollNo:  43,
						Courses: []enum.Course{enum.A, enum.B, enum.C, enum.D},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		err := tc.repo.Delete(tc.rollNo)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}

func TestSave(t *testing.T) {
	tests := []struct {
		scenario string
		users    []user.User
		err      error
		repo     Repository
	}{
		{
			scenario: "save users ",
			users: []user.User{
				{
					Name:    "rohan",
					Address: "a and c naranpura",
					RollNo:  41,
					Courses: []enum.Course{enum.A, enum.B, enum.C, enum.D},
				},
				{
					Name:    "sahaj",
					Age:     21,
					Address: "d and c unico",
					RollNo:  43,
					Courses: []enum.Course{enum.A, enum.B, enum.C, enum.D},
				},
			},
			repo: Repository{},
		},
	}

	for _, tc := range tests {
		tc.repo.Load("user_data_exists.json")
		err := tc.repo.Save(tc.users)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}
