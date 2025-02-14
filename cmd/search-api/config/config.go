package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Rakuten RakutenConfig `yaml:"Rakuten"`
}

type RakutenConfig struct {
	ApplicationID     string `yaml:"ApplicationID"`
	ApplicationSecret string `yaml:"ApplicationSecret"`
	AffiliateID       string `yaml:"AffiliateID"`
}

func Load(path string) (*Config, error) {
	// 設定ファイルを読み込む
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := yaml.Unmarshal(buf, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
