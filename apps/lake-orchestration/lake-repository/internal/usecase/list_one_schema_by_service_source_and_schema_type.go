package usecase

import (
     "apps/lake-orchestration/lake-repository/internal/entity"
     outputDTO "libs/dtos/golang/dto-repository/output"
)

type ListOneSchemaByServiceSourceAndSchemaTypeUseCase struct {
     SchemaRepository entity.SchemaInterface
}

func NewListOneSchemaByServiceSourceAndSchemaTypeUseCase(
     repository entity.SchemaInterface,
) *ListOneSchemaByServiceSourceAndSchemaTypeUseCase {
     return &ListOneSchemaByServiceSourceAndSchemaTypeUseCase{
          SchemaRepository: repository,
     }
}

func (la *ListOneSchemaByServiceSourceAndSchemaTypeUseCase) Execute(service string, source string, schemaType string) (outputDTO.SchemaDTO, error) {
     item, err := la.SchemaRepository.FindOneByServiceSourceAndSchemaType(service, source, schemaType)
     if err != nil {
          return outputDTO.SchemaDTO{}, err
     }

     dto := outputDTO.SchemaDTO{
          ID:         string(item.ID),
          SchemaType: item.SchemaType,
          Service:    item.Service,
          Source:     item.Source,
          JsonSchema: item.JsonSchema,
          SchemaID:   string(item.SchemaID),
          CreatedAt:  item.CreatedAt,
          UpdatedAt:  item.UpdatedAt,
     }

     return dto, nil
}

