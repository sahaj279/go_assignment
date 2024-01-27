package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	familytree "github.com/sahaj279/go_assignment/service"

	"github.com/pkg/errors"
)

const (
	Accept = "y"
	Deny   = "n"
)

func Init() error {
	familyTree := familytree.NewFamilyTree()

	var moreInput bool = true
	for moreInput {
		showMenu()

		userChoice, err := GetChoice()
		if err != nil {
			return errors.Wrap(err, "err in getChoice")
		}

		switch userChoice {
		case "1":
			if err := ListParent(familyTree); err != nil {
				fmt.Println(err)
			}
		case "2":
			if err := ListChild(familyTree); err != nil {
				fmt.Println(err)
			}
		case "3":
			if err := ListAncestors(familyTree); err != nil {
				fmt.Println(err)
			}
		case "4":
			if err := ListDescendants(familyTree); err != nil {
				fmt.Println(err)
			}
		case "5":
			if err := deleteDependency(familyTree); err != nil {
				fmt.Println(err)
			}
		case "6":
			if err := deleteNode(familyTree); err != nil {
				fmt.Println(err)
			}
		case "7":
			if err := addDependency(familyTree); err != nil {
				fmt.Println(err)
			}
		case "8":
			if err := addNode(familyTree); err != nil {
				fmt.Println(err)
			}
		case "9":
			moreInput = false
		default:
			fmt.Println("Invalid choice")
		}
	}

	return nil
}

func showMenu() {
	fmt.Println("-------------------")
	fmt.Println("1. Disply the immediate parents of a node")
	fmt.Println("2. Display the immediate children of a node")
	fmt.Println("3. Display the ancestors of a node")
	fmt.Println("4. Display the descendants of a node")
	fmt.Println("5. Delete the dependency")
	fmt.Println("6. Delete the node")
	fmt.Println("7. Add dependency")
	fmt.Println("8. Add node")
	fmt.Println("9. Exit")
	fmt.Println("-------------------")
}

func GetChoice() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var userChoice string

	if scanner.Scan() {
		userChoice = scanner.Text()
		userChoice = strings.TrimSpace(userChoice)
	}
	if err := scanner.Err(); err != nil {
		return "", errors.Wrap(err, "scan for user's choice failed")
	}

	return userChoice, nil
}

func ListParent(familyTree familytree.Svc) error {
	scanner := bufio.NewScanner(os.Stdin)

	var id int

	fmt.Printf("Enter id: ")

	if scanner.Scan() {
		i, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println("scan for node's id (List immediate parent) failed")
			return errors.Wrap(err, "scan for node's id (List immediate parent) failed")
		}
		id = i
	}
	if err := scanner.Err(); err != nil {
		log.Println("scan for node's id (List immediate parent) failed")
		return errors.Wrap(err, "scan for node's id (List immediate parent) failed")
	}

	nodes, err := familyTree.GetParents(id)
	if err != nil {
		return err
	}

	fmt.Println(familyTree.ListNodesID(nodes))

	return nil
}

func ListChild(familyTree familytree.Svc) error {
	scanner := bufio.NewScanner(os.Stdin)

	var id int

	fmt.Printf("Enter id: ")

	if scanner.Scan() {
		i, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println("scan for node's id (List immediate child) failed")
			return errors.Wrap(err, "scan for node's id (List immediate child) failed")
		}
		id = i
	}
	if err := scanner.Err(); err != nil {
		log.Println("scan for node's id (List immediate child) failed")
		return errors.Wrap(err, "scan for node's id (List immediate child) failed")
	}

	nodes, err := familyTree.GetChildren(id)
	if err != nil {
		return err
	}

	fmt.Println(familyTree.ListNodesID(nodes))

	return nil
}

func ListAncestors(familyTree familytree.Svc) error {
	scanner := bufio.NewScanner(os.Stdin)

	var id int

	fmt.Printf("Enter id: ")

	if scanner.Scan() {
		i, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println("scan for node's id (List ancestors) failed")
			return errors.Wrap(err, "scan for node's id (List ancestors) failed")
		}
		id = i
	}
	if err := scanner.Err(); err != nil {
		log.Println("scan for node's id (List ancestors) failed")
		return errors.Wrap(err, "scan for node's id (List ancestors) failed")
	}

	nodes, err := familyTree.GetAncestors(id)
	if err != nil {
		return err
	}

	fmt.Println(familyTree.ListNodesID(nodes))

	return nil
}

func ListDescendants(familyTree familytree.Svc) error {
	scanner := bufio.NewScanner(os.Stdin)

	var id int

	fmt.Printf("Enter id: ")

	if scanner.Scan() {
		i, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println("scan for node's id (List descendants) failed")
			return errors.Wrap(err, "scan for node's id (List descendants) failed")
		}
		id = i
	}
	if err := scanner.Err(); err != nil {
		log.Println("scan for node's id (List descendants) failed")
		return errors.Wrap(err, "scan for node's id (List descendants) failed")
	}

	nodes, err := familyTree.GetDescendants(id)
	if err != nil {
		return err
	}

	fmt.Println(familyTree.ListNodesID(nodes))

	return nil
}

