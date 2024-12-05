package main

import (
	"fmt"

	"github.com/khodaid/Sablon/internal/config"
)

func main() {
	fmt.Println("Hello, World!")
	config.InitFlags()
	config.Run()
}
