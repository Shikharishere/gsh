package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("gsh> ")
		// Read the keyboard input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// handle exec
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	// remove the new line character
	input = strings.TrimSuffix(input, "\n")

	// split the input to separate the command and args
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}

		return os.Chdir(args[1])
	case "exit":
		fmt.Println("Seeya 👋")
		os.Exit(0)
	}

	// prepare the command to exec
	cmd := exec.Command(args[0], args[1:]...)

	// setting the output devices
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// exec the command
	return cmd.Run()
}
