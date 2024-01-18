package main

import "assignment2/menu"

func main() {
	if err := menu.Init(); err != nil {
		menu.PrintError(err)
	}
}
