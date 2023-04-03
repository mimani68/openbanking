package main

import (
	"log"

	"github.com/mimani68/fintech-core/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
