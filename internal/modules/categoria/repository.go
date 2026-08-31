package categoria

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) AddCategoria(nome string) error {

	sql := `
		INSERT INTO categorias (nome)
		VALUES ($1)
	`
	_, err := r.db.Exec(
		context.Background(),
		sql,
		nome,
	)

	return err
}

func (r *Repository) ListarCategorias() ([]Categoria, error) {
	sql := `
		SELECT id, nome FROM categorias
	`
	linhas, err := r.db.Query(context.Background(), sql)

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

func (r *Repository) BuscarCategoriaPeloId(idCategoria int) (Categoria, error) {

	var categoria Categoria

	sql := `
		SELECT id, nome FROM categorias WHERE id = $1
	`
	err := r.db.QueryRow(
		context.Background(),
		sql,
		idCategoria,
	).Scan(
		&categoria.Id,
		&categoria.Nome,
	)

	return categoria, err
}

func (r *Repository) AtualizarCategoria(idCategoria int, novoNome string) error {
	sql := `
		UPDATE categorias SET nome = $1 WHERE id = $2
	`
	_, err := r.db.Exec(
		context.Background(),
		sql,
		novoNome,
		idCategoria,
	)

	return err
}

func (r *Repository) DeletarCategoria(idCategoria int) error {
	sql := `
		DELETE FROM categorias WHERE id = $1
	`
	_, err := r.db.Exec(
		context.Background(),
		sql,
		idCategoria,
	)

	return err
}
