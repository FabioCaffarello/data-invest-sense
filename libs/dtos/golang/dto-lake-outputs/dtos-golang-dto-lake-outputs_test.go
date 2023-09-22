package dtos_golang_dto_lake_outputs

import (
	"testing"
)

func TestDtoLakeOutputs(t *testing.T) {
	result := DtoLakeOutputs("works")
	if result != "DtoLakeOutputs works" {
		t.Error("Expected DtoLakeOutputs to append 'works'")
	}
}
