package user

import (
	"fmt"

	"github.com/pkg/errors"

	enum "github.com/sahaj279/go_assignment/user/course"

	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	Name    string        `json:"name,omitempty"`
	Age     int           `json:"age,omitempty"`
	Address string        `json:"address,omitempty"`
	RollNo  int           `json:"roll_no"`
	Courses []enum.Course `json:"courses,omitempty"` // set of courses
}

func NewUser(name string, age int, address string, rollNo int, courses []string) (user User, err error) {
	user.Name = name
	user.Age = age
	user.Address = address
	user.RollNo = rollNo

	var c []enum.Course
	c, err = getCourses(courses)
	if err != nil {
		err = errors.Wrap(err, "newUser")
		return
	}

	user.Courses = c

	if err = user.validate(); err != nil {
		err = errors.Wrap(err, "newUser")
		return
	}

	return
}

func getCourses(courses []string) (courseEnum []enum.Course, err error) {
	for _, c := range courses {
		course, err := enum.CourseString(c)
		if err != nil {
			return []enum.Course{}, errors.Wrap(err, "getCourse")
		}
		courseEnum = append(courseEnum, course)
	}

	return
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
			err := errors.New("should be unique")
			return errors.Wrap(err, "validate")
		}
		set[c] = struct{}{}
	}
	return nil
}

func CheckPositive(value interface{}) error {
	val := value.(int)
	if val <= 0 {
		return errors.Wrap(errors.New("must be positive"), "validate")
	}
	return nil
}

func (user User) String() string {
	return fmt.Sprintf("%s	|	%d	|	%s	|	%d	|	%s	|\n", user.Name, user.Age, user.Address, user.RollNo, courseString(user.Courses))
}

func courseString(course []enum.Course) (courses []string) {
	for _, c := range course {
		courses = append(courses, c.String())
	}
	return courses
}
