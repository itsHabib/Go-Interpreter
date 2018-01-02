package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/itsHabib/go-interp/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the MH programming language!\n",
		user.Username)
	fmt.Print("Type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
