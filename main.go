package main

import (
	"context"
	"fmt"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
	p "basico-crud-go/produtos"
)

func main() {
	url := "postgres://postgres:102030@localhost:5432/basicodb"
	db, err := pgxpool.New(context.Background(), url)

	if err != nil {
		log.Fatal("Erro de conexão", err)
	}

	defer db.Close()

	// Criar uma categoria
	// err = p.AddCategoria(db, "Informatica")
	
	// categoria, err := p.ListarCategorias(db)
	categoria, err := p.BuscarCategoriaPeloId(db, 1)

	if err != nil {
		log.Fatal("Error.: ", err)
	}

	// for _, categoria := range categorias {
	// 	fmt.Println(
	// 		categoria.Id,
	// 		categoria.Nome,
	// 	)
	// }

	// fmt.Println("Categoria registrada com sucess o!!!" )
	fmt.Printf("%d - %s\n", categoria.Id, categoria.Nome)
}