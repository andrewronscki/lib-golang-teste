package main

import (
	"context"

	"github.com/andrewronscki/lib-golang-teste/internal/app/ioc"

	"github.com/andrewronscki/lib-golang-teste/pkg/config"
	"github.com/andrewronscki/lib-golang-teste/pkg/datadog/logger"
	"github.com/andrewronscki/lib-golang-teste/pkg/hosting"
)

// @title Balance API
// @version 2.0
// @description Code styles Golang API Example
// @termsOfService http://swagger.io/terms/

// @contact.name André
// @contact.url https://www.devtoolshq.dev
// @contact.email contato@devtoolshq.dev

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
// @schemes http
const devEnvPath = "../config/dev.env"

func main() {
	config.LoadEnv("", devEnvPath)

	logger.ConfigureLogger()

	if _, err := ioc.Configure(); err != nil {
		logger.Fatal(context.Background()).AnErr("error", err).Send()
	}

	host := &hosting.Host{
		Addr: config.Env.GetString("ADDR"),
	}

	host.Start()
}
