package config

import (
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/ini.v1"
)

type FileConfig struct {
	Cfg *ini.File
}

func NewFileConfig(path string) (*FileConfig, error) {
	cfg, err := ini.Load(strings.TrimLeft(path, "/"))
	if err != nil {
		return nil, errors.Wrap(err, "Fail to read file")
	}
	result := new(FileConfig)
	result.Cfg = cfg
	return result, nil
}
func (p *FileConfig) String(k string) string {
	items := strings.Split(k, ".")
	switch len(items) {
	case 0:
		return ""
	case 1:
		return p.Cfg.Section("").Key(items[0]).String()
	default:
		section := strings.Join(items[:len(items)-1], ".")
		return p.Cfg.Section(section).Key(items[len(items)-1]).String()
	}
}

func (p *FileConfig) Int64(k string) int64 {
	items := strings.Split(k, ".")
	switch len(items) {
	case 0:
		return 0
	case 1:
		res, _ := p.Cfg.Section("").Key(items[0]).Int64()
		return res
	default:
		section := strings.Join(items[:len(items)-1], ".")
		res, _ := p.Cfg.Section(section).Key(items[len(items)-1]).Int64()
		return res
	}
}
