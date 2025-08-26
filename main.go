package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func main() {
	if len(os.Args) > 1 {
		// Mode 1: Command line argument
		prompt := strings.Join(os.Args[1:], " ")
		processPrompt(prompt)
	} else {
		// Mode 2: Interactive mode
		interactiveMode()
	}
}

func processPrompt(prompt string) {
	fmt.Printf("Processing: %s\n", prompt)
	// Add your prompt processing logic here
}

func interactiveMode() {
	var lines []string
	var currentLine string
	ghostShown := true
	
	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	
	fmt.Print("> \033[90mEnter prompt\033[0m")
	fmt.Print("\033[13D")
	
	for {
		b := make([]byte, 1)
		os.Stdin.Read(b)
		c := b[0]
		
		if c == 13 || c == 10 { // Enter
			if ghostShown {
				fmt.Print("\033[K")
				ghostShown = false
			}
			fmt.Print("\r\n")
			if currentLine == "" && len(lines) > 0 {
				prompt := strings.Join(lines, "\n")
				processPrompt(prompt)
				return
			}
			lines = append(lines, currentLine)
			currentLine = ""
			if len(lines) == 1 {
				fmt.Print("> \033[90mPress enter again to run the command\033[0m")
				fmt.Print("\033[39D")
			} else {
				fmt.Print("> ")
			}
			ghostShown = len(lines) > 0
		} else if c == 127 || c == 8 { // Backspace
			if len(currentLine) > 0 {
				currentLine = currentLine[:len(currentLine)-1]
				fmt.Print("\b \b")
			}
		} else if c >= 32 { // Printable character
			if ghostShown {
				fmt.Print("\033[K")
				ghostShown = false
			}
			currentLine += string(c)
			fmt.Print(string(c))
		}
	}
}