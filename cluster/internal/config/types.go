package config

// StorageConfig stores the default values for MongoDB cluster.
type StorageConfig struct {
	URI       string `koanf:"uri"`
	Database  string `koanf:"database"`
	Partition string `koanf:"-"`
}

// ReplicaConfig stores the replicas information in cluster to setup.
type ReplicaConfig struct {
	Name string `koanf:"name"`
	Port int    `koanf:"port"`
}
