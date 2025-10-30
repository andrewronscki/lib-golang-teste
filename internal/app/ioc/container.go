package ioc

import (
	"github.com/andrewronscki/lib-golang-teste/internal/app/behaviors"
	userioc "github.com/andrewronscki/lib-golang-teste/internal/user/ioc"
	"go.uber.org/dig"
	"go.uber.org/multierr"
)

func Configure() (*dig.Container, error) {
	container := dig.New()

	return container, multierr.Combine(
		behaviors.Configure(container),
		userioc.Configure(container),
	)
}
