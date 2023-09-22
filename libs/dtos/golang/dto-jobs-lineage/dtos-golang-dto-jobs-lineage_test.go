package dtos_golang_dto_jobs_lineage

import (
	"testing"
)

func TestDtoJobsLineage(t *testing.T) {
	result := DtoJobsLineage("works")
	if result != "DtoJobsLineage works" {
		t.Error("Expected DtoJobsLineage to append 'works'")
	}
}
