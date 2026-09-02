package produto

import (
	"basico-crud-go/internal/modules/categoria"
	"time"
)

type Produto struct {
	Id         int                 `json:"id"`
	Descricao  string              `json:"descricao"`
	Preco      float64             `json:"preco"`
	Quantidade int                 `json:"quantidade"`
	Categoria  categoria.Categoria `json:"categoria"`
	CriadoEm   time.Time           `json:"criadoEm"`
}
