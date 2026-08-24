package produtos

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Categoria struct {
	Id int
	Nome string
}

func AddCategoria(db *pgxpool.Pool, nome string) error {
	
	sql := `
		INSERT INTO categorias (nome)
		VALUES ($1)
	`
	_, err := db.Exec(
		context.Background(),
		sql,
		nome,
	)

	return err
}

func ListarCategorias(db *pgxpool.Pool)([]Categoria, error) {
	sql := `
		SELECT id, nome FROM categorias
	`
	linhas, err := db.Query(context.Background(), sql)

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	categorias := []Categoria{}

	for linhas.Next() {
		var categoria Categoria

		err := linhas.Scan(
			&categoria.Id,
			&categoria.Nome,
		) 

		if err != nil {
			return nil, err
		}

		categorias = append(categorias, categoria)
	}

	return categorias, nil
}

func BuscarCategoriaPeloId(db *pgxpool.Pool, idCategoria int)(Categoria, error) {
	
	var categoria Categoria

	sql := `
		SELECT id, nome FROM categorias WHERE id = $1
	`
	err := db.QueryRow(
		context.Background(),
		sql,
		idCategoria,
	).Scan(
		&categoria.Id,
		&categoria.Nome,
	)

	return categoria, err
}