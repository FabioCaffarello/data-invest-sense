package usecase

import (
     "apps/lake-governance/lake-catalog/internal/entity"
     outputDTO "libs/dtos/golang/dto-lake-catalog/output"
)

type ListOneSchemaCatalogByServiceAndSourceUseCase struct {
     SchemaCatalogRepository entity.SchemaCatalogInterface
}

func NewListOneSchemaCatalogByServiceAndSourceUseCase(
     repository entity.SchemaCatalogInterface,
) *ListOneSchemaCatalogByServiceAndSourceUseCase {
     return &ListOneSchemaCatalogByServiceAndSourceUseCase{
          SchemaCatalogRepository: repository,
     }
}

func (la *ListOneSchemaCatalogByServiceAndSourceUseCase) Execute(service string, source string) (outputDTO.SchemaCatalogDTO, error) {
     item, err := la.SchemaCatalogRepository.FindOneByServiceAndSource(service, source)
     if err != nil {
          return outputDTO.SchemaCatalogDTO{}, err
     }

     dto := outputDTO.SchemaCatalogDTO{
          ID:         string(item.ID),
          Service:    item.Service,
          Source:     item.Source,
          Context:    item.Context,
          LakeLayer:  item.LakeLayer,
          SchemaType: item.SchemaType,
          CatalogID:  string(item.CatalogID),
          Catalog:    item.Catalog,
          CreatedAt:  item.CreatedAt,
          UpdatedAt:  item.UpdatedAt,
     }

     return dto, nil
}

