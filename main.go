package main

import (
	"github.com/pkg/errors"
	"github.com/sahaj279/go-assignment/cli"
)

func main() {
	if err := cli.Init(); err != nil {
		cli.LogError(errors.Wrap(err, "main"))
	}
}
