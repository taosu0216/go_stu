package config

type Logger struct {
	Level     string `yaml:"level"`
	Prefix    string `yaml:"prefix"`
	Director  string `yaml:"director"`
	ShowLine  bool   `yaml:"show_line"`
	LogInLine bool   `yaml:"log_in_console"`
}
