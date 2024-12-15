package config

func Default() Config {
	return Config{
		ResponseLimit: 0,
		Port:          5001,
		Shards:        make([]ShardConfig, 0),
	}
}

func DefaultIPTable() IPTable {
	return IPTable{
		Endpoints: make(map[string]string),
		Services:  make(map[string]string),
	}
}
