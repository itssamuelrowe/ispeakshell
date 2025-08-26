package source

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strings"
)

func ConfirmAndExecute(command string) {
	prompt := color.New(color.FgYellow)
	prompt.Print("Run this command? ")
	options := color.New(color.FgHiBlack)
	options.Print("[y/n/save]: ")

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		choice := strings.ToLower(strings.TrimSpace(input))

		switch choice {
		case "y", "yes":
			executeCommand(command)
			return
		case "n", "no":
			skipped := color.New(color.FgRed)
			skipped.Println("Command discarded.")
			return
		case "save":
			SaveCommand("", command)
			return
		default:
			error := color.New(color.FgRed)
			error.Print("Invalid input. Please enter ")
			options.Print("[y/n/save]: ")
		}
	}
}

func executeCommand(command string) {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		errorMsg := color.New(color.FgRed, color.Bold)
		errorMsg.Printf("Error executing command: %v\n", err)
	}
	fmt.Println()
}
