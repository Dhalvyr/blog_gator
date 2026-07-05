package main

import (
	"github.com/dhalvyr/blog_gator/internal/config"
	"fmt"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	
	err = cfg.SetUser("Dhalvyr")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cfg)
}



