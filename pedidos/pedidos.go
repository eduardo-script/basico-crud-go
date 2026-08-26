package pedidos

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ItemPedido struct {
	Id         int
	PedidoId   int
	ProdutoId  int
	Quantidade int
	Preco      float64
}

type Pedido struct {
	Id         int
	ClienteId  int
	DataPedido time.Time
	Itens      []ItemPedido
}

// CriarPedido insere o pedido e seus itens no banco
func CriarPedido(db *pgxpool.Pool, clienteId int, itens []ItemPedido) (int, error) {
	ctx := context.Background()

	var pedidoId int
	sqlPedido := `INSERT INTO pedidos (cliente_id) VALUES ($1) RETURNING id`
	err := db.QueryRow(ctx, sqlPedido, clienteId).Scan(&pedidoId)
	if err != nil {
		return 0, err
	}

	sqlItem := `INSERT INTO itens_pedidos (pedido_id, produto_id, quantidade, preco) VALUES ($1, $2, $3, $4)`
	for _, item := range itens {
		_, err := db.Exec(ctx, sqlItem, pedidoId, item.ProdutoId, item.Quantidade, item.Preco)
		if err != nil {
			return 0, err
		}
	}

	return pedidoId, nil
}
