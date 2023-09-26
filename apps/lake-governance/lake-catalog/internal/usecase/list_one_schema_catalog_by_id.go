package usecase

import (
	"apps/lake-governance/lake-catalog/internal/entity"
	outputDTO "libs/dtos/golang/dto-lake-catalog/output"
)

type ListOneSchemaCatalogByIdUseCase struct {
	SchemaCatalogRepository entity.SchemaCatalogInterface
}

func NewListOneSchemaCatalogByIdUseCase(
	repository entity.SchemaCatalogInterface,
) *ListOneSchemaCatalogByIdUseCase {
	return &ListOneSchemaCatalogByIdUseCase{
		SchemaCatalogRepository: repository,
	}
}

func (la *ListOneSchemaCatalogByIdUseCase) Execute(id string) (outputDTO.SchemaCatalogDTO, error) {
	item, err := la.SchemaCatalogRepository.FindOneById(id)
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
