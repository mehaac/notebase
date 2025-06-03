package notebasesync

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/biozz/wow/notebase/internal/utils"
	"github.com/gobwas/glob"
	"github.com/goccy/go-yaml"
	"github.com/pocketbase/pocketbase/core"
	"github.com/syncthing/notify"
)

type File struct {
	JSONFrontmatter string
	RawFrontmatter  string
	Content         string
	AbsPath         string
	RelPath         string
	Slug            string

	// Versioning
	Origin  string
	Version string
}

func parse(rootPath string, curPath string) (File, error) {
	relPath, _ := filepath.Rel(rootPath, curPath)
	fileName := filepath.Base(relPath)
	slug := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	content, err := os.ReadFile(curPath)
	if err != nil {
		return File{}, err
	}

	contentStr := string(content)
	extracted := utils.ExtractFrontMatter(contentStr)
	xattrs, _ := utils.GetXAttrs(curPath)

	data := File{
		AbsPath:         curPath,
		RelPath:         relPath,
		Content:         extracted.MainContent,
		Slug:            slug,
		JSONFrontmatter: "{}",
		RawFrontmatter:  extracted.FrontMatter,
		Version:         xattrs.Version,
		Origin:          xattrs.Origin,
	}

	if len(extracted.FrontMatter) > 0 {
		yamlFrontmatter := yaml.MapSlice{}
		yaml.Unmarshal([]byte(extracted.FrontMatter), &yamlFrontmatter)
		jsonBytes, err := yaml.MarshalWithOptions(yamlFrontmatter, yaml.JSON())
		if err == nil {
			data.JSONFrontmatter = string(jsonBytes)
		}
	}

	data.Content = extracted.MainContent

	return data, nil
}

func (h *SyncHandler) fileWatcher(watcher chan notify.EventInfo, patterns []glob.Glob) {
	for {
		select {
		case ei := <-watcher:
			relPath, _ := filepath.Rel(h.root, ei.Path())
			if utils.IsExcluded(patterns, relPath) {
				continue
			}
			h.app.Logger().Info("file watcher event", "event", ei.Event().String(), "path", ei.Path())
			if filepath.Ext(ei.Path()) != ".md" {
				continue
			}
			if ei.Event() == notify.Create {
				fstat, _ := os.Stat(ei.Path())
				if fstat.IsDir() {
					continue
				}
			}
			h.handleFSNotifyEvent(ei)
		}
	}
}

type WatcherEvent struct {
	EventType string
	Path      string
}

func (h *SyncHandler) handleFSNotifyEvent(event notify.EventInfo) {
	switch event.Event() {
	case notify.Create, notify.Write:
		data, err := parse(h.root, event.Path())
		if err != nil {
			h.app.Logger().Error("unable to parse", err)
			return
		}
		if event.Event() == notify.Create {
			if data.Slug == "Untitled" {
				// this is probably a placeholder file, created by Obsidian, ignore it
				return
			}

			if err := h.createFile(data); err != nil {
				h.app.Logger().Error("unable to create file", err)
			}
			return
		}
		if event.Event() == notify.Write {
			if err := h.updateFile(data); err != nil {
				h.app.Logger().Error("unable to write file", err)
			}
			return
		}
	case notify.Rename, notify.Remove:
		if err := h.softDeleteFile(event.Path()); err != nil {
			h.app.Logger().Error("unable to delete file", err)
		}
	}
}

func (h *SyncHandler) createFile(data File) error {
	filesCol, err := h.app.FindCollectionByNameOrId("files")
	if err != nil {
		return err
	}
	fileRec := core.NewRecord(filesCol)
	version := utils.GetVersion()

	data.Origin = "fs"
	data.Version = version
	fillFileRecFromData(fileRec, data)

	// Remove deleted timestamp if we are restoring the file
	// (important for docker volume changes tracking,
	// where the file is remove and recreated and not updated)
	fileRec.Set("deleted", nil)

	if err := h.app.Save(fileRec); err != nil {
		h.app.Logger().Error("Error saving file record", "error", err)
		return err
	}

	utils.SetFileXAttrs(data.AbsPath, utils.XAttrs{Version: version, Origin: "fs"})

	return nil
}

func fillFileRecFromData(fileRec *core.Record, data File) {
	fileRec.Load(map[string]any{
		"path":            data.RelPath,
		"slug":            data.Slug,
		"content":         data.Content,
		"frontmatter":     data.JSONFrontmatter,
		"raw_frontmatter": data.RawFrontmatter,
		"origin":          data.Origin,
		"version":         data.Version,
	})
}

func (h *SyncHandler) updateFile(data File) error {
	fileRec, err := h.app.FindFirstRecordByData("files", "path", data.RelPath)
	if err != nil {
		return err
	}

	xattrs, _ := utils.GetXAttrs(data.AbsPath)
	dbVersion, _ := time.Parse(time.RFC3339Nano, fileRec.GetString("version"))
	fsVersion, _ := time.Parse(time.RFC3339Nano, xattrs.Version)

	dbHash := utils.GetDBHash(fileRec.GetString("frontmatter"), fileRec.GetString("content"))
	fsHash := utils.GetFSHash(data.AbsPath)

	if dbHash == fsHash {
		h.app.Logger().Debug("file hash is not changed, skipping update", "path", data.RelPath)
		return nil
	}

	if dbVersion.After(fsVersion) {
		h.app.Logger().Debug("file version is not newer, skipping update", "path", data.RelPath)
		return nil
	}

	data.Origin = "fs"
	data.Version = xattrs.Version
	fillFileRecFromData(fileRec, data)

	if err := h.app.Save(fileRec); err != nil {
		h.app.Logger().Error("Error updating file record", "error", err)
		return err
	}

	// Skip writing xattrs here to avoid triggering another FS event
	// utils.SetFileXAttrs(data.AbsPath, utils.XAttrs{Version: data.Version, Origin: data.Origin})

	return nil
}

func (h *SyncHandler) softDeleteFile(path string) error {
	relPath, _ := filepath.Rel(h.root, path)
	fileRec, err := h.app.FindFirstRecordByData("files", "path", relPath)
	if err != nil {
		return err
	}
	fileRec.Set("deleted", time.Now())
	if err := h.app.Save(fileRec); err != nil {
		h.app.Logger().Error("Error deleting file record", "error", err)
		return err
	}

	// Remove xattr notebase.version on delete is optional, ignoring error
	_ = utils.RemoveFileXAttrVersion(path)

	return nil
}
