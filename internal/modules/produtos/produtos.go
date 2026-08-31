package produtos

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Declare a struct Categoria aqui para a struct Produto conseguir usá-la
type Categoria struct {
	Id   int
	Nome string
}

type Produto struct {
	Id         int
	Descricao  string
	Preco      float64
	Quantidade int
	Categoria  Categoria
	CriadoEm   time.Time
}

func AddProduto(db *pgxpool.Pool, produto Produto) error {
	sql := `
        INSERT INTO produtos (descricao, preco, quantidade, categoria_id)
        VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(
		context.Background(),
		sql,
		produto.Descricao,
		produto.Preco,
		produto.Quantidade,
		produto.Categoria.Id,
	)
	return err
}

func ListarProdutos(db *pgxpool.Pool) ([]Produto, error) {
	sql := `
        SELECT p.id, p.descricao, p.preco, p.quantidade, c.id, c.nome
        FROM produtos p
        JOIN categorias c ON c.id = p.categoria_id
        ORDER BY p.id`

	linhas, err := db.Query(context.Background(), sql)
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
