package config

func Default() Config {
	return Config{
		LogLevel: "error",
		Name:     "C0",
		Handler: HandlerConfig{
			Instances: 0,
			QueueSize: 0,
		},
		Storage: StorageConfig{
			URI:      "mongo://localhost:1272",
			Database: "C0",
		},
		Shard: ShardConfig{
			Name:  "D0",
			Range: make([]int, 0),
		},
		Replicas: []ReplicaConfig{},
	}
}

func DefaultIPTable() IPTable {
	return IPTable{
		Endpoints: make(map[string]string),
		Services:  make(map[string]string),
	}
}
