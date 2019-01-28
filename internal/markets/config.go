package markets

import "vega/internal/logging"

type Config struct {
	log logging.Logger
	level logging.Level
}

func NewConfig() *Config {
	level := logging.DebugLevel
	logger := logging.NewLogger()
	logger.InitConsoleLogger(level)
	logger.AddExitHandler()
	return &Config{
		log: logger,
		level: level,
	}
}