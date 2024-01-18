package repository

import (
	"assignment2/user"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	Name    = "Name"
	Age     = "Age"
	RollNo  = "RollNo"
	Address = "Address"
)

type RepositoryOps interface {
	Load(dataFilePath string) error
	Add(user.User) error
	List(field string, ASCOrder bool) (users []user.User, err error)
	Delete(rollno int) error
	Save(users []user.User) error
	Close() error
}

type Repository struct {
	users map[int]user.User
	file  *os.File
}

func NewRepo() *Repository {
	return &Repository{}
}

func (r *Repository) Load(dataFilePath string) error {
	if err := open(r, dataFilePath); err != nil {
		return err
	}

	r.users = make(map[int]user.User)

	users, err := retrieveData(r)
	if err != nil {
		return err
	}

	for _, user := range users {
		r.users[user.RollNo] = user
	}

	return nil
}

func open(r *Repository, dataFilePath string) error {
	if r.file != nil {
		return nil
	}

	file, err := os.OpenFile(dataFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	r.file = file

	return nil
}

func retrieveData(r *Repository) ([]user.User, error) {
	fs, err := r.file.Stat()
	if err != nil {
		return []user.User{}, err
	}

	len := fs.Size()
	if len == 0 {
		return []user.User{}, err
	}

	dataB := make([]byte, len)
	_, err = r.file.Read(dataB)
	if err != nil {
		return []user.User{}, err
	}

	users, err := user.DecodeUsers(dataB)
	if err != nil {
		return []user.User{}, err
	}

	return users, nil
}

func (r *Repository) Add(user user.User) error {
	if _, exist := r.users[user.RollNo]; exist {
		err := fmt.Errorf("user already exists for %d roll number", user.RollNo)
		return err
	}

	r.users[user.RollNo] = user
	return nil
}

func (r *Repository) List(field string, AscOrder bool) ([]user.User, error) {
	var users []user.User
	for _, user := range r.users {
		users = append(users, user)
	}

	if AscOrder {
		sortAscCustom(users, field)
	} else {
		sortDescCustom(users, field)
	}

	return users, nil
}

func sortAscCustom(users []user.User, field string) {
	sort.SliceStable(users, func(i, j int) bool {
		switch field {
		case Name:
			return (strings.Compare(users[i].Name, users[j].Name) == -1)
		case RollNo:
			return (users[i].RollNo < users[j].RollNo)
		case Address:
			return (strings.Compare(users[i].Address, users[j].Address) == -1)
		case Age:
			return (users[i].Age < users[j].Age)
		default:
			return true
		}
	})
}

func sortDescCustom(users []user.User, field string) {
	sort.SliceStable(users, func(i, j int) bool {
		switch field {
		case Name:
			return (strings.Compare(users[i].Name, users[j].Name) == 1)
		case RollNo:
			return (users[i].RollNo > users[j].RollNo)
		case Address:
			return (strings.Compare(users[i].Address, users[j].Address) == 1)
		case Age:
			return (users[i].Age > users[j].Age)
		default:
			return true
		}
	})
}

func (r *Repository) Delete(rollno int) error {
	if _, exist := r.users[rollno]; !exist {
		err := fmt.Errorf("user does not exist for %d roll number", rollno)

		return err
	}

	delete(r.users, rollno)

	return nil
}

func (r *Repository) Save(users []user.User) error {
	dataB, err := user.EncodeUsers(users)
	if err != nil {
		return err
	}

	if err = r.file.Truncate(0); err != nil {
		return err
	}
	_, err = r.file.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = r.file.Write(dataB)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Close() error {
	return r.file.Close()
}
