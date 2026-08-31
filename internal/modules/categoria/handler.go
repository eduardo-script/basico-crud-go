package categoria

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ListarCategorias(
	response http.ResponseWriter,
	request *http.Request,
) {
	categorias, err := h.service.ListarTodasCategorias()

	if err != nil {
		http.Error(
			response,
			"Error ao carregar categorias",
			http.StatusInternalServerError,
		)
		return
	}

	response.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(response).Encode(categorias)
}
