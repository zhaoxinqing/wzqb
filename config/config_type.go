package config

// ConfigInformation config info
type ConfigInformation struct {
	DB struct {
		Host   string `yaml:"host"`
		User   string `yaml:"user"`
		Pwd    string `yaml:"pwd"`
		DBname string `yaml:"dbname"`
	} `yaml:"db"`
}
