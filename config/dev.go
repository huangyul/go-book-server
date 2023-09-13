package config

var Config = WebookConfig{
	DB: DBConfig{
		DSB: "root:root@tcp(localhost:13316)/webook",
	},
	Redis: RedisConfig{
		Addr: "localhost:6379",
	},
}
