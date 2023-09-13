package config

type WebookConfig struct {
	DB    DBConfig
	Redis RedisConfig
}

type DBConfig struct {
	DSB string
}

type RedisConfig struct {
	Addr string
}
