package produto

import "errors"

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) novaProduto(produto Produto) error {

	if produto.Descricao == "" {
		return errors.New("Descricao do produto é obrigatório")
	}

	if produto.Preco == 0.0 {
		return errors.New("Preço do produto é obrigatório")
	}

	if produto.Quantidade == 0 {
		return errors.New("Quantidade do produto é obrigatório")
	}

	if produto.Categoria.Id == 0 {
		return errors.New("Categoria do produto é obrigatório")
	}

	return s.repository.addProduto(produto)
}

func (s *Service) listarTodosProduto() ([]Produto, error) {
	return s.repository.listarProdutos()
}
