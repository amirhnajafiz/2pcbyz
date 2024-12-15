package config

func Default() Config {
	return Config{
		LogLevel: "debug",
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
			Cluster: "C0",
			From:    0,
			To:      0,
		},
		Shards:   make([]ShardConfig, 0),
		Replicas: make([]ReplicaConfig, 0),
	}
}

func DefaultIPTable() IPTable {
	return IPTable{
		Endpoints: make(map[string]string),
		Services:  make(map[string]string),
	}
}
