package config

func Default() Config {
	return Config{
		LogLevel: "error",
		Name:     "C0",
		Storage: StorageConfig{
			URI:      "mongo://localhost:1272",
			Database: "C0",
		},
		Replicas: []ReplicaConfig{},
	}
}
