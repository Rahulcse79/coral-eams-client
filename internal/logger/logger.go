package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	*zerolog.Logger
}

var logger *zerolog.Logger

func InitLogger(logFilePath string) *zerolog.Logger {
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	multi := io.MultiWriter(file, consoleWriter)

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
	logger.Info().Fields(fieldsMap(fields)).Msg(msg)
}

func Debug(msg string, fields ...interface{}) {
	logger.Debug().Fields(fieldsMap(fields)).Msg(msg)
}

func Warn(msg string, fields ...interface{}) {
	logger.Warn().Fields(fieldsMap(fields)).Msg(msg)
}

func Error(msg string, fields ...interface{}) {
	logger.Error().Fields(fieldsMap(fields)).Msg(msg)
}

func Fatal(msg string, fields ...interface{}) {
	logger.Fatal().Fields(fieldsMap(fields)).Msg(msg)
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
