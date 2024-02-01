package menu

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	repo "github.com/sahaj279/go_assignment/repository"
)

const (
	AddUser int = iota + 1
	DisplayUser
	DeleteUser
	SaveUser
	Exit
)

func NewMenu(repository repo.Svc) *Menu {
	return &Menu{
		repository: repository,
	}
}

type Menu struct {
	repository repo.Svc
}

func (m *Menu) Init() error {
	defer m.repository.Close()
	if err := m.repository.Load(DataFilePath); err != nil {
		return errors.Wrap(err, "init")
	}

	defer fmt.Println("Menu application finished!")

	for {
		showMenu()
		var choice int
		var err error

		choice, err = getChoice()
		if err != nil {
			return errors.Wrap(err, "init")
		}

		switch choice {
		case AddUser:
			if err = m.addUser(); err != nil {
				return errors.Wrap(err, "init")
			}
		case DisplayUser:
			if err = m.displayUsers(); err != nil {
				return errors.Wrap(err, "init")
			}
		case DeleteUser:
			if err = m.deleteUser(); err != nil {
				return errors.Wrap(err, "init")
			}
		case SaveUser:
			if err = m.saveUser(); err != nil {
				return errors.Wrap(err, "init")
			}
		case Exit:
			if err = m.confirmSave(); err != nil {
				return errors.Wrap(err, "init")
			}
			return nil
		}
	}
}

func getChoice() (choice int, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	var userChoice string

	fmt.Println("\nEnter a digit in between 1-5 :")
	if scanner.Scan() {
		userChoice = scanner.Text()
		userChoice = strings.TrimSpace(userChoice)
	}

	if err = scanner.Err(); err != nil {
		return 0, errors.Wrap(err, "getChoice")
	}

	choice, err = strconv.Atoi(userChoice)
	if err != nil {
		return 0, errors.Wrap(err, "getChoice")
	}

	if choice < 1 || choice > 5 {
		err = (errors.New("choice should be in between 1 and 5"))
		return 0, errors.Wrap(err, "getChoice")
	}
	return
}

func showMenu() {
	fmt.Println("\nMenu")
	fmt.Println("1. Add user details")
	fmt.Println("2. Display user details")
	fmt.Println("3. Delete user details")
	fmt.Println("4. Save user details")
	fmt.Println("5. Exit")
}

func PrintError(err error) {
	log.Println(err)
}
