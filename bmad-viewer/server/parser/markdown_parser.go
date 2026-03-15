package parser

import (
	"os"
	"strings"

	"bmad-viewer/server/model"
	"gopkg.in/yaml.v3"
)

// ParseMarkdown reads a Markdown file and extracts frontmatter and content.
func ParseMarkdown(filePath string) (*model.Document, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	raw := string(data)
	frontmatter, content := splitFrontmatter(raw)

	var fm map[string]interface{}
	if frontmatter != "" {
		if err := yaml.Unmarshal([]byte(frontmatter), &fm); err != nil {
			fm = make(map[string]interface{})
		}
	} else {
		fm = make(map[string]interface{})
	}

	title := extractTitle(fm, content)

	return &model.Document{
		Content:     content,
		Frontmatter: fm,
		Title:       title,
	}, nil
}

// splitFrontmatter separates YAML frontmatter from Markdown content.
func splitFrontmatter(raw string) (string, string) {
	if !strings.HasPrefix(strings.TrimSpace(raw), "---") {
		return "", raw
	}

	trimmed := strings.TrimSpace(raw)
	// Find the closing ---
	rest := trimmed[3:] // skip opening ---
	idx := strings.Index(rest, "\n---")
	if idx == -1 {
		return "", raw
	}

	frontmatter := strings.TrimSpace(rest[:idx])
	content := strings.TrimSpace(rest[idx+4:]) // skip \n---
	return frontmatter, content
}

// extractTitle gets the title from frontmatter or first heading.
func extractTitle(fm map[string]interface{}, content string) string {
	if title, ok := fm["title"]; ok {
		if s, ok := title.(string); ok && s != "" {
			return s
		}
	}

	for _, line := range strings.Split(content, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "# ") {
			return strings.TrimPrefix(trimmed, "# ")
		}
	}

	return ""
}
