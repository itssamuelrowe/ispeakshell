package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/term"
	"ispeakshell/source"
	"log"
	"os"
	"strings"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	if len(os.Args) > 1 {
		// Mode 1: Command line argument
		prompt := strings.Join(os.Args[1:], " ")
		processPrompt(prompt)
	} else {
		// Mode 2: Interactive mode
		interactiveMode()
	}
}

type Command struct {
	Code        string `json:"code"`
	Explanation string `json:"explanation"`
	Summary     string `json:"summary"`
}

type Response struct {
	Summary  string    `json:"summary"`
	Commands []Command `json:"commands"`
}

func processPrompt(prompt string) {
	pm := source.NewPromptManager()
	renderedPrompt, err := pm.RenderTemplate("basic.template", map[string]interface{}{
		"Prompt": prompt,
	})
	if err != nil {
		fmt.Printf("Error rendering template: %v\n", err)
		return
	}

	response, err := source.CallGemini(renderedPrompt)
	if err != nil {
		fmt.Println(err)
		return
	}

	var parsed Response
	if err := json.Unmarshal([]byte(response), &parsed); err != nil {
		fmt.Printf("Error parsing response: %v\n", err)
		return
	}

	// Pretty header
	header := color.New(color.FgCyan, color.Bold)
	header.Printf("\n%s\n\n", strings.TrimSpace(parsed.Summary))

	// Create table for commands
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Task", "Description"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)

	// Set column widths
	table.SetColMinWidth(0, 3)  // # column
	table.SetColMinWidth(1, 25) // Task column
	table.SetColMinWidth(2, 80) // Description column (max 80 chars)
	if width, _, err := term.GetSize(int(os.Stdout.Fd())); err == nil {
		table.SetColWidth(width - 10)
	}

	for i, cmd := range parsed.Commands {
		table.Append([]string{
			color.New(color.FgYellow, color.Bold).Sprintf("%d", i+1),
			color.New(color.FgGreen).Sprintf("%s", cmd.Summary),
			color.New(color.FgWhite).Sprintf("%s", cmd.Explanation),
		})
	}
	table.Render()
	fmt.Println()

	// Execute commands with pretty formatting
	for i, cmd := range parsed.Commands {
		cmdColor := color.New(color.FgMagenta, color.Bold)
		cmdColor.Printf("[%d] %s\n", i+1, cmd.Code)
		source.ConfirmAndExecute(cmd.Code)
	}
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
				term.Restore(int(os.Stdin.Fd()), oldState)
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
