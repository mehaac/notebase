package config

import (
	"fmt"
	"os"
	"path"

	"github.com/goccy/go-yaml"
)

type NotebaseConfig struct {
	ClearOnStartup bool     `yaml:"clear_on_startup"`
	Exclude        []string `yaml:"exclude"`
	SyncWorkers    int      `yaml:"sync_workers"`
	SyncBatchSize  int      `yaml:"sync_batch_size"`
}

func Load(root string) (NotebaseConfig, error) {
	conf := NotebaseConfig{}
	data, err := os.ReadFile(path.Join(root, ".notebase.yml"))
	if err != nil {
		return conf, fmt.Errorf("error reading config", err)
	}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return conf, fmt.Errorf("error parsing config", err)
	}
	if conf.SyncBatchSize == 0 {
		conf.SyncBatchSize = 200
	}
	if conf.SyncWorkers == 0 {
		conf.SyncWorkers = 5
	}
	return conf, nil
}
