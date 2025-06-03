package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gobwas/glob"
	"github.com/goccy/go-yaml"
	"github.com/pkg/xattr"
	"github.com/pocketbase/pocketbase"
)

func GetSetting(app *pocketbase.PocketBase, key string) string {
	rec, err := app.FindFirstRecordByData("settings", "key", key)
	if err != nil {
		app.Logger().Error("unable to find setting", err)
		return ""
	}
	return rec.GetString("value")
}

func Debounce[T any](input <-chan T, duration time.Duration) <-chan T {
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

type ExtractFrontMatterResult struct {
	FrontMatter string
	MainContent string
}

func ExtractFrontMatter(content string) ExtractFrontMatterResult {
	extracted := ExtractFrontMatterResult{
		MainContent: content,
	}
	// Check if content starts with "---"
	if !strings.HasPrefix(content, "---\n") {
		return extracted
	}

	// Find the closing "---"
	endIndex := strings.Index(content[4:], "---")
	if endIndex == -1 {
		return extracted
	}

	// Adjust for the initial offset "---\n"
	endIndex += 4
	extracted.FrontMatter = content[4:endIndex]
	mainContentOffset := 3
	if len(content) >= endIndex+4 {
		// Catch the last newline after the frontmatter
		mainContentOffset = 4
	}
	extracted.MainContent = content[endIndex+mainContentOffset:]
	return extracted
}

func SaveToDisk(filePath string, content string, rawFrontmatter string) error {
	var b strings.Builder
	if rawFrontmatter != "" {
		b.WriteString("---\n")
		b.WriteString(rawFrontmatter)
		b.WriteString("---\n")
	}
	b.WriteString(content)

	return os.WriteFile(filePath, []byte(b.String()), 0644)
}

func JsonToYaml(jsonRaw string) string {
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

func IsExcluded(patterns []glob.Glob, path string) bool {
	for _, pattern := range patterns {
		if pattern.Match(path) {
			return true
		}
	}
	return false
}

func RemoveFileXAttrVersion(filePath string) error {
	// Remove only the Linux style user namespace key
	return xattr.Remove(filePath, "user.notebase.version")
}

func GetVersion() string {
	return time.Now().UTC().Format(time.RFC3339Nano)
}

type XAttrs struct {
	Version string
	Origin  string
}

func GetXAttrs(filePath string) (XAttrs, error) {
	version, _ := xattr.Get(filePath, "user.notebase.version")
	origin, _ := xattr.Get(filePath, "user.notebase.origin")
	return XAttrs{
		Version: string(version),
		Origin:  string(origin),
	}, nil
}

func SetFileXAttrs(filePath string, xattrs XAttrs) {
	_ = xattr.Set(filePath, "user.notebase.version", []byte(xattrs.Version))
	_ = xattr.Set(filePath, "user.notebase.origin", []byte(xattrs.Origin))
}

func GetDBHash(rawFrontmatter, content string) string {
	h := sha256.New()
	if rawFrontmatter != "" {
		h.Write([]byte("---\n"))
		h.Write([]byte(rawFrontmatter))
		h.Write([]byte("---\n"))
	}
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil))
}

func GetFSHash(filePath string) string {
	result, _ := os.ReadFile(filePath)
	h := sha256.New()
	h.Write(result)
	return hex.EncodeToString(h.Sum(nil))
}
