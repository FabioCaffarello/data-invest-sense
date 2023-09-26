package output

import (
	sharedDTO "libs/dtos/golang/dto-lake-outputs/shared"
)

type ServiceOutputDTO struct {
	ID        string                 `json:"id"`
	Data      map[string]interface{} `json:"data"`
	Metadata  sharedDTO.Metadata     `json:"metadata"`
	Service   string                 `json:"service"`
	Source    string                 `json:"source"`
	Context   string                 `json:"context"`
	CreatedAt string                 `json:"created_at"`
	UpdatedAt string                 `json:"updated_at"`
}