func deleteDependency(familyTree familytree.Svc) error {
	scanner := bufio.NewScanner(os.Stdin)

	var id1 int

	fmt.Println("Enter ids of nodes")

	if scanner.Scan() {
		i1, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println("scan for node's id-1 failed while deleting dependency")
			return errors.Wrap(err, "scan for node's id-1 failed while deleting dependency")
		}
		id1 = i1
	}
	if err := scanner.Err(); err != nil {
		err = errors.Wrap(err, "scan for node's id-1 failed while deleting dependency")
		log.Println(err)
		return err
	}

	var id2 int

	if scanner.Scan() {
		i2, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println("scan for node's id-2 failed while deleting dependency")
			return errors.Wrap(err, "scan for node's id-2 failed while deleting dependency")
		}
		id2 = i2
	}
	if err := scanner.Err(); err != nil {
		log.Println("scan for node's id-2 failed while deleting dependency")
		return errors.Wrap(err, "scan for node's id-2 failed while deleting dependency")
	}

	if err := familyTree.DeleteEdge(id1, id2); err != nil {
		return err
	}

	return nil
}

func deleteNode(familyTree familytree.Svc) error {
	scanner := bufio.NewScanner(os.Stdin)

	var id int

	fmt.Println("Enter id of node")

	if scanner.Scan() {
		i, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println("scan for node's id failed while deleting")
			return errors.Wrap(err, "scan for node's id failed while deleting")
		}
		id = i
	}
	if err := scanner.Err(); err != nil {
		log.Println("scan for node's id failed while deleting")
		return errors.Wrap(err, "scan for node's id failed while deleting")
	}

	if err := familyTree.DeleteNode(id); err != nil {
		return err
	}

	return nil
}

func addDependency(familyTree familytree.Svc) error {
	scanner := bufio.NewScanner(os.Stdin)

	var id1 int

	fmt.Println("Enter ids of nodes")

	if scanner.Scan() {
		i1, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println("scan for node's id-1 failed while adding dependency")
			return errors.Wrap(err, "scan for node's id-1 failed while adding dependency")
		}
		id1 = i1
	}
	if err := scanner.Err(); err != nil {
		log.Println("scan for node's id-1 failed while adding dependency")
		return errors.Wrap(err, "scan for node's id-1 failed while adding dependency")
	}

	var id2 int

	if scanner.Scan() {
		i2, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println("scan for node's id-2 failed while adding dependency")
			return errors.Wrap(err, "scan for node's id-2 failed while adding dependency")
		}
		id2 = i2
	}
	if err := scanner.Err(); err != nil {
		log.Println("scan for node's id-2 failed while adding dependency")
		return errors.Wrap(err, "scan for node's id-2 failed while adding dependency")
	}

	if err := familyTree.AddEdge(id1, id2); err != nil {
		return err
	}

	return nil
}

func addNode(familyTree familytree.Svc) error {
	scanner := bufio.NewScanner(os.Stdin)

	var id int

	fmt.Printf("Id: ")

	if scanner.Scan() {
		i, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println("scan for node's id failed")
			return errors.Wrap(err, "scan for node's id failed")
		}
		id = i
	}
	if err := scanner.Err(); err != nil {
		log.Println("scan for node's id failed")
		return errors.Wrap(err, "scan for node's id failed")
	}

	var name string

	fmt.Printf("Name: ")

	if scanner.Scan() {
		name = scanner.Text()
		name = strings.TrimSpace(name)
	}
	if err := scanner.Err(); err != nil {
		log.Println("scan for node's name failed")
		return errors.Wrap(err, "scan for node's name failed")
	}

	metaData := make(map[string]string)

	if err := ListAdditionInfo(metaData); err != nil {
		log.Println("scan for node's metadata failed")
		return errors.Wrap(err, "scan for node's metadata failed")
	}

	if err := familyTree.AddNode(id, name, metaData); err != nil {
		return err
	}

	return nil
}

func ListAdditionInfo(metaData map[string]string) error {
	scanner := bufio.NewScanner(os.Stdin)

	var userChoice string

	fmt.Printf("Additional Info (" + Accept + "/" + Deny + "): ")

	if scanner.Scan() {
		userChoice = scanner.Text()
		userChoice = strings.TrimSpace(userChoice)
	}
	if err := scanner.Err(); err != nil {
		log.Println("scan for user's choice for meta data failed")
		return errors.Wrap(err, "scan for user's choice for meta data failed")
	}

	if userChoice != Accept {
		return nil
	}
	var key string

	fmt.Print("Enter key: ")

	if scanner.Scan() {
		key = scanner.Text()
		key = strings.TrimSpace(key)
	}
	if err := scanner.Err(); err != nil {
		log.Println("scan for node's meta data failed")
		return errors.Wrap(err, "scan for node's meta data failed")
	}

	var value string

	fmt.Print("Enter value: ")

	if scanner.Scan() {
		value = scanner.Text()
		value = strings.TrimSpace(value)
	}
	if err := scanner.Err(); err != nil {
		log.Println("scan for node's meta data failed")
		return errors.Wrap(err, "scan for node's meta data failed")
	}

	metaData[key] = value

	if err := ListAdditionInfo(metaData); err != nil {
		return err
	}

	return nil
}
