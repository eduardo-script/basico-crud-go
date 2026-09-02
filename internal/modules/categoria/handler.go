package categoria

import (
	"encoding/json"
	"net/http"
	// "strconv"
	// "github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CadastrarCategorias(
	response http.ResponseWriter,
	request *http.Request,
) {

	var categoria Categoria

	err := json.NewDecoder(request.Body).Decode(&categoria)

	if err != nil {
		http.Error(
			response,
			"JSON invalido",
			http.StatusBadRequest,
		)

		return
	}

	err = h.service.NovaCategoria(categoria)

	if err != nil {
		http.Error(
			response,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	response.WriteHeader(http.StatusCreated)
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
