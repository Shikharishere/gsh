package main

import (
	"bufio"
	"container/list"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	hostname, err := os.Hostname()
	history := list.New()

	for {
		dir, errDir := os.Getwd()
		if err == nil && errDir == nil {
			if len(dir) > 20 {
				splittedDir := strings.Split(dir, "/")
				fmt.Print("👀 <" + hostname + "> '" + "../" + splittedDir[len(splittedDir)-1] + "' ")
			} else {
				fmt.Print("👀 <" + hostname + "> '" + dir + "' ")
			}
		} else {
			fmt.Print("👀 gsh ")
		}
		// Read the keyboard input.
		input, err := reader.ReadString('\n')

		// Read up or down keys
		// input, err = reader.ReadString("^[[A")

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// handle exec
		if err = execInput(input, history); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string, history *list.List) error {
	// remove the new line character
	input = strings.TrimSuffix(input, "\n")

	history.PushFront(input)

	// split the input to separate the command and args
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}

		return os.Chdir(args[1])
	case "history":
		head := history.Front()
		for head != nil {
			fmt.Println(head.Value)
			head = head.Next()
		}
		return nil;
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
