package log

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
)

var (
	instance zerolog.Logger
	once     sync.Once
)

func GetLogger() zerolog.Logger {
	once.Do(func() {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		instance = zerolog.New(os.Stdout).With().Timestamp().Logger()
	})
	return instance
}

func WithLayer(layer string) zerolog.Logger {
	return GetLogger().With().Str("layer", layer).Logger()
}
