package entity

type SchemaCatalogInterface interface {
     SaveSchemaCatalog(schemaCatalog *SchemaCatalog) error
     FindOneById(id string) (*SchemaCatalog, error)
     FindAll() ([]*SchemaCatalog, error)
     FindAllByService(service string) ([]*SchemaCatalog, error)
     FindAllByServiceAndSource(service string, source string) ([]*SchemaCatalog, error)
     DeleteOneById(id string) error
     FindOneByServiceAndSource(service string, source string) (*SchemaCatalog, error)
}
