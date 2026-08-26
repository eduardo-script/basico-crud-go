package clientes

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// A struct deve começar com 'C' maiúsculo
type Cliente struct {
	Id           int
	Nome         string
	CpfCnpj      string
	Telefone     string
	DataCadastro time.Time
}

func AddCliente(db *pgxpool.Pool, cliente Cliente) error {
	sql := `INSERT INTO clientes (nome, cpf_cnpj, telefone) VALUES ($1, $2, $3)`
	_, err := db.Exec(context.Background(), sql, cliente.Nome, cliente.CpfCnpj, cliente.Telefone)
	return err
}

func ListarClientes(db *pgxpool.Pool) ([]Cliente, error) {
	sql := `SELECT id, nome, cpf_cnpj, telefone, data_cadastro FROM clientes`
	linhas, err := db.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var clientes []Cliente
	for linhas.Next() {
		var c Cliente
		err := linhas.Scan(&c.Id, &c.Nome, &c.CpfCnpj, &c.Telefone, &c.DataCadastro)
		if err != nil {
			return nil, err
		}
		clientes = append(clientes, c)
	}
	return clientes, nil
}
