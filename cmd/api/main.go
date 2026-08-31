package main

import (
	"context"
	"log"
	"net/http"

	"basico-crud-go/internal/modules/categoria"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	url := "postgres://postgres:102030@localhost:5432/basicodb"
	db, err := pgxpool.New(context.Background(), url)

	if err != nil {
		log.Fatal("Erro de conexão", err)
	}

	defer db.Close()

	// Inicialização das camadas do módulo de categoria
	repo := categoria.NewRepository(db)
	service := categoria.NewService(repo)
	handler := categoria.NewHandler(&service)

	// Configuração do roteador
	r := chi.NewRouter()

	r.Get("/categorias", handler.ListarCategorias)

	log.Println("Servidor rodando na porta :8080")
	http.ListenAndServe(":8080", r)
}
