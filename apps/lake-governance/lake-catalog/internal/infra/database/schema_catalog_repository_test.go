package database

import (
     "apps/lake-governance/lake-catalog/internal/entity"
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SchemaCatalogRepositorySuite struct {
	suite.Suite
	log        *log.Logger
	Client     *mongo.Client
	Database   string
	Collection string
	repo       *SchemaCatalogRepository
}

func TestSchemaCatalogRepositorySuite(t *testing.T) {
	suite.Run(t, new(SchemaCatalogRepositorySuite))
}

func (suite *SchemaCatalogRepositorySuite) SetupTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	mongoURI := "mongodb://localhost:27017"
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	suite.Client = client
	suite.Database = "test-database"
	suite.Collection = "test-service"
	suite.log = log.New(log.Writer(), "[SCHEMA-CATALOG-REPOSITORY] ", log.LstdFlags)
	suite.repo = NewSchemaCatalogRepository(suite.Client, suite.Database)
}

func (suite *SchemaCatalogRepositorySuite) TearDownTest() {
	suite.Client.Database(suite.Database).Drop(context.Background())
	err := suite.Client.Disconnect(context.Background())
	suite.NoError(err)
}

func (suite *SchemaCatalogRepositorySuite) TestFindOneById() {
     schemaCatalog := map[string]interface{}{
          "test": "test",
     }
     catalog, err := entity.NewSchemaCatalog("test", "test", "test", "test", "test", schemaCatalog)
     suite.NoError(err)
     err = suite.repo.SaveSchemaCatalog(catalog)
     suite.NoError(err)
     result, err := suite.repo.getOneById(string(catalog.ID))
     suite.NoError(err)
     suite.Equal(catalog, result)
}

func (suite *SchemaCatalogRepositorySuite) TestFindAll() {
     schemaCatalog1 := map[string]interface{}{
          "test": "test1",
     }
     schemaCatalog2 := map[string]interface{}{
          "test": "test2",
     }
     catalog1, err := entity.NewSchemaCatalog("test1", "test1", "test1", "test1", "test1", schemaCatalog1)
     suite.NoError(err)
     catalog2, err := entity.NewSchemaCatalog("test2", "test2", "test2", "test2", "test2", schemaCatalog2)
     suite.NoError(err)

     err = suite.repo.SaveSchemaCatalog(catalog1)
     suite.NoError(err)
     err = suite.repo.SaveSchemaCatalog(catalog2)
     suite.NoError(err)

     result, err := suite.repo.FindAll()
     suite.NoError(err)
     suite.NotNil(result)
     suite.Equal(2, len(result))
}

func (suite *SchemaCatalogRepositorySuite) TestFindAllWhenNoCatalogs() {
     result, err := suite.repo.FindAll()
     suite.NoError(err)
     suite.Nil(result)
     suite.Equal(0, len(result))
}

func (suite *SchemaCatalogRepositorySuite) TestFindAllByService() {
     schemaCatalog1 := map[string]interface{}{
          "test": "test1",
     }
     schemaCatalog2 := map[string]interface{}{
          "test": "test2",
     }
     catalog1, err := entity.NewSchemaCatalog("test1", "test1", "test1", "test1", "test1", schemaCatalog1)
     suite.NoError(err)
     catalog2, err := entity.NewSchemaCatalog("test2", "test2", "test2", "test2", "test2", schemaCatalog2)
     suite.NoError(err)

     err = suite.repo.SaveSchemaCatalog(catalog1)
     suite.NoError(err)
     err = suite.repo.SaveSchemaCatalog(catalog2)
     suite.NoError(err)

     result, err := suite.repo.FindAllByService("test1")
     suite.NoError(err)
     suite.NotNil(result)
     suite.Equal(1, len(result))
}

func (suite *SchemaCatalogRepositorySuite) TestFindAllByServiceWhenNoCatalogs() {
     result, err := suite.repo.FindAllByService("test1")
     suite.NoError(err)
     suite.Nil(result)
     suite.Equal(0, len(result))
}

func (suite *SchemaCatalogRepositorySuite) TestFindAllByServiceAndSource() {
     schemaCatalog1 := map[string]interface{}{
          "test": "test1",
     }
     schemaCatalog2 := map[string]interface{}{
          "test": "test2",
     }
     catalog1, err := entity.NewSchemaCatalog("test1", "test1", "test1", "test1", "test1", schemaCatalog1)
     suite.NoError(err)
     catalog2, err := entity.NewSchemaCatalog("test2", "test2", "test2", "test2", "test2", schemaCatalog2)
     suite.NoError(err)

     err = suite.repo.SaveSchemaCatalog(catalog1)
     suite.NoError(err)
     err = suite.repo.SaveSchemaCatalog(catalog2)
     suite.NoError(err)

     result, err := suite.repo.FindAllByServiceAndSource("test1", "test1")
     suite.NoError(err)
     suite.NotNil(result)
     suite.Equal(1, len(result))
}

func (suite *SchemaCatalogRepositorySuite) TestFindAllByServiceAndSourceWhenNoCatalogs() {
     result, err := suite.repo.FindAllByServiceAndSource("test1", "test1")
     suite.NoError(err)
     suite.Nil(result)
     suite.Equal(0, len(result))
}

func (suite *SchemaCatalogRepositorySuite) TestDeleteOneById() {
     schemaCatalog := map[string]interface{}{
          "test": "test",
     }
     catalog, err := entity.NewSchemaCatalog("test", "test", "test", "test", "test", schemaCatalog)
     suite.NoError(err)
     err = suite.repo.SaveSchemaCatalog(catalog)
     suite.NoError(err)
     err = suite.repo.DeleteOneById(string(catalog.ID))
     suite.NoError(err)
     result, err := suite.repo.FindOneById(string(catalog.ID))
	suite.Error(err)
	suite.Nil(result)
}

func (suite *SchemaCatalogRepositorySuite) TestDeleteOneByIdWhenCatalogDoesNotExist() {
     err := suite.repo.DeleteOneById("test")
     suite.NoError(err)
}

func (suite *SchemaCatalogRepositorySuite) TestSaveSchemaCatalogWhenSchemaCatalogDoesNotExist() {
     //
     schemaCatalog := map[string]interface{}{
          "test": "test",
     }
     catalog, err := entity.NewSchemaCatalog("test", "test", "test", "test", "test", schemaCatalog)
     suite.NoError(err)
     err = suite.repo.SaveSchemaCatalog(catalog)
     suite.NoError(err)
}

func (suite *SchemaCatalogRepositorySuite) TestSaveSchemaCatalogWhenSchemaCatalogDoesExist() {
     schemaCatalog1 := map[string]interface{}{
          "test": "test1",
     }
     schemaCatalog2 := map[string]interface{}{
          "test": "test2",
     }
     catalog1, err := entity.NewSchemaCatalog("test", "test", "test", "test", "test", schemaCatalog1)
     suite.NoError(err)
     catalog2, err := entity.NewSchemaCatalog("test", "test", "test", "test", "test", schemaCatalog2)
     suite.NoError(err)

     err = suite.repo.SaveSchemaCatalog(catalog1)
     suite.NoError(err)
     err = suite.repo.SaveSchemaCatalog(catalog2)
     suite.NoError(err)

     result, err := suite.repo.FindOneById(string(catalog1.ID))
     suite.NoError(err)
     suite.NotNil(result)
     suite.Equal(catalog2, result)
}
