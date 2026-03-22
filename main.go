package main

import (
	"fmt"
	"os"
)

func gitInit() {
	fmt.Println("initialized")
	os.Mkdir(".minigit", 0755)
	os.Mkdir(".minigit/objects", 0755)
	os.Mkdir(".minigit/refs", 0755)
	os.Mkdir(".minigit/refs/heads", 0755)
	os.WriteFile(".minigit/HEAD", []byte("ref: refs/heads/main"), 0644)
}

func main() {
	command := os.Args[1]
	switch command {
	case "init":
		gitInit()
	default:
		panic("no command")
	}
}
