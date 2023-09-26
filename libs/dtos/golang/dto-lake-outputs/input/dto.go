package input

import (
	sharedDTO "libs/dtos/golang/dto-lake-outputs/shared"
)

type ServiceOutputDTO struct {
	Data     map[string]interface{} `json:"data"`
	Metadata sharedDTO.Metadata     `json:"metadata"`
     Context  string                 `json:"context"`
}
