package usecase

import (
     apiClient "libs/api-client/golang/go-outputs"
     outputsInputDTO "libs/dtos/golang/dto-lake-outputs/input"
     outputsOutputDTO "libs/dtos/golang/dto-lake-outputs/output"
)

type CreateOutputUseCase struct {
     OutputsAPIClient apiClient.Client
}

func NewCreateOutputUseCase() *CreateOutputUseCase {
     return &CreateOutputUseCase{
          OutputsAPIClient: *apiClient.NewClient(),
     }
}

func (cou *CreateOutputUseCase) Execute(service string, serviceOutput outputsInputDTO.ServiceOutputDTO) (outputsOutputDTO.ServiceOutputDTO, error) {
     serviceOutputCreated, err := cou.OutputsAPIClient.CreateOutput(service, serviceOutput)
     if err != nil {
          return outputsOutputDTO.ServiceOutputDTO{}, err
     }
     return serviceOutputCreated, nil
}
