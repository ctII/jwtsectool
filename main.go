package main

import (
	"github.com/ctii/jwtsectool/cmd"
	"log"
)

func main() {
	err := cmd.Main()
	if err != nil {
		log.Fatalln(err)
	}
}
