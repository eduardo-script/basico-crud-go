package main

import (
	"basico-crud-go/clientes"
	"basico-crud-go/pedidos"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	url := "postgres://postgres:102030@localhost:5432/basicodb"
	db, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatal("Erro de conexão:", err)
	}
	defer db.Close()

	// 1. Add novo cliente
	novoCliente := clientes.Cliente{
		Nome:     "João Silva",
		CpfCnpj:  "123.456.789-00",
		Telefone: "(86) 99999-8888",
	}

	err = clientes.AddCliente(db, novoCliente)
	if err != nil {
		log.Println("Erro ao adicionar cliente:", err)
	}

	// 2. Listar clientes
	lista, err := clientes.ListarClientes(db)
	if err != nil {
		log.Fatal("Erro ao listar:", err)
	}

	fmt.Println("--- Lista de Clientes ---")
	for _, c := range lista {
		fmt.Printf("%d - %s | CPF/CNPJ: %s | Tel: %s\n", c.Id, c.Nome, c.CpfCnpj, c.Telefone)
	}

	// 3. Criar um novo pedido para o Cliente ID 1
	fmt.Println("\n--- Criando Pedido ---")
	itens := []pedidos.ItemPedido{
		{ProdutoId: 1, Quantidade: 2, Preco: 750.00},
	}

	pedidoId, err := pedidos.CriarPedido(db, 1, itens)
	if err != nil {
		log.Println("Erro ao criar pedido:", err)
		return
	}

	fmt.Printf("Pedido #%d registrado com sucesso!\n", pedidoId)
}
