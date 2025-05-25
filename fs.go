package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pocketbase/pocketbase"
)

type FSHandler struct {
	app  *pocketbase.PocketBase
	root string
}

func NewFSHandler(app *pocketbase.PocketBase, root string) *FSHandler {
	return &FSHandler{
		app:  app,
		root: root,
	}
}

func updateFrontmatterJSON(filePath string, frontmatterJSON string, content string) error {
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	yamlFrontmatter := jsonToYaml(frontmatterJSON)
	if yamlFrontmatter == "" {
		return fmt.Errorf("empty yaml")
	}

	contentStr := string(contentBytes)
	startIdx := strings.Index(contentStr, "---\n")

	if startIdx == -1 {
		// no frontmatter, just add new frontmatter and the content (or existing content if content arg empty)
		useContent := content
		if useContent == "" {
			useContent = contentStr
		}
		// ensure exactly one newline between frontmatter and content
		newContent := "---\n" + yamlFrontmatter + "---\n\n" + useContent
		return os.WriteFile(filePath, []byte(newContent), 0644)
	}

	// find end of frontmatter delimiter
	endIdx := strings.Index(contentStr[startIdx+4:], "---")
	if endIdx == -1 {
		// no closing delimiter, treat as no frontmatter
		useContent := content
		if useContent == "" {
			useContent = contentStr
		}
		// ensure exactly one newline between frontmatter and content
		newContent := "---\n" + yamlFrontmatter + "---\n\n" + useContent
		return os.WriteFile(filePath, []byte(newContent), 0644)
	}

	endIdx += startIdx + 4

	// Compose new content with new frontmatter and the given content if any, else existing content from file after frontmatter
	useContent := content
	if useContent == "" {
		useContent = contentStr[endIdx+3:]
	}

	// trim leading newlines on useContent to avoid excess newlines
	useContent = strings.TrimLeft(useContent, "\r\n")

	// ensure exactly one newline between frontmatter and content
	newContent := "---\n" + yamlFrontmatter + "---\n\n" + useContent

	return os.WriteFile(filePath, []byte(newContent), 0644)
}
