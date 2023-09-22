package dtos_golang_dto_treatment_domain_registry

import (
	"testing"
)

func TestDtoTreatmentDomainRegistry(t *testing.T) {
	result := DtoTreatmentDomainRegistry("works")
	if result != "DtoTreatmentDomainRegistry works" {
		t.Error("Expected DtoTreatmentDomainRegistry to append 'works'")
	}
}
