package produto

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

func (h *Handler) carregarProdutos(
	response http.ResponseWriter,
	request *http.Request,
) {

	produtos, err := h.service.listarTodosProduto()

	if err != nil {
		http.Error(
			response,
			"Error ao carregar produtos",
			http.StatusInternalServerError,
		)

		return
	}

	response.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(response).Encode(produtos)
}
