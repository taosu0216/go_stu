package config

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	System System `yaml:"system"`
	Logger Logger `yaml:"logger"`
}
