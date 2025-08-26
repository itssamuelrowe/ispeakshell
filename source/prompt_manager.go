package source

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type PromptManager struct {
	promptsDir string
}

func NewPromptManager() *PromptManager {
	return &PromptManager{
		promptsDir: "prompts",
	}
}

func (pm *PromptManager) RenderTemplate(handle string, data map[string]interface{}) (string, error) {
	templatePath := filepath.Join(pm.promptsDir, handle)

	content, err := os.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to read template %s: %v", handle, err)
	}

	tmpl, err := template.New(handle).Parse(string(content))
	if err != nil {
		return "", fmt.Errorf("failed to parse template %s: %v", handle, err)
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template %s: %v", handle, err)
	}

	return buf.String(), nil
}
