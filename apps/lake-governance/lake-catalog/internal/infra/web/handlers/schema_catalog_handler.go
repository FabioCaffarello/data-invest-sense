package handlers

import (
	"encoding/json"
	"net/http"

	"apps/lake-governance/lake-catalog/internal/entity"
	"apps/lake-governance/lake-catalog/internal/usecase"
	inputDTO "libs/dtos/golang/dto-lake-catalog/input"
	"github.com/go-chi/chi/v5"
)

type WebSchemaCatalogHandler struct {
	SchemaCatalogRepository entity.SchemaCatalogInterface
}

func NewWebSchemaCatalogHandler(
	repository entity.SchemaCatalogInterface,
) *WebSchemaCatalogHandler {
	return &WebSchemaCatalogHandler{
		SchemaCatalogRepository: repository,
	}
}

func (wsh *WebSchemaCatalogHandler) CreateSchemaCatalog(w http.ResponseWriter, r *http.Request) {
	var dto inputDTO.SchemaCatalogDTO
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createSchemaCatalog := usecase.NewCreateSchemaCatalogUseCase(
		wsh.SchemaCatalogRepository,
	)

	output, err := createSchemaCatalog.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (wsh *WebSchemaCatalogHandler) ListAllSchemaCatalogs(w http.ResponseWriter, r *http.Request) {
     listAllSchemaCatalogs := usecase.NewListAllSchemaCatalogsUseCase(
          wsh.SchemaCatalogRepository,
     )

     output, err := listAllSchemaCatalogs.Execute()
     if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
     }

     err = json.NewEncoder(w).Encode(output)
     if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
     }
}

func (wsh *WebSchemaCatalogHandler) ListOneSchemaCatalogById(w http.ResponseWriter, r *http.Request) {
     id := chi.URLParam(r, "id")

     listOneSchemaCatalogById := usecase.NewListOneSchemaCatalogByIdUseCase(
          wsh.SchemaCatalogRepository,
     )

     output, err := listOneSchemaCatalogById.Execute(id)
     if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
     }

     err = json.NewEncoder(w).Encode(output)
     if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
     }
}

func (wsh *WebSchemaCatalogHandler) ListAllSchemaCatalogsByService(w http.ResponseWriter, r *http.Request) {
     service := chi.URLParam(r, "service")

     listAllSchemaCatalogsByService := usecase.NewListAllSchemaCatalogsByServiceUseCase(
          wsh.SchemaCatalogRepository,
     )

     output, err := listAllSchemaCatalogsByService.Execute(service)
     if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
     }

     err = json.NewEncoder(w).Encode(output)
     if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
     }
}

func (wsh *WebSchemaCatalogHandler) ListOneSchemaCatalogByServiceAndSource(w http.ResponseWriter, r *http.Request) {
     service := chi.URLParam(r, "service")
     source := chi.URLParam(r, "source")

     listOneSchemaCatalogByServiceAndSource := usecase.NewListOneSchemaCatalogByServiceAndSourceUseCase(
          wsh.SchemaCatalogRepository,
     )

     output, err := listOneSchemaCatalogByServiceAndSource.Execute(service, source)
     if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
     }

     err = json.NewEncoder(w).Encode(output)
     if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
     }
}
