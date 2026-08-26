package produtos

import (
	"context"
	"time"

	"basico-crud-go/categorias"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Produto struct {
	Id         int
	Descricao  string
	Preco      float64
	Quantidade int
	Categoria  categorias.Categoria
	CriadoEm   time.Time
}

func AddProduto(db *pgxpool.Pool, produto Produto) error {
	sql := `
		INSERT INTO produtos (descricao, preco, quantidade, categoria_id)
		VALUES ($1, $2, $3, $4)
	`

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
		SELECT 
			p.id,
			p.descricao,
			p.preco,
			p.quantidade,
			p.criado_em,
			c.id,
			c.nome
		FROM produtos p
		INNER JOIN categorias c ON p.categoria_id = c.id
	`

	linhas, err := db.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var produtos []Produto

	for linhas.Next() {
		var p Produto

		err := linhas.Scan(
			&p.Id,
			&p.Descricao,
			&p.Preco,
			&p.Quantidade,
			&p.CriadoEm,
			&p.Categoria.Id,
			&p.Categoria.Nome,
		)

		if err != nil {
			return nil, err
		}

		produtos = append(produtos, p)
	}

	if err := linhas.Err(); err != nil {
		return nil, err
	}

	return produtos, nil
}
