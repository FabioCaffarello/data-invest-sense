// go:build wireinject
// +build wireinject

package main

import (
	"apps/lake-orchestration/lake-outputs/internal/entity"
	"apps/lake-orchestration/lake-outputs/internal/infra/database"
	webHandler "apps/lake-orchestration/lake-outputs/internal/infra/web/handlers"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var setServiceOutputRepositoryDependency = wire.NewSet(
	database.NewServiceOutputRepository,
	wire.Bind(
		new(entity.ServiceOutputInterface),
		new(*database.ServiceOutputRepository),
	),
)

func NewWebServiceOutputHandler(client *mongo.Client, database string) *webHandler.WebServiceOutputHandler {
	wire.Build(
		setServiceOutputRepositoryDependency,
		webHandler.NewWebServiceOutputHandler,
	)
	return &webHandler.WebServiceOutputHandler{}
}
