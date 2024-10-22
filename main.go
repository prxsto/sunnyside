package main

import (
	"log"

	"github.com/prxsto/sunnyside/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
