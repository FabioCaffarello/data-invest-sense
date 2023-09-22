package database

import (
	"apps/lake-governance/lake-catalog/internal/entity"
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SchemaCatalogRepository struct {
	log        *log.Logger
	Client     *mongo.Client
	Database   string
	Collection *mongo.Collection
}

func NewSchemaCatalogRepository(client *mongo.Client, database string) *SchemaCatalogRepository {
	return &SchemaCatalogRepository{
		log:        log.New(os.Stdout, "[SCHEMA-CATALOG-REPOSITORY] ", log.LstdFlags),
		Client:     client,
		Database:   database,
		Collection: client.Database(database).Collection("catalog"),
	}
}

func (scr *SchemaCatalogRepository) getOneById(id string) (*entity.SchemaCatalog, error) {
	filter := bson.M{"id": id}
	existingDoc := scr.Collection.FindOne(context.Background(), filter)
	// Check if the document does not exist
	if existingDoc.Err() != nil {
		return nil, existingDoc.Err()
	}

	var result entity.SchemaCatalog
	if err := existingDoc.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (scr *SchemaCatalogRepository) SaveSchemaCatalog(schemaCatalog *entity.SchemaCatalog) error {
	// Check if the document already exists based on the ID
	existingSchemaCatalog, err := scr.getOneById(string(schemaCatalog.ID))
	if err != nil {
		// Insert new document
		schemaCatalog.CreatedAt = time.Now().Format(time.RFC3339)
		_, err := scr.Collection.InsertOne(context.Background(), bson.M{
			"id":          schemaCatalog.ID,
			"service":     schemaCatalog.Service,
			"source":      schemaCatalog.Source,
			"context":     schemaCatalog.Context,
			"lake_layer":  schemaCatalog.LakeLayer,
			"schema_type": schemaCatalog.SchemaType,
			"catalog_id":  schemaCatalog.CatalogID,
			"catalog":     schemaCatalog.Catalog,
			"created_at":  schemaCatalog.CreatedAt,
			"updated_at":  schemaCatalog.UpdatedAt,
		})
		if err != nil {
			return err
		}
	} else {
		// Update existing document
		schemaCatalog.UpdatedAt = time.Now().Format(time.RFC3339)
		_, err := scr.Collection.UpdateOne(context.Background(), bson.M{
			"id": existingSchemaCatalog.ID,
		}, bson.M{
			"$set": bson.M{
				"service":     schemaCatalog.Service,
				"source":      schemaCatalog.Source,
				"context":     schemaCatalog.Context,
				"lake_layer":  schemaCatalog.LakeLayer,
				"schema_type": schemaCatalog.SchemaType,
				"catalog_id":  schemaCatalog.CatalogID,
				"catalog":     schemaCatalog.Catalog,
				"updated_at":  schemaCatalog.UpdatedAt,
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (scr *SchemaCatalogRepository) FindOneById(id string) (*entity.SchemaCatalog, error) {
	result, err := scr.getOneById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (scr *SchemaCatalogRepository) FindAll() ([]*entity.SchemaCatalog, error) {
     cursor, err := scr.Collection.Find(context.Background(), bson.M{})
     if err != nil {
          return nil, err
     }
     defer cursor.Close(context.Background())

     var results []*entity.SchemaCatalog
     for cursor.Next(context.Background()) {
          var result entity.SchemaCatalog
          if err := cursor.Decode(&result); err != nil {
               return nil, err
          }
          results = append(results, &result)
     }
     if err := cursor.Err(); err != nil {
          return nil, err
     }

     return results, nil
}

func (scr *SchemaCatalogRepository) FindAllByService(service string) ([]*entity.SchemaCatalog, error) {
     filter := bson.M{"service": service}
     cursor, err := scr.Collection.Find(context.Background(), filter)
     if err != nil {
          return nil, err
     }
     defer cursor.Close(context.Background())

     var results []*entity.SchemaCatalog
     for cursor.Next(context.Background()) {
          var result entity.SchemaCatalog
          if err := cursor.Decode(&result); err != nil {
               return nil, err
          }
          results = append(results, &result)
     }

     return results, nil
}

func (scr *SchemaCatalogRepository) FindAllByServiceAndSource(service string, source string) ([]*entity.SchemaCatalog, error) {
     filter := bson.M{"service": service, "source": source}
     cursor, err := scr.Collection.Find(context.Background(), filter)
     if err != nil {
          return nil, err
     }
     defer cursor.Close(context.Background())

     var results []*entity.SchemaCatalog
     for cursor.Next(context.Background()) {
          var result entity.SchemaCatalog
          if err := cursor.Decode(&result); err != nil {
               return nil, err
          }
          results = append(results, &result)
     }

     return results, nil
}

func (scr *SchemaCatalogRepository) DeleteOneById(id string) error {
     filter := bson.M{"id": id}
     _, err := scr.Collection.DeleteOne(context.Background(), filter)
     if err != nil {
          return err
     }
     return nil
}

func (scr *SchemaCatalogRepository) FindOneByServiceAndSource(service string, source string) (*entity.SchemaCatalog, error) {
     filter := bson.M{"service": service, "source": source}
     existingDoc := scr.Collection.FindOne(context.Background(), filter)
     // Check if the document does not exist
     if existingDoc.Err() != nil {
          return nil, existingDoc.Err()
     }

     var result entity.SchemaCatalog
     if err := existingDoc.Decode(&result); err != nil {
          return nil, err
     }

     return &result, nil
}
