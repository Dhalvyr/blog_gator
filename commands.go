package main

import (
	"github.com/dhalvyr/blog_gator/internal/config"
	"fmt"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	commandList map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.commandList[cmd.name]
	if !ok {
		return fmt.Errorf("Command not found.")
	}
	err := handler(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandList[name] = f
}