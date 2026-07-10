package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/dhalvyr/blog_gator/internal/config"
	"github.com/dhalvyr/blog_gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	dbQueries := database.New(db)
	
	savedState := state{db: dbQueries, cfg: &cfg}

	savedCommands := commands{commandList: map[string]func(*state, command) error{}}

	savedCommands.register("login", handlerLogin)
	savedCommands.register("register", handlerRegister)

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



