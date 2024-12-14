package config

// ShardConfig stores the values of shards creation.
type ShardConfig struct {
	Name    string `koanf:"name"`
	Cluster string `koanf:"cluster"`
	Range   string `koanf:"range"`
}
