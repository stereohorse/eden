package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:])
}

func getSubcommand(args []string) {
	if len(args) < 2 {
		return
	}
}
