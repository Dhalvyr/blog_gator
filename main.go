package main

import (
	"github.com/dhalvyr/blog_gator/internal/config"
	"fmt"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	
	savedState := state{cfg: &cfg}

	savedCommands := commands{commandList: map[string]func(*state, command) error{}}

	savedCommands.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("A command name is required")
		os.Exit(1)
	}
	
	cmd := command{name: args[1], args: args[2:]}
	err = savedCommands.run(&savedState, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}



