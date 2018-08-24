package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	kalosLoop()

	os.Exit(0)
}

func kalosLoop() {
	for {
		fmt.Printf("> ")
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		tokens := splitLine(line)
		launchCommand(tokens)
	}
}

func splitLine(line string) []string {
	tokens := strings.FieldsFunc(line, func(r rune) bool {
		return r == '\t' || r == '\r' || r == '\n' || r == '\a' || r == ' '
	})

	return tokens
}

func launchCommand(tokens []string) {
	cmd := exec.Command(tokens[0], tokens[1:]...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("There was a problem running command: %s\n", err)
	}
	fmt.Printf("%v", string(output))
}
