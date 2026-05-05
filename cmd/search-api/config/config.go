package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// getEnvOrDefault 環境変数が設定されていればその値を、なければデフォルト値を返す
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

type AmazonConfig struct {
	AssociateTag string `yaml:"AssociateTag"`
	AccessKey    string `yaml:"AccessKey"`
	SecretKey    string `yaml:"SecretKey"`
}

func (a *AmazonConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type rawAmazonConfig AmazonConfig
	raw := rawAmazonConfig{}
	if err := unmarshal(&raw); err != nil {
		return err
	}

	a.AssociateTag = getEnvOrDefault("AMAZON_ASSOCIATE_TAG", raw.AssociateTag)
	a.AccessKey = getEnvOrDefault("AMAZON_ACCESS_KEY", raw.AccessKey)
	a.SecretKey = getEnvOrDefault("AMAZON_SECRET_KEY", raw.SecretKey)
	return nil
}

type RakutenConfig struct {
	ApplicationID     string `yaml:"ApplicationID"`
	ApplicationSecret string `yaml:"ApplicationSecret"`
	AffiliateID       string `yaml:"AffiliateID"`
}

func (r *RakutenConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type rawRakutenConfig RakutenConfig
	raw := rawRakutenConfig{}
	if err := unmarshal(&raw); err != nil {
		return err
	}

	r.ApplicationID = getEnvOrDefault("RAKUTEN_APPLICATION_ID", raw.ApplicationID)
	r.ApplicationSecret = getEnvOrDefault("RAKUTEN_APPLICATION_SECRET", raw.ApplicationSecret)
	r.AffiliateID = getEnvOrDefault("RAKUTEN_AFFILIATE_ID", raw.AffiliateID)
	return nil
}

type Config struct {
	Amazon  AmazonConfig  `yaml:"Amazon"`
	Rakuten RakutenConfig `yaml:"Rakuten"`
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
