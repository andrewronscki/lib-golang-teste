package ioc

import (
	"github.com/andrewronscki/lib-golang-teste/internal/app/behaviors"
	"go.uber.org/dig"
	"go.uber.org/multierr"
)

func Configure() (*dig.Container, error) {
	container := dig.New()

	return container, multierr.Combine(
		behaviors.Configure(container),
	)
}
