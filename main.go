package main

import (
	"github.com/pkg/errors"
	"github.com/sahaj279/go-assignment/cli"
	"github.com/sahaj279/go-assignment/item"
)

func main() {
	// Passing itemSvc which implements itemHandler to create cli to be able to mock createItem and calculateTax
	// Dependency injection

	var itemScv item.ItemSvc
	newCli := cli.NewCli(itemScv)

	if err := newCli.Init(); err != nil {
		cli.LogError(errors.Wrap(err, "main"))
	}
}
