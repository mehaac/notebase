package notebasesync

import (
	"path/filepath"
	"time"

	"github.com/biozz/wow/notebase/internal/utils"
	"github.com/pocketbase/pocketbase/core"
)

func (h *SyncHandler) OnRecordUpdate(record *core.Record) {
	deleted := record.GetString("deleted")
	if deleted != "" {
		h.app.Logger().Debug("file was soft deleted, no need to proceed")
		return
	}
	path := filepath.Join(h.root, record.GetString("path"))
	xattrs, _ := utils.GetXAttrs(path)

	dbVersion, _ := time.Parse(time.RFC3339Nano, record.GetString("version"))
	fsVersion, _ := time.Parse(time.RFC3339Nano, xattrs.Version)

	frontmatterJSON := record.GetString("frontmatter")
	newFrontmatter := utils.JsonToYaml(frontmatterJSON)
	frontmatter := record.GetString("raw_frontmatter")
	if newFrontmatter != frontmatter {
		h.app.Logger().Debug("frontmatter changed, updating it and exiting")
		frontmatter = newFrontmatter
		record.Set("raw_frontmatter", frontmatter)
		if err := h.app.Save(record); err != nil {
			h.app.Logger().Error("error saving record", "error", err)
			return
		}
		return
	}
	content := record.GetString("content")
	dbHash := utils.GetDBHash(frontmatter, content)
	fsHash := utils.GetFSHash(path)

	h.app.Logger().Debug("record update", "path", path, "dbVersion", dbVersion, "fsVersion", fsVersion, "dbHash", dbHash, "fsHash", fsHash)

	if dbVersion == fsVersion && dbHash == fsHash {
		h.app.Logger().Debug("same hashes, no need to update in db")
		return
	}

	err := utils.SaveToDisk(path, content, frontmatter)
	if err != nil {
		h.app.Logger().Error("error saving file from DB", "path", path, "error", err)
		return
	}

	newVersion := utils.GetVersion()
	utils.SetFileXAttrs(path, utils.XAttrs{Version: newVersion, Origin: "unknown"})
	record.Set("origin", "db")
	record.Set("version", newVersion)
	if err := h.app.SaveNoValidate(record); err != nil {
		h.app.Logger().Error("error updating record origin/version", "path", path, "error", err)
	}
}
