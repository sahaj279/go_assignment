package menu

import (
	repo "assignment2/repository"
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Init() error {
	// users repository for storing in memory
	repository := repo.NewRepo()

	// loading users from file
	if err := repository.Load(DataFilePath); err != nil {
		return err
	}

	// closing the file at end
	defer repository.Close()

	// show option menu
	for {
		showMenu()
		var choice int
		var err error

		// Getting user choice until we get a valid choice

		choice, err = getUserChoice()
		if err != nil {
			return err
		} else if choice < 1 || choice > 5 {
			return (errors.New("choice should be in between 1 and 5"))
		}

		// Performing operation based on selected choice
		switch choice {
		case 1:
			if err = addUser(*repository); err != nil {
				return err
			}
		case 2:
			if err = displayUsers(repository); err != nil {
				return err
			}
		case 3:
		case 4:
		case 5:
		}
	}
}

func getUserChoice() (choice int, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	var userChoice string

	fmt.Println("\nEnter a digit in between 1-5 :")
	if scanner.Scan() {
		userChoice = scanner.Text()
		userChoice = strings.TrimSpace(userChoice)
	}
	if err = scanner.Err(); err != nil {
		return
	}

	choice, err = strconv.Atoi(userChoice)
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
