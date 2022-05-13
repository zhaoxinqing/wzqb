package conf

// ConfInfos config info
type ConfigInfos struct {
	DSN   string      `yaml:"dsn"`   // db
	Redis RedisConfig `yaml:"redis"` // redis
}

// RedisConfig redis
type RedisConfig struct {
	Host string `yarm:"host"`
	Pwd  string `yarm:"pwd"`
	DB   int    `yarm:"db"`
}
