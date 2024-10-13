package config

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func NewZapLogger(cfg *viper.Viper) (*zap.Logger, error) {
	logLevel := cfg.GetString("log.level")
	logDir := cfg.GetString("log.dir")
	logName := cfg.GetString("log.filename")

	var level zap.AtomicLevel

	switch logLevel {
	case `debug`:
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case `warning`:
		level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case `error`:
		level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case `panic`:
		level = zap.NewAtomicLevelAt(zap.PanicLevel)
	case `fatal`:
		level = zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	path := filepath.Join(logDir, logName)
	fmt.Println(path)

	conf := zap.Config{
		Level:            level,
		Encoding:         "json",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout", path},
		ErrorOutputPaths: []string{"stderr", path},
	}

	return conf.Build()
}
