package config

func Default() Config {
	return Config{
		Port: 5001,
		Storage: StorageConfig{
			URI:      "mongo://localhost:1272",
			Database: "C0",
		},
	}
}

func DefaultIPTable() IPTable {
	return IPTable{
		Endpoints: make(map[string]string),
		Services:  make(map[string]string),
	}
}
