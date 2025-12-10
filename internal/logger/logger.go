package logger

import (
	"io"
	"os"
	"time"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger *zerolog.Logger

func InitLogger(logFilePath string) *zerolog.Logger {
	var multi io.Writer
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	if logFilePath != "" {
		file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		multi = io.MultiWriter(file, consoleWriter)
	} else {
		multi = consoleWriter
	}

	level := zerolog.InfoLevel
	if os.Getenv("LOG_LEVEL") == "DEBUG" {
		level = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(level)
	l := zerolog.New(multi).With().Timestamp().Logger()
	logger = &l
	log.Logger = l

	return logger
}

func Info(msg string, fields ...interface{}) {
	if logger != nil {
		logger.Info().Fields(fieldsMap(fields)).Msg(msg)
	}
}

func Debug(msg string, fields ...interface{}) {
	if logger != nil {
		logger.Debug().Fields(fieldsMap(fields)).Msg(msg)
	}
}

func Warn(msg string, fields ...interface{}) {
	if logger != nil {
		logger.Warn().Fields(fieldsMap(fields)).Msg(msg)
	}
}

func Error(msg string, fields ...interface{}) {
	if logger != nil {
		logger.Error().Fields(fieldsMap(fields)).Msg(msg)
	}
}

func Fatal(msg string, fields ...interface{}) {
	if logger != nil {
		logger.Fatal().Fields(fieldsMap(fields)).Msg(msg)
	}
}

func fieldsMap(fields []interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	for i := 0; i < len(fields)-1; i += 2 {
		key, ok := fields[i].(string)
		if !ok {
			continue
		}
		out[key] = fields[i+1]
	}
	return out
}
