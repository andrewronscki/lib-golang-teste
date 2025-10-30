package logger

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func ConfigureLogger() {
	logger = zerolog.New(os.Stdout).With().Caller().Timestamp().Logger()
}

func Info(ctx context.Context) *zerolog.Event {
	event := logger.Info().Ctx(ctx)

	return event
}

func Warn(ctx context.Context) *zerolog.Event {
	event := logger.Warn().Ctx(ctx)
	return event
}

func Err(ctx context.Context, err error) *zerolog.Event {
	event := logger.Err(err).Ctx(ctx)

	return event
}

func Error(ctx context.Context) *zerolog.Event {
	event := logger.Error().Ctx(ctx)

	return event
}

func Debug(ctx context.Context) *zerolog.Event {
	event := logger.Debug().Ctx(ctx)

	return event
}

func Fatal(ctx context.Context) *zerolog.Event {
	event := logger.Fatal().Ctx(ctx)

	return event
}
