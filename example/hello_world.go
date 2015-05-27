package main

import (
	"fmt"
	"github.com/choix/goluajit"
)

func main() {
	L := luajit.NewState()
	defer L.Close()

	L.OpenLibs()

	lines := []string{
		`text = "Hello"`,
		`name = "Wolrd"`,
		`print(text .. ", " .. name .. "!")`,
	}
	var err error

	for _, line := range lines {
		err = L.LoadString(line)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		err = L.Run()
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
	}
}
