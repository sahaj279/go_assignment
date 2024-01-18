package menu

import (
	repo "assignment2/repository"
	"assignment2/user"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func addUser(repository repo.Repository) error {
	fmt.Printf("Add user details ")

	// getting user details from cli
	name, age, address, rollNo, courses, err := getDetailsFromCLI()
	if err != nil {
		return err
	}

	// creating user object after it gets validated
	user, err := user.NewUser(name, age, address, rollNo, courses)
	if err != nil {
		return err
	}

	// adding user to repository
	if err := repository.Add(user); err != nil {
		return err
	}

	fmt.Print("\nuser added successfully!\n")
	return nil
}

func getDetailsFromCLI() (name string, age int, address string, rollNo int, courses []string, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Full Name: ")
	if scanner.Scan() {
		name = scanner.Text()
		name = strings.TrimSpace(name)
	}
	if err = scanner.Err(); err != nil {
		return
	}

	fmt.Printf("Age: ")
	if scanner.Scan() {
		age, err = strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			return
		}
	}
	if err = scanner.Err(); err != nil {
		return
	}

	fmt.Printf("Address: ")
	if scanner.Scan() {
		address = scanner.Text()
		address = strings.TrimSpace(address)
	}
	if err = scanner.Err(); err != nil {
		return
	}

	fmt.Printf("Roll No : ")
	if scanner.Scan() {
		rollNo, err = strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			return
		}
	}
	if err = scanner.Err(); err != nil {
		return
	}

	courses, err = getCourse()
	if err != nil {
		return
	}
	return
}

func getCourse() ([]string, error) {
	// getting courses which can be 4 and maximum 6 and should be unique
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter number of courses you want to enrol (at least %d) ", MinCourses)

	var courses []string
	var numCourses int

	// asking for number of courses one wants to apply for
	if scanner.Scan() {
		n, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			return []string{}, err
		}
		numCourses = n
	}
	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	if numCourses < MinCourses || numCourses > TotalCourses {
		err := fmt.Errorf("select at least %d and not more than %d", MinCourses, TotalCourses)
		return []string{}, err
	}

	// entering the courses
	for i := 1; i <= numCourses; i++ {
		fmt.Printf("Enter course - %d: (A,B,C,D,E,F)", i)
		var course string

		if scanner.Scan() {
			course = scanner.Text()
			course = strings.TrimSpace(course)
		}
		if err := scanner.Err(); err != nil {
			return []string{}, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}

func displayUsers(repository *repo.Repository) error {
	// display users present in memory
	fmt.Println("Display Users")

	// get users in a sorted order
	users, err := getAll(repository)
	if err != nil {
		return err
	}
	// print details
	display(users)
	return nil
}

func getAll(repository *repo.Repository) (users []user.User, err error) {
	// enter field name to sort the users
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Field Name to sort details on: ")
	var field string

	if scanner.Scan() {
		field = scanner.Text()
		field = strings.TrimSpace(field)
	}
	if err = scanner.Err(); err != nil {
		return
	}

	fmt.Print("\nIn which order should records be sorted\n[1] Ascending \n[2]Descending\n")
	var order string

	if scanner.Scan() {
		order = scanner.Text()
		order = strings.TrimSpace(order)
	}
	if err = scanner.Err(); err != nil {
		return
	}
	if order != Ascending && order != Descending {
		err = errors.New("enter either 1 or 2")
		return
	}

	var ASCOrder bool = true
	if order == Ascending {
		ASCOrder = false
	}

	users, err = repository.List(field, ASCOrder)
	if err != nil {
		return []user.User{}, err
	}

	return users, nil
}

func display(users []user.User) {
	fmt.Print("\n	Name	|	Age	|	Address	|	RollNo	|	Courses	|\n")
	fmt.Println()
	for _, user := range users {
		fmt.Println(user.String())
	}
}

func deleteByRollNo(repository repo.Repository) error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter roll no to delete: ")

	var rollNo int

	if scanner.Scan() {
		r, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {

			return err
		}
		rollNo = r
	}
	if err := scanner.Err(); err != nil {

		return err
	}

	if err := repository.Delete(rollNo); err != nil {
		return err
	}

	fmt.Print("\nuser deleted successfully\n")

	return nil
}

func save(repository repo.Repository) error {
	//saving data in ascending order of name
	users, err := repository.List("name", true)
	if err != nil {
		return err
	}

	if err = repository.Save(users); err != nil {
		return err
	}

	fmt.Println("saved successfully")

	return nil
}

func confirmSave(repository repo.Repository) error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Do you want to save the data(%s/%s)?", Accept, Reject)

	var userChoice string

	if scanner.Scan() {
		userChoice = scanner.Text()
		userChoice = strings.TrimSpace(userChoice)
	}
	if err := scanner.Err(); err != nil {

		return err
	}

	if err := validateConfirmation(userChoice); err != nil {
		return err
	}

	if userChoice == Accept {
		if err := save(repository); err != nil {
			return err
		}
	}
	return nil
}

func validateConfirmation(userChoice string) error {
	if userChoice != Accept && userChoice != Reject {
		err := fmt.Errorf("%s: %s", "invalid choice", userChoice)

		return err
	}

	return nil
}
