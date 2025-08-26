package source

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type OllamaRequest struct {
	Model    string    `json:"model"`
	Stream   bool      `json:"stream"`
	Messages []Message `json:"messages"`
	Format   string    `json:"format"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OllamaResponse struct {
	Message struct {
		Content string `json:"content"`
	} `json:"message"`
}

func CallOllama(prompt string) (string, error) {
	// log.Printf("Calling Ollama with prompt: %s", prompt)

	req := OllamaRequest{
		Model:  "codegemma:7b",
		Stream: false,
		Format: "json",
		Messages: []Message{
			{Role: "user", Content: prompt},
		},
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	log.Println("Calling Ollama API...")
	resp, err := http.Post("http://127.0.0.1:11434/api/chat", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("Error: Ollama server not running")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Ollama API error: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var ollamaResp OllamaResponse
	if err := json.Unmarshal(body, &ollamaResp); err != nil {
		return "", err
	}

	content := strings.TrimSpace(ollamaResp.Message.Content)
	fmt.Println(ollamaResp)

	if content == "" {
		return "", fmt.Errorf("Could not generate a valid command")
	}

	return content, nil
}
