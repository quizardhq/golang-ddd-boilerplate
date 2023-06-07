//go:build wireinject
// +build wireinject

package main

import (
	"github.com/quizardhq/golang-ddd-boilerplate/configs"
	"github.com/quizardhq/golang-ddd-boilerplate/internal/application"
	"github.com/quizardhq/golang-ddd-boilerplate/internal/domain/repository"
	"github.com/quizardhq/golang-ddd-boilerplate/internal/entrypoint/handler"
	"github.com/quizardhq/golang-ddd-boilerplate/internal/entrypoint/http"
	"github.com/quizardhq/golang-ddd-boilerplate/internal/entrypoint/mapper"
	"github.com/quizardhq/golang-ddd-boilerplate/internal/infrastructure/persistence"

	"github.com/google/wire"
)

func initHttpService() http.HttpService {
	wire.Build(
		configs.ProvideConfig,
		persistence.ProvideDatabase,
		mapper.MapperSet,
		repository.RepositorySet,
		application.ApplicationSet,
		handler.HandlerSet,
		http.HttpSet,
	)

	return nil
}
