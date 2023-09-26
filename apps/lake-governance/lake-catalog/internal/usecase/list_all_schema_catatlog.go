package usecase

import (
	"apps/lake-governance/lake-catalog/internal/entity"
	outputDTO "libs/dtos/golang/dto-lake-catalog/output"
)

type ListAllSchemaCatalogsUseCase struct {
	SchemaCatalogRepository entity.SchemaCatalogInterface
}

func NewListAllSchemaCatalogsUseCase(
	repository entity.SchemaCatalogInterface,
) *ListAllSchemaCatalogsUseCase {
	return &ListAllSchemaCatalogsUseCase{
		SchemaCatalogRepository: repository,
	}
}

func (la *ListAllSchemaCatalogsUseCase) Execute() ([]outputDTO.SchemaCatalogDTO, error) {
	items, err := la.SchemaCatalogRepository.FindAll()
	if err != nil {
		return []outputDTO.SchemaCatalogDTO{}, err
	}
	var result []outputDTO.SchemaCatalogDTO
	for _, item := range items {
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
		result = append(result, dto)
	}
	return result, nil
}
