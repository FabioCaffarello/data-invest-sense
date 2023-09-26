package usecase

import (
     "apps/lake-governance/lake-catalog/internal/entity"
     inputDTO "libs/dtos/golang/dto-lake-catalog/input"
     outputDTO "libs/dtos/golang/dto-lake-catalog/output"
)

type CreateSchemaCatalogUseCase struct {
     SchemaCatalogRepository entity.SchemaCatalogInterface
}

func NewCreateSchemaCatalogUseCase(
     repository entity.SchemaCatalogInterface,
) *CreateSchemaCatalogUseCase {
     return &CreateSchemaCatalogUseCase{
          SchemaCatalogRepository: repository,
     }
}

func (ccu *CreateSchemaCatalogUseCase) Execute(schemaCatalog inputDTO.SchemaCatalogDTO) (outputDTO.SchemaCatalogDTO, error) {
     schemaCatalogEntity, err := entity.NewSchemaCatalog(
          schemaCatalog.Service,
          schemaCatalog.Source,
          schemaCatalog.Context,
          schemaCatalog.LakeLayer,
          schemaCatalog.SchemaType,
          schemaCatalog.Catalog,
     )
     if err != nil {
          return outputDTO.SchemaCatalogDTO{}, err
     }

     err = ccu.SchemaCatalogRepository.SaveSchemaCatalog(schemaCatalogEntity)
     if err != nil {
          return outputDTO.SchemaCatalogDTO{}, err
     }

     dto := outputDTO.SchemaCatalogDTO{
          ID:         string(schemaCatalogEntity.ID),
          Service:    schemaCatalogEntity.Service,
          Source:     schemaCatalogEntity.Source,
          Context:    schemaCatalogEntity.Context,
          LakeLayer:  schemaCatalogEntity.LakeLayer,
          SchemaType: schemaCatalogEntity.SchemaType,
          CatalogID:  string(schemaCatalogEntity.CatalogID),
          Catalog:    schemaCatalogEntity.Catalog,
          CreatedAt:  schemaCatalogEntity.CreatedAt,
          UpdatedAt:  schemaCatalogEntity.UpdatedAt,
     }

     return dto, nil
}
