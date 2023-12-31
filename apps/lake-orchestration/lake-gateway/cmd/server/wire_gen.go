// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"apps/lake-orchestration/lake-gateway/internal/entity"
	"apps/lake-orchestration/lake-gateway/internal/event"
	"apps/lake-orchestration/lake-gateway/internal/infra/database"
	"apps/lake-orchestration/lake-gateway/internal/infra/web/handlers"
	"apps/lake-orchestration/lake-gateway/internal/usecase"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"libs/golang/events"
)

// Injectors from wire.go:

// [Use Case]
func NewCreateInputUseCase(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database2 string) *usecase.CreateInputUseCase {
	inputRepository := database.NewInputRepository(client, database2)
	inputCreated := event.NewInputCreated()
	createInputUseCase := usecase.NewCreateInputUseCase(inputRepository, inputCreated, eventDispatcher)
	return createInputUseCase
}

func NewUpdateStatusInputUseCase(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database2 string) *usecase.UpdateStatusInputUseCase {
	inputRepository := database.NewInputRepository(client, database2)
	inputUpdated := event.NewInputUpdated()
	updateStatusInputUseCase := usecase.NewUpdateStatusInputUseCase(inputRepository, inputUpdated, eventDispatcher)
	return updateStatusInputUseCase
}

func NewCreateStagingJobUseCase(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database2 string) *usecase.CreateStagingJobUseCase {
	stagingJobRepository := database.NewStagingJobRepository(client, database2)
	createStagingJobUseCase := usecase.NewCreateStagingJobUseCase(stagingJobRepository)
	return createStagingJobUseCase
}

func NewRemoveStagingJobUseCase(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database2 string) *usecase.RemoveStagingJobUseCase {
	stagingJobRepository := database.NewStagingJobRepository(client, database2)
	removeStagingJobUseCase := usecase.NewRemoveStagingJobUseCase(stagingJobRepository)
	return removeStagingJobUseCase
}

func NewListAllByServiceAndSourceUseCase(client *mongo.Client, database2 string) *usecase.ListAllByServiceAndSourceUseCase {
	inputRepository := database.NewInputRepository(client, database2)
	listAllByServiceAndSourceUseCase := usecase.NewListAllByServiceAndSourceUseCase(inputRepository)
	return listAllByServiceAndSourceUseCase
}

func NewListAllByServiceUseCase(client *mongo.Client, database2 string) *usecase.ListAllByServiceUseCase {
	inputRepository := database.NewInputRepository(client, database2)
	listAllByServiceUseCase := usecase.NewListAllByServiceUseCase(inputRepository)
	return listAllByServiceUseCase
}

func NewListOneByIdAndServiceUseCase(client *mongo.Client, database2 string) *usecase.ListOneByIdAndServiceUseCase {
	inputRepository := database.NewInputRepository(client, database2)
	listOneByIdAndServiceUseCase := usecase.NewListOneByIdAndServiceUseCase(inputRepository)
	return listOneByIdAndServiceUseCase
}

func NewListOneStagingJobUsingServiceSourceIdUseCase(client *mongo.Client, database2 string) *usecase.ListOneStagingJobUsingServiceSourceIdUseCase {
	stagingJobRepository := database.NewStagingJobRepository(client, database2)
	listOneStagingJobUsingServiceSourceIdUseCase := usecase.NewListOneStagingJobUsingServiceSourceIdUseCase(stagingJobRepository)
	return listOneStagingJobUsingServiceSourceIdUseCase
}

// [Web Handler]
func NewWebInputStatusHandler(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database2 string) *handlers.WebInputStatusHandler {
	inputRepository := database.NewInputRepository(client, database2)
	inputUpdated := event.NewInputUpdated()
	webInputStatusHandler := handlers.NewWebInputStatusHandler(eventDispatcher, inputRepository, inputUpdated)
	return webInputStatusHandler
}

func NewWebInputHandler(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database2 string) *handlers.WebInputHandler {
	inputRepository := database.NewInputRepository(client, database2)
	inputCreated := event.NewInputCreated()
	webInputHandler := handlers.NewWebInputHandler(eventDispatcher, inputRepository, inputCreated)
	return webInputHandler
}

func NewWebStagingJobHandler(client *mongo.Client, database2 string) *handlers.WebStagingJobHandler {
	stagingJobRepository := database.NewStagingJobRepository(client, database2)
	webStagingJobHandler := handlers.NewWebStagingJobHandler(stagingJobRepository)
	return webStagingJobHandler
}

// wire.go:

var setInputRepositoryDependency = wire.NewSet(database.NewInputRepository, wire.Bind(
	new(entity.InputInterface),
	new(*database.InputRepository),
),
)

var setStagingJobRepositoryDependency = wire.NewSet(database.NewStagingJobRepository, wire.Bind(
	new(entity.StagingJobInterface),
	new(*database.StagingJobRepository),
),
)

var setEventDispatcherDependency = wire.NewSet(events.NewEventDispatcher, event.NewInputCreated, event.NewInputUpdated, wire.Bind(new(events.EventInterface), new(*event.InputCreated)), wire.Bind(new(events.EventInterface), new(*event.InputUpdated)), wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)))

var setInputCreatedEvent = wire.NewSet(event.NewInputCreated, wire.Bind(new(events.EventInterface), new(*event.InputCreated)))

var setInputUpdatedEvent = wire.NewSet(event.NewInputUpdated, wire.Bind(new(events.EventInterface), new(*event.InputUpdated)))
