package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
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

func (h *FSHandler) Routes(se *core.ServeEvent) {
	fsGroup := se.Router.Group("/fs")
	fsGroup.Bind(apis.RequireSuperuserAuth())

	fsGroup.PATCH("/frontmatter", func(e *core.RequestEvent) error {
		return h.handleModifyFrontmatter(e)
	})
	fsGroup.PUT("/content", func(e *core.RequestEvent) error {
		return h.handleModifyContent(e)
	})
}

type FrontmatterRequest struct {
	Path string          `json:"path"`
	Data json.RawMessage `json:"data"`
}

func (h *FSHandler) handleModifyFrontmatter(e *core.RequestEvent) error {
	req := FrontmatterRequest{}
	if err := e.BindBody(&req); err != nil {
		return fmt.Errorf("invalid json body: %w", err)
	}

	fileRec, err := h.app.FindFirstRecordByData("files", "path", req.Path)
	if err != nil {
		return fmt.Errorf("file not found: %w", err)
	}

	newJSONBytes, err := json.Marshal(json.RawMessage(req.Data))
	if err != nil {
		return fmt.Errorf("invalid frontmatter json data: %w", err)
	}

	absPath := filepath.Join(h.root, fileRec.GetString("path"))
	err = updateFrontmatterJSON(absPath, string(newJSONBytes), "")
	if err != nil {
		return fmt.Errorf("failed to update frontmatter on disk: %w", err)
	}

	return nil
}

type ContentRequest struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

func (h *FSHandler) handleModifyContent(e *core.RequestEvent) error {
	req := ContentRequest{}
	if err := e.BindBody(&req); err != nil {
		return fmt.Errorf("invalid json body: %w", err)
	}

	fileRec, err := h.app.FindFirstRecordByData("files", "path", req.Path)
	if err != nil {
		return fmt.Errorf("file not found: %w", err)
	}

	// Keep frontmatter from the record to preserve it
	frontmatterJSON := fileRec.GetString("frontmatter")

	// Update content in the record
	fileRec.Set("content", req.Content)

	absPath := filepath.Join(h.root, fileRec.GetString("path"))

	err = saveToDisk(absPath, req.Content, frontmatterJSON)
	if err != nil {
		return fmt.Errorf("failed to save file on disk: %w", err)
	}

	// Save updated record
	if err := h.app.Save(fileRec); err != nil {
		return fmt.Errorf("failed to save updated file record: %w", err)
	}

	return nil
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
