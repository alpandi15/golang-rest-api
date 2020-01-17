package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var share *_Configuration

type _Configuration struct {
	Server struct {
		Host         string        `json:"host"`
		Port         int           `json:"port"`
		ReadTimeout  time.Duration `json:"read_timeout"`
		WriteTimeout time.Duration `json:"write_timeout"`
	} `json:"server"`

	Log struct {
		Verbose bool `json:"verbose"`
	} `json:"log"`
}

func init() {
	if share != nil {
		return
	}

	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
		return
	}

	bts, err := ioutil.ReadFile(filepath.Join(basePath, "config", "config.json"))
	if err != nil {
		panic(err)
		return
	}

	share = new(_Configuration)
	err = json.Unmarshal(bts, &share)
	if err != nil {
		panic(err)
		return
	}
}

func Configuration() _Configuration {
	return *share
}
