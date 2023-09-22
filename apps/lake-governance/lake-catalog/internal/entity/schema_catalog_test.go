package entity

import (
     "testing"

     "github.com/stretchr/testify/suite"
)

type SchemaCatalogSuite struct {
     suite.Suite
}

func TestSchemaCatalogSuite(t *testing.T) {
     suite.Run(t, new(SchemaCatalogSuite))
}

func (suite *SchemaCatalogSuite) TestNewSchemaCatalogWhenIsANewSchemaCatalog() {
     catalog := map[string]interface{}{
          "test": "test",
     }
     schemaCatalog, err := NewSchemaCatalog("test", "test", "test", "test", "test", catalog)
     suite.NoError(err)
     suite.NotNil(schemaCatalog)
}

func (suite *SchemaCatalogSuite) TestNewSchemaCatalogWhenServiceIsEmpty() {
     catalog := map[string]interface{}{
          "test": "test",
     }
     schemaCatalog, err := NewSchemaCatalog("", "test", "test", "test", "test", catalog)
     suite.Error(err)
     suite.Nil(schemaCatalog)
}

func (suite *SchemaCatalogSuite) TestNewSchemaCatalogWhenSourceIsEmpty() {
     catalog := map[string]interface{}{
          "test": "test",
     }
     schemaCatalog, err := NewSchemaCatalog("test", "", "test", "test", "test", catalog)
     suite.Error(err)
     suite.Nil(schemaCatalog)
}

func (suite *SchemaCatalogSuite) TestNewSchemaCatalogWhenContextIsEmpty() {
     catalog := map[string]interface{}{
          "test": "test",
     }
     schemaCatalog, err := NewSchemaCatalog("test", "test", "", "test", "test", catalog)
     suite.Error(err)
     suite.Nil(schemaCatalog)
}

func (suite *SchemaCatalogSuite) TestNewSchemaCatalogWhenCatalogIsEmpty() {
     catalog := map[string]interface{}{}
     schemaCatalog, err := NewSchemaCatalog("test", "test", "test", "test", "test", catalog)
     suite.Error(err)
     suite.Nil(schemaCatalog)
}

func (suite *SchemaCatalogSuite) TestNewSchemaCatalogWhenCatalogIsNil() {
     schemaCatalog, err := NewSchemaCatalog("test", "test", "test", "test", "test", nil)
     suite.Error(err)
     suite.Nil(schemaCatalog)
}

func (suite *SchemaCatalogSuite) TestNewSchemaCatalogWhenSchemaTypeIsEmpty() {
     catalog := map[string]interface{}{
          "test": "test",
     }
     schemaCatalog, err := NewSchemaCatalog("test", "test", "test", "test", "", catalog)
     suite.Error(err)
     suite.Nil(schemaCatalog)
}

func (suite *SchemaCatalogSuite) TestNewSchemaCatalogWhenLakeLayerIsEmpty() {
     catalog := map[string]interface{}{
          "test": "test",
     }
     schemaCatalog, err := NewSchemaCatalog("test", "test", "test", "", "test", catalog)
     suite.Error(err)
     suite.Nil(schemaCatalog)
}

func (suite *SchemaCatalogSuite) TestNewSchemaCatalogWhenIsValid() {
     catalog := map[string]interface{}{
          "test": "test",
     }
     schemaCatalog, err := NewSchemaCatalog("test", "test", "test", "test", "test", catalog)
     suite.Equal("config-test-test", schemaCatalog.ID)
     suite.Equal("test", schemaCatalog.Service)
     suite.Equal("test", schemaCatalog.Source)
     suite.Equal("test", schemaCatalog.Context)
     suite.Equal("test", schemaCatalog.LakeLayer)
     suite.Equal("test", schemaCatalog.SchemaType)
     suite.Equal("31eb80bb-b6ad-5d13-976f-b01e33e5029a", schemaCatalog.CatalogID)
     suite.Equal(catalog, schemaCatalog.Catalog)
     suite.NoError(err)
     suite.NotNil(schemaCatalog)
}
