package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	log.Print("This Program is executed command!\n")
	out := exec.Command("sleep", "2")
	error := command.Start()

	if error != nil {
		panic(fmt.Sprintf("Error1: %v", error))
	}

	log.Printf("Command start!")

	erorr = command.Wait()

	if error != nil {
		panic(fmt.Sprintf("Error2: %v", error))
	}

	log.Printf("Command finish!")
}
