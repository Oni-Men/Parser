package main

import (
	"fmt"
	"os"
	"os/user"
	"./repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s, Peach Programming Language\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
