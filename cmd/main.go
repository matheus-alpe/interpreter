package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/matheus-alpe/interpreter/internal/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s, welcome to Shitscript!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
