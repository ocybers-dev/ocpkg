package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/kr/pretty"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

var (
	conf *Config
	once sync.Once
)

// Config holds all the configuration sections.
type Config struct {
	OcModule OcModule `yaml:"oc_module"` // oc module configuration
	Mongo    Mongo    `yaml:"mongo"`     // MongoDB configuration
	Redis    Redis    `yaml:"redis"`     // Redis configuration
}

// OcModule holds configuration for a specific module.
type OcModule struct {
	Name     string `yaml:"name"`      // Module name
	Addr     string `yaml:"addr"`      // HTTP address of the module
	LogLevel string `yaml:"log_level"` // Log level for the module
}

// Mongo holds MongoDB connection settings.
type Mongo struct {
	Addr     string `yaml:"addr"`     // MongoDB address
	Username string `yaml:"username"` // MongoDB username
	Password string `yaml:"password"` // MongoDB password
	Database string `yaml:"database"` // MongoDB database name
}

// Redis holds Redis connection settings.
type Redis struct {
	Address  string `yaml:"address"`  // Redis address
	Username string `yaml:"username"` // Redis username
	Password string `yaml:"password"` // Redis password
	DB       int    `yaml:"db"`       // Redis database ID
}

// GetConf returns a singleton configuration instance.
func GetConf(OcModuleName string) *Config {
	once.Do(func() {
		err := initConf(OcModuleName)
		if err != nil {
			panic(fmt.Sprintf("Failed to initialize configuration: %v", err))
		}
	})
	return conf
}

// initConf initializes the configuration by loading from a YAML file.
func initConf(OcModuleName string) error {
	prefix := "conf"
	confFileRelPath := filepath.Join(prefix, OcModuleName, "conf.yaml")
	content, err := os.ReadFile(confFileRelPath)
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}
	conf = new(Config)
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		return fmt.Errorf("error unmarshaling YAML: %w", err)
	}
	if err := validator.Validate(conf); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	pretty.Printf("%+v\n", conf)
	return nil
}
