package main

import (
	"fmt"

	"github.com/skwb/realengo-conflict/core"
	"github.com/skwb/realengo-conflict/core/config"
)

func main() {
	var cfg, _ = config.LoadConfig()
	fmt.Println(cfg)
	core.StartGame()
}
