package config

import (
	"log"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
)

// IPTable holds two maps for the nodes and clusters addresses.
type IPTable struct {
	Endpoints map[string]string `koanf:"endpoints"`
	Services  map[string]string `koanf:"services"`
}

// NewIPTable reads an iptable file inside a IPTable instance.
func NewIPTable(path string) IPTable {
	var instance IPTable

	k := koanf.New(".")

	// load default configuration from file
	if err := k.Load(structs.Provider(DefaultIPTable(), "koanf"), nil); err != nil {
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
