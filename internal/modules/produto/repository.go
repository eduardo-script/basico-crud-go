package produto

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func newRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) addProduto(produto Produto) error {
	sql := `
		INSERT INTO produtos (descricao, preco, quantidade, categoria_id) 
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.Exec(
		context.Background(),
		sql,
		produto.Descricao,
		produto.Preco,
		produto.Quantidade,
		produto.Categoria.Id,
	)
	return err
}

func (r *Repository) listarProdutos() ([]Produto, error) {
	sql := `
		SELECT p.id, p.descricao, p.preco, p.quantidade, c.id, c.nome 
		FROM produtos p
		JOIN categorias c ON c.id = p.categoria_id
		ORDER BY p.id
	`
	linhas, err := r.db.Query(context.Background(), sql)

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	produtos := []Produto{}

	for linhas.Next() {
		var produto Produto

		err := linhas.Scan(
			&produto.Id,
			&produto.Descricao,
			&produto.Preco,
			&produto.Quantidade,
			&produto.Categoria.Id,
			&produto.Categoria.Nome,
		)

		if err != nil {
			return nil, err
		}

		produtos = append(produtos, produto)
	}

	return produtos, nil
}
