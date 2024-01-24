package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/pkg/errors"
	enum "github.com/sahaj279/go_assignment/repository/data_field"
	"github.com/sahaj279/go_assignment/user"
)

type Svc interface {
	Load(dataFilePath string) error
	Add(user.User) error
	List(field enum.DataField, ASCOrder bool) (users []user.User)
	Delete(rollNo int) error
	Save(users []user.User) error
	Close() error
}

type Repository struct {
	users map[int]user.User
	file  *os.File
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Load(dataFilePath string) error {
	if err := open(r, dataFilePath); err != nil {
		return errors.Wrap(err, "load")
	}

	r.users = make(map[int]user.User)

	users, err := retrieveData(r)
	if err != nil {
		return errors.Wrap(err, "load")
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
		return errors.Wrap(err, "open")
	}

	r.file = file

	return nil
}

func retrieveData(r *Repository) ([]user.User, error) {
	fs, err := r.file.Stat()
	if err != nil {
		return []user.User{}, errors.Wrap(err, "retrieveData")
	}

	len := fs.Size()
	if len == 0 {
		return []user.User{}, errors.Wrap(err, "retrieveData")
	}

	dataB := make([]byte, len)
	_, err = r.file.Read(dataB)
	if err != nil {
		return []user.User{}, errors.Wrap(err, "retrieveData")
	}

	users, err := DeserializeUsers(dataB)
	if err != nil {
		return []user.User{}, errors.Wrap(err, "retrieveData")
	}

	return users, nil
}

func (r *Repository) Add(user user.User) error {
	if _, exist := r.users[user.RollNo]; exist {
		err := fmt.Errorf("user already exists for %d roll number", user.RollNo)
		return errors.Wrap(err, "add")
	}

	r.users[user.RollNo] = user
	return nil
}

func (r *Repository) List(field enum.DataField, ascOrder bool) (users []user.User) {
	for _, user := range r.users {
		users = append(users, user)
	}
	sortUsers(users, field)
	if !ascOrder {
		slices.Reverse(users)
	}

	return
}

func sortUsers(users []user.User, field enum.DataField) {
	sort.SliceStable(users, func(i, j int) bool {
		switch field {
		case enum.Name:
			if strings.Compare(users[i].Name, users[j].Name) == 0 {
				return (users[i].RollNo < users[j].RollNo)
			}
			return (strings.Compare(users[i].Name, users[j].Name) == -1)
		case enum.RollNo:
			return (users[i].RollNo < users[j].RollNo)
		case enum.Address:
			if strings.Compare(users[i].Address, users[j].Address) == 0 {
				return (users[i].RollNo < users[j].RollNo)
			}
			return (strings.Compare(users[i].Address, users[j].Address) == -1)
		case enum.Age:
			if users[i].RollNo == users[j].RollNo {
				return (users[i].RollNo < users[j].RollNo)
			}
			return (users[i].Age < users[j].Age)
		default:
			return true
		}
	})
}

func (r *Repository) Delete(rollNo int) error {
	if _, exist := r.users[rollNo]; !exist {
		err := errors.Errorf("user does not exist for roll number %d", rollNo)
		return errors.Wrap(err, "delete")
	}

	delete(r.users, rollNo)
	return nil
}

func (r *Repository) Save(users []user.User) error {
	dataB, err := SerializeUsers(users)
	if err != nil {
		return errors.Wrap(err, "save")
	}

	if err = r.file.Truncate(0); err != nil {
		return errors.Wrap(err, "save")
	}

	_, err = r.file.Seek(0, 0)
	if err != nil {
		return errors.Wrap(err, "save")
	}

	_, err = r.file.Write(dataB)
	if err != nil {
		return errors.Wrap(err, "save")
	}

	return nil
}

func SerializeUsers(users []user.User) ([]byte, error) {
	userB, err := json.Marshal(users)
	if err != nil {
		return []byte{}, errors.Wrap(err, "encodeUse")
	}

	return userB, nil
}

func DeserializeUsers(userB []byte) ([]user.User, error) {
	var users []user.User
	if err := json.Unmarshal(userB, &users); err != nil {
		return []user.User{}, errors.Wrap(err, "decodeUser")
	}

	return users, nil
}

func (r *Repository) Close() error {
	return r.file.Close()
}
