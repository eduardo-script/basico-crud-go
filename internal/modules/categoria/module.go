package categoria

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRegisterModule(db *pgxpool.Pool, router chi.Router) {
	repository := newRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	router.Get("/categorias", handler.ListarCategorias)
	router.Post("/categorias", handler.CadastrarCategorias)
}
