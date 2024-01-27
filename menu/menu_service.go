package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	enum "github.com/sahaj279/go_assignment/repository/data_field"
	"github.com/sahaj279/go_assignment/user"
)

const (
	DataFilePath = "user_data.json"
	MinCourses   = 4
	TotalCourses = 6
	Accept       = "y"
	Reject       = "n"
	Ascending    = "1"
	Descending   = "2"
)

func (m *Menu) addUser() error {
	fmt.Printf("Add user details ")

	name, age, address, rollNo, courses, err := GetDetails()
	if err != nil {
		return errors.Wrap(err, "addUser")
	}

	// creating user object after it gets validated
	user, err := user.NewUser(name, age, address, rollNo, courses)
	if err != nil {
		return errors.Wrap(err, "addUser")
	}

	if err := m.repository.Add(user); err != nil {
		return errors.Wrap(err, "addUser")
	}

	fmt.Print("\nuser added successfully!\n")
	return nil
}

func GetDetails() (name string, age int, address string, rollNo int, courses []string, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Full Name: ")
	if scanner.Scan() {
		name = scanner.Text()
		name = strings.TrimSpace(name)
	}
	if err = scanner.Err(); err != nil {
		err = errors.Wrap(err, "getDetails")
		return
	}

	fmt.Printf("Age: ")
	if scanner.Scan() {
		age, err = strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			err = errors.Wrap(err, "getDetails")
			return
		}
	}
	if err = scanner.Err(); err != nil {
		err = errors.Wrap(err, "getDetails")
		return
	}

	fmt.Printf("Address: ")
	if scanner.Scan() {
		address = scanner.Text()
		address = strings.TrimSpace(address)
	}
	if err = scanner.Err(); err != nil {
		err = errors.Wrap(err, "getDetails")
		return
	}

	fmt.Printf("Roll No : ")
	if scanner.Scan() {
		rollNo, err = strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			err = errors.Wrap(err, "getDetails")
			return
		}
	}
	if err = scanner.Err(); err != nil {
		err = errors.Wrap(err, "getDetails")
		return
	}

	courses, err = getCourse(scanner)
	if err != nil {
		err = errors.Wrap(err, "getDetails")
		return
	}
	return
}

func getCourse(scanner *bufio.Scanner) ([]string, error) {
	fmt.Printf("Enter number of courses you want to enroll (at least %d) ", MinCourses)

	var courses []string
	var numCourses int = -1

	if scanner.Scan() {
		n, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			return []string{}, errors.Wrap(err, "getCourse")
		}
		numCourses = n
	}

	if err := scanner.Err(); err != nil {
		return []string{}, errors.Wrap(err, "getCourse")
	}

	if numCourses < MinCourses || numCourses > TotalCourses {
		err := fmt.Errorf("select at least %d and not more than %d while you entered %d", MinCourses, TotalCourses, numCourses)
		return []string{}, errors.Wrap(err, "getCourse")
	}

	for i := 1; i <= numCourses; i++ {
		fmt.Printf("Enter course - %d: (A,B,C,D,E,F)", i)
		var course string

		if scanner.Scan() {
			course = scanner.Text()
			course = strings.TrimSpace(course)
		}
		if err := scanner.Err(); err != nil {
			return []string{}, errors.Wrap(err, "getCourse")
		}

		courses = append(courses, course)
	}

	return courses, nil
}

func (m *Menu) displayUsers() error {
	fmt.Println("Display Users")

	field, ascOrder, err := getSortBy()
	if err != nil {
		return errors.Wrap(err, "displayUsers")
	}

	users := m.getAll(field, ascOrder)
	display(users)
	return nil
}

func getSortBy() (field enum.DataField, ascOrder bool, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Field Name to sort details on: ")
	var dataField string

	if scanner.Scan() {
		dataField = scanner.Text()
		dataField = strings.TrimSpace(dataField)
	}

	if err = scanner.Err(); err != nil {
		err = errors.Wrap(err, "getSortBy")
		return
	}

	field, err = enum.DataFieldString(dataField)
	if err != nil {
		err = errors.Wrap(err, "getSortBy")
		return
	}

	fmt.Print("\nIn which order should records be sorted\n[1] Ascending \n[2]Descending\n")
	var order string

	if scanner.Scan() {
		order = scanner.Text()
		order = strings.TrimSpace(order)
	}

	if err = scanner.Err(); err != nil {
		err = errors.Wrap(err, "getSortBy")
		return
	}

	if order != Ascending && order != Descending {
		err = errors.New("enter either 1 or 2")
		err = errors.Wrap(err, "getSortBy")
		return
	}

	if order == Ascending {
		ascOrder = true
	}

	return
}

func (m *Menu) getAll(field enum.DataField, ascOrder bool) (users []user.User) {
	users = m.repository.List(field, ascOrder)
	return users
}

func display(users []user.User) {
	fmt.Print("\n	Name	|	Age	|	Address	|	RollNo	|	Courses	|\n")
	fmt.Println()
	for _, user := range users {
		fmt.Println(user.String())
	}
}

func (m *Menu) deleteUser() error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter roll no to delete: ")
	var rollNo int

	if scanner.Scan() {
		r, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			return errors.Wrap(err, "deleteUser")
		}
		rollNo = r
	}
	if err := scanner.Err(); err != nil {
		return errors.Wrap(err, "deleteUser")
	}

	if err := m.repository.Delete(rollNo); err != nil {
		return errors.Wrap(err, "deleteUser")
	}

	fmt.Print("\nuser deleted successfully\n")
	return nil
}

func (m *Menu) saveUser() error {
	// saving data in ascending order of name
	users := m.repository.List(enum.Name, true)

	if err := m.repository.Save(users); err != nil {
		return errors.Wrap(err, "saveUser")
	}

	fmt.Println("saved successfully")
	return nil
}

func (m *Menu) confirmSave() error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Do you want to save the data(%s/%s)?", Accept, Reject)

	var userChoice string

	if scanner.Scan() {
		userChoice = scanner.Text()
		userChoice = strings.TrimSpace(userChoice)
	}

	if err := scanner.Err(); err != nil {
		return errors.Wrap(err, "confirmSave")
	}

	if err := validateConfirmation(userChoice); err != nil {
		return errors.Wrap(err, "confirmSave")
	}

	if userChoice == Accept {
		if err := m.saveUser(); err != nil {
			return errors.Wrap(err, "confirmSave")
		}
	}
	return nil
}

func validateConfirmation(userChoice string) error {
	if userChoice != Accept && userChoice != Reject {
		err := fmt.Errorf("%s: %s", "invalid choice", userChoice)

		return errors.Wrap(err, "confirmSave")
	}

	return nil
}
