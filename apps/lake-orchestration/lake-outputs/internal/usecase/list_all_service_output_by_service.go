package usecase

import (
	"apps/lake-orchestration/lake-outputs/internal/entity"
	outputDTO "libs/dtos/golang/dto-lake-outputs/output"
	sharedDTO "libs/dtos/golang/dto-lake-outputs/shared"
)

type ListAllServiceOutputByServiceUseCase struct {
	ServiceOutputRepository entity.ServiceOutputInterface
}

func NewListAllServiceOutputByServiceUseCase(
	repository entity.ServiceOutputInterface,
) *ListAllServiceOutputByServiceUseCase {
	return &ListAllServiceOutputByServiceUseCase{
		ServiceOutputRepository: repository,
	}
}

func (la *ListAllServiceOutputByServiceUseCase) Execute(service string) ([]outputDTO.ServiceOutputDTO, error) {
	items, err := la.ServiceOutputRepository.FindAllByService(service)
	if err != nil {
		return []outputDTO.ServiceOutputDTO{}, err
	}
	var result []outputDTO.ServiceOutputDTO
	for _, item := range items {
		dto := outputDTO.ServiceOutputDTO{
			ID:   string(item.ID),
			Data: item.Data,
			Metadata: sharedDTO.Metadata{
				InputId: item.Metadata.InputID,
				Input: sharedDTO.MetadataInput{
					ID:                  item.Metadata.Input.ID,
					Data:                item.Metadata.Input.Data,
					ProcessingId:        item.Metadata.Input.ProcessingId,
					ProcessingTimestamp: item.Metadata.Input.ProcessingTimestamp,
				},
				Service: item.Metadata.Service,
				Source:  item.Metadata.Source,
			},
			Service:   item.Service,
			Source:    item.Source,
			Context:   item.Context,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
		result = append(result, dto)
	}
	return result, nil
}
