package config

type Logger struct {
	Level      string `mapstructure:"level"`
	FilePath   string `mapstructure:"file_path"`
	FileName   string `mapstructure:"file_name"`
	Format     string `mapstructure:"format"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}
