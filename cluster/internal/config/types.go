package config

// HandlerConfig stores the handler module config values.
type HandlerConfig struct {
	Instances int `koanf:"instances"`
	QueueSize int `koanf:"queue_size"`
}

// StorageConfig stores the default values for MongoDB cluster.
type StorageConfig struct {
	URI      string `koanf:"uri"`
	Database string `koanf:"database"`
}

// ShardConfig stores the values of the cluster shards.
type ShardConfig struct {
	Cluster string `koanf:"name"`
	From    int    `koanf:"from"`
	To      int    `koanf:"to"`
}

// ReplicaConfig stores the replicas information in cluster to setup.
type ReplicaConfig struct {
	Name string `koanf:"name"`
	Port int    `koanf:"port"`
}
