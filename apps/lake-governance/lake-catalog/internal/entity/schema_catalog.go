package entity

import (
	"errors"
	"libs/golang/goid/config"
	uuid "libs/golang/goid/schema/uuid"
	"time"
)

var (
	ErrSchemaCatalogServiceEmpty = errors.New("schema catalog service is empty")
	ErrSchemaCatalogSourceEmpty  = errors.New("schema catalog source is empty")
	ErrSchemaCatalogContextEmpty = errors.New("schema catalog context is empty")
	ErrSchemaCatalogEmpty        = errors.New("schema catalog is empty")
     ErrSchemaCatalogLakeLayerEmpty = errors.New("schema catalog lake layer is empty")
     ErrSchemaCatalogSchemaTypeEmpty = errors.New("schema catalog schema type is empty")
)

type SchemaCatalog struct {
	ID         config.ID              `bson:"id"`
	Service    string                 `bson:"service"`
	Source     string                 `bson:"source"`
	Context    string                 `bson:"context"`
	LakeLayer  string                 `bson:"lake_layer"`
	SchemaType string                 `bson:"schema_type"`
	CatalogID  uuid.ID                `bson:"catalog_id"`
	Catalog    map[string]interface{} `bson:"catalog"`
	CreatedAt  string                 `bson:"created_at"`
	UpdatedAt  string                 `bson:"updated_at"`
}

func NewSchemaCatalog(
     service string,
     source string,
     context string,
     lakeLayer string,
     schemaType string,
     catalog map[string]interface{},
) (*SchemaCatalog, error) {
     catalogId, err := uuid.GenerateSchemaID(schemaType, catalog)
     if err != nil {
          return nil, err
     }

     schemaCatalog := &SchemaCatalog{
          ID:         config.NewID(service, source),
          Service:    service,
          Source:     source,
          Context:    context,
          LakeLayer:  lakeLayer,
          SchemaType: schemaType,
          CatalogID:  catalogId,
          Catalog:    catalog,
          CreatedAt:  time.Now().Format(time.RFC3339),
          UpdatedAt:  time.Now().Format(time.RFC3339),
     }
     err = schemaCatalog.IsSchemaCatalogValid()
     if err != nil {
          return nil, err
     }
     return schemaCatalog, nil
}

func (schemaCatalog *SchemaCatalog) IsSchemaCatalogValid() error {
	if schemaCatalog.Service == "" {
		return ErrSchemaCatalogServiceEmpty
	}
	if schemaCatalog.Source == "" {
		return ErrSchemaCatalogSourceEmpty
	}
	if schemaCatalog.Context == "" {
		return ErrSchemaCatalogContextEmpty
	}
	if len(schemaCatalog.Catalog) == 0 {
		return ErrSchemaCatalogEmpty
	}
     if schemaCatalog.LakeLayer == "" {
          return ErrSchemaCatalogLakeLayerEmpty
     }
     if schemaCatalog.SchemaType == "" {
          return ErrSchemaCatalogSchemaTypeEmpty
     }
     return nil
}
