package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execInput(input string) error {
	// to remove the newline thingy
	input = strings.TrimSuffix(input, "\n")

	// split the input
	args := strings.Split(input, " ")

	// doing this because the "cd" command doesn't actually exist
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("Path is required")
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	// execute the command and return the error
	return cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">>")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			// the ; makes the variable available only inside the if statement, scope setting and all :3
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
