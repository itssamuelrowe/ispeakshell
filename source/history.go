package source

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type HistoryEntry struct {
	Prompt  string `json:"prompt"`
	Command string `json:"command"`
}

func SaveCommand(prompt, command string) {
	homeDir, _ := os.UserHomeDir()
	historyDir := filepath.Join(homeDir, ".ispeakshell")
	historyFile := filepath.Join(historyDir, "history.json")

	os.MkdirAll(historyDir, 0755)

	var history []HistoryEntry
	if data, err := os.ReadFile(historyFile); err == nil {
		json.Unmarshal(data, &history)
	}

	entry := HistoryEntry{Prompt: prompt, Command: command}
	history = append(history, entry)

	data, _ := json.MarshalIndent(history, "", "  ")
	os.WriteFile(historyFile, data, 0644)

	fmt.Println("Command saved to history.")
}
