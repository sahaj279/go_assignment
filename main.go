package main

import (
	"github.com/pkg/errors"
	"github.com/sahaj279/go_assignment/menu"
	repo "github.com/sahaj279/go_assignment/repository"
)

func main() {
	repository := repo.Repository{}
	newMenu := menu.NewMenu(&repository)
	if err := newMenu.Init(); err != nil {
		menu.PrintError(errors.Wrap(err, "main"))
	}
}
