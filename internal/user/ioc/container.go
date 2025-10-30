package ioc

import (
	helloworldioc "github.com/andrewronscki/lib-golang-teste/internal/user/features/hello-world"
	"go.uber.org/dig"
	"go.uber.org/multierr"
)

func Configure(container *dig.Container) error {
	multierr.Combine(
		helloworldioc.Configure(container),
	)

	return nil
}
