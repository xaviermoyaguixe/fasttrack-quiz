package main

import (
	"fasttrackquiz/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal()
	}
}
