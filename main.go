package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execInput(input string) error {
	// to remove the newline thingy
	input = strings.TrimSuffix(input, "\n")

	// prepare the command to execute
	cmd := exec.Command(input)

	// set the correct output device
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
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
