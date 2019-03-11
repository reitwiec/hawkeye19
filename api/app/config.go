package app

var defaultConfig = Settings{
	ServerAddr: ":8080",
	DBUsername: "hawkadmin",
	DBPassword: "neverat12",
	DBName:     "hawkeyedb",
	HashKey:    "hawk-secret-1234",
	BlockKey:   "hawk-on-the-blok",
}

func ConfigureInstance() {
	Configuration = Settings{
		ServerAddr: GetEnv("HAWK_SERVER_ADDR", defaultConfig.ServerAddr),
		DBUsername: GetEnv("HAWK_DB_USER", defaultConfig.DBUsername),
		DBPassword: GetEnv("HAWK_DB_PASSWORD", defaultConfig.DBPassword),
		DBName:     GetEnv("HAWK_DB_NAME", defaultConfig.DBName),
		HashKey:    GetEnv("HAWK_HASH_KEY", defaultConfig.HashKey),
		BlockKey:   GetEnv("HAWK_BLOCK_KEY", defaultConfig.BlockKey),
	}
}
