package config

// StorageConfig stores the default values for MongoDB cluster.
type StorageConfig struct {
	URI      string `koanf:"uri"`
	Database string `koanf:"database"`
}
