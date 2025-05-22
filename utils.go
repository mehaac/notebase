package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gobwas/glob"
	"github.com/goccy/go-yaml"
	"github.com/pocketbase/pocketbase"
)

func getSetting(app *pocketbase.PocketBase, key string) string {
	rec, err := app.FindFirstRecordByData("settings", "key", key)
	if err != nil {
		app.Logger().Error("unable to find setting", err)
		return ""
	}
	return rec.GetString("value")
}

func debounce[T any](input <-chan T, duration time.Duration) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		var (
			timer     *time.Timer
			lastEvent T
		)

		// Helper function to reset/stop the timer safely.
		resetTimer := func() {
			if timer == nil {
				timer = time.NewTimer(duration)
			} else {
				if !timer.Stop() {
					// Drain the channel if needed.
					select {
					case <-timer.C:
					default:
					}
				}
				timer.Reset(duration)
			}
		}

		for {
			var timerC <-chan time.Time
			if timer != nil {
				timerC = timer.C
			}
			select {
			// Read a new event from the watcher.
			case event, ok := <-input:
				if !ok {
					return // Input closed
				}
				lastEvent = event
				resetTimer()
			// The debounce period finished; send out the last event.
			case <-timerC:
				out <- lastEvent
				// After emitting the event, nil the timer so a new one will be created upon a new event.
				timer = nil
			}
		}
	}()
	return out
}

func extractFrontMatter(content string) (frontMatter, mainContent string) {
	// Check if content starts with "---"
	if !strings.HasPrefix(content, "---\n") {
		return "", content
	}

	// Find the closing "---"
	endIndex := strings.Index(content[4:], "---")
	if endIndex == -1 {
		return "", content
	}

	endIndex += 4 // Adjust for the initial offset
	frontMatter = content[4:endIndex]
	mainContent = content[endIndex+3:] // Skip past the closing "---"

	return frontMatter, mainContent
}

func saveToDisk(filePath string, content string, frontmatterJSON string) error {
	yamlFrontmatter := jsonToYaml(frontmatterJSON)
	if yamlFrontmatter == "" {
		yamlFrontmatter = ""
	}
	var b strings.Builder
	if yamlFrontmatter != "" {
		b.WriteString("---\n")
		b.WriteString(yamlFrontmatter)
		b.WriteString("---\n")
	}
	b.WriteString(content)

	return os.WriteFile(filePath, []byte(b.String()), 0644)
}

func jsonToYaml(jsonRaw string) string {
	if jsonRaw == "" {
		return ""
	}
	var yamlData yaml.MapSlice
	err := yaml.Unmarshal([]byte(jsonRaw), &yamlData)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	yamlBytes, err := yaml.Marshal(yamlData)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(yamlBytes)
}

func isExcluded(patterns []glob.Glob, path string) bool {
	for _, pattern := range patterns {
		if pattern.Match(path) {
			return true
		}
	}
	return false
}
