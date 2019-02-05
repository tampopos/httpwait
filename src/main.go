package main

import (
	"fmt"
	"os"

	"github.com/tampopos/httpwait/src/di"
	"github.com/tampopos/httpwait/src/httpwait"
)

func main() {
	container, err := di.CreateContainer()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	if err := container.Invoke(executeHttpwait); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func executeHttpwait(useCase httpwait.UseCase) {
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
}
