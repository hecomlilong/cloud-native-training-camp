package config

import (
	"fmt"
	"net/url"

	"github.com/hecomlilong/cloud-native-training-camp/pkg/utils/errors"
	errPkg "github.com/pkg/errors"
)

type Config interface {
	String(string) string
	Int64(string) int64
}

var DefaultConfig Config

func InitConfig(uri string) error {
	urlObj, err := url.Parse(uri)
	if err != nil {
		return errPkg.Wrap(err, "InitConfig")
	}
	fmt.Printf("urlObj:%v\n", urlObj)
	if urlObj == nil {
		return errors.ErrorParseConfigNil
	}
	switch urlObj.Scheme {
	case "file":
		DefaultConfig, err = NewFileConfig(urlObj.Path)
		if err != nil {
			return errPkg.Wrap(err, "InitConfig")
		}
		return nil
	default:
		return errors.ErrorConfigSchemaNotSupport
	}
}

func GetConfig() Config {
	return DefaultConfig
}
