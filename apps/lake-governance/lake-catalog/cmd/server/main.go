package main

import (
	"apps/lake-governance/lake-catalog/internal/infra/web/webserver"
	"context"
	"libs/golang/go-config/configs"
	mongoClient "libs/golang/go-mongodb/client"
	"log"
	"time"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Connect to MongoDB with retries
	mongoDB := getMongoDBClient(configs, ctx)
	client := mongoDB.Client
	defer client.Disconnect(ctx)

     webserver := webserver.NewWebServer(configs.WebServerPort)
     webSchemaCatalogHandler := NewWebSchemaCatalogHandler(client, configs.DBName)

     webserver.AddHandler("/schemas", "POST", "/schemas-catalog", webSchemaCatalogHandler.CreateSchemaCatalog)
     webserver.AddHandler("/schemas", "GET", "/schemas-catalog", webSchemaCatalogHandler.ListAllSchemaCatalogs)
     webserver.AddHandler("/schemas", "GET", "/schemas-catalog/{id}", webSchemaCatalogHandler.ListOneSchemaCatalogById)
     webserver.AddHandler("/schemas", "GET", "/schemas-catalog/service/{service}", webSchemaCatalogHandler.ListAllSchemaCatalogsByService)
     webserver.AddHandler("/schemas", "GET", "/schemas-catalog/service/{service}/source/{source}", webSchemaCatalogHandler.ListOneSchemaCatalogByServiceAndSource)

     log.Printf("Server started at port %s", configs.WebServerPort)
     webserver.Start()

     select {}
}

func getMongoDBClient(configs configs.Config, ctx context.Context) *mongoClient.MongoDB {
	mongoDB := mongoClient.NewMongoDB(
		configs.DBDriver,
		configs.DBUser,
		configs.DBPassword,
		configs.DBHost,
		configs.DBPort,
		configs.DBName,
		ctx,
	)

	_, err := mongoDB.Connect()
	if err != nil {
		panic(err)
	}

	return mongoDB
}
