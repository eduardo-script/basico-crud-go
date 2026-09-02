package main

import (
	database "basico-crud-go/infra/databse"
	"basico-crud-go/internal/modules/categoria"
	produto "basico-crud-go/internal/modules/produto"

	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	url := os.Getenv("DATABASE_URL")

	if url == "" {
		log.Fatal("DATABASE_URL não foi definida")
	}

	db, err := database.NewPostgresPool(url)

	if err != nil {
		log.Fatal("Erro de conexão", err)
	}

	defer db.Close()

	router := chi.NewRouter()
	categoria.NewRegisterModule(db, router)
	produto.NewRegisterModule(db, router)

	log.Println(
		"Servidor executando em http://localhost:8082",
	)

	err = http.ListenAndServe(
		":8082",
		router,
	)

	if err != nil {
		log.Fatal(err)
	}
}
