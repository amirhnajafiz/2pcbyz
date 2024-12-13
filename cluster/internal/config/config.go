package config

import (
	"log"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
)

// Config struct is a module that stores system configs.
type Config struct {
	LogLevel string          `koanf:"log_level"` // log level can be (debug, info, warn, error, panic, fatal)
	Name     string          `koanf:"name"`      // cluster name
	Replicas []ReplicaConfig `koanf:"replicas"`  // types.ReplicaConfig
	Storage  StorageConfig   `koanf:"storage"`   // types.StorageConfig
}

// New reads configuration with koanf,
// by loading a yaml config path into the Config struct.
func New(path string) Config {
	var instance Config

	k := koanf.New(".")

	// load default configuration from file
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	// load configuration from file
	if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
		log.Printf("error loading config.yml: %s", err)
	}

	// unmarshad the instance
	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}

	return instance
}
