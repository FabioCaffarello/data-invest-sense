// go:build wireinject
// +build wireinject

package main

import (
     "apps/lake-governance/lake-catalog/internal/entity"
     "apps/lake-governance/lake-catalog/internal/infra/database"
     webHandler "apps/lake-governance/lake-catalog/internal/infra/web/handlers"
     "github.com/google/wire"
     "go.mongodb.org/mongo-driver/mongo"
)

var setSchemaCatalogRepositoryDependency = wire.NewSet(
     database.NewSchemaCatalogRepository,
     wire.Bind(
          new(entity.SchemaCatalogInterface),
          new(*database.SchemaCatalogRepository),
     ),
)

func NewWebSchemaCatalogHandler(client *mongo.Client, database string) *webHandler.WebSchemaCatalogHandler {
     wire.Build(
          setSchemaCatalogRepositoryDependency,
          webHandler.NewWebSchemaCatalogHandler,
     )
     return &webHandler.WebSchemaCatalogHandler{}
}


