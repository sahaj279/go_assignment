package user

import (
	enum "assignment2/user/course"
	"encoding/json"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	Name    string
	Age     int
	Address string
	RollNo  int
	Courses []enum.Course // set of courses
}

func NewUser(name string, age int, address string, rollNo int, courses []string) (user User, err error) {
	// creating user object
	user.Name = name
	user.Age = age
	user.Address = address
	user.RollNo = rollNo

	var c []enum.Course
	c, err = getCourses(courses)
	if err != nil {
		return
	}

	user.Courses = c

	// validating it
	if err = user.validate(); err != nil {
		return
	}

	// only returning if validation successful
	return
}

func getCourses(courses []string) ([]enum.Course, error) {
	var courseEnum []enum.Course
	for _, c := range courses {
		course, err := enum.CourseString(c)
		if err != nil {
			return []enum.Course{}, err
		}
		courseEnum = append(courseEnum, course)
	}

	return courseEnum, nil
}

func (user User) validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.Required),
		validation.Field(&user.Age, validation.Required, validation.By(CheckPositive)),
		validation.Field(&user.Address, validation.Required),
		validation.Field(&user.RollNo, validation.Required, validation.By(CheckPositive)),
		validation.Field(&user.Courses, validation.Required, validation.By(CheckUnique)),
	)
}

func CheckUnique(value interface{}) error {
	courses := value.([]enum.Course)
	set := map[enum.Course]struct{}{}

	for _, c := range courses {
		_, exists := set[c]
		if exists {
			return errors.New("should be unique")
		}
		set[c] = struct{}{}
	}
	return nil
}

func CheckPositive(value interface{}) error {
	val := value.(int)
	if val <= 0 {
		return errors.New("must be positive")
	}
	return nil
}

func (user User) String() string {
	return fmt.Sprintf("	%s	|	%d	|	%s	|	%d	|	%s|\n", user.Name, user.Age, user.Address, user.RollNo, courseString(user.Courses))
}

func courseString(course []enum.Course) []string {
	var courses []string
	for _, c := range course {
		courses = append(courses, c.String())
	}

	return courses
}

func EncodeUsers(users []User) ([]byte, error) {
	userB, err := json.Marshal(users)
	if err != nil {
		return []byte{}, err
	}

	return userB, nil
}

func DecodeUsers(userB []byte) ([]User, error) {
	var users []User
	if err := json.Unmarshal(userB, &users); err != nil {
		return []User{}, err
	}

	return users, nil

}
