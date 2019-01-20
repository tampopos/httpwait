package main

import (
	"fmt"
	"os"

	"./client"
	"./httpwait"
	"./stopwatch"
)

func main() {
	executeHttpwait()
}

func executeHttpwait() {
	var c = client.Create()
	var s = stopwatch.New()
	var useCase = httpwait.CreateUseCase(s, c)
	var args, err = useCase.CreateArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	err = useCase.Wait(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Success!\n")
}
