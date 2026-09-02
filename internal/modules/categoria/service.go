package categoria

import "errors"

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) NovaCategoria(categoria Categoria) error {

	if categoria.Nome == "" {
		return errors.New("Nome de categoria é obrigatória")
	}

	return s.repository.AddCategoria(categoria)
}

func (s *Service) ListarTodasCategorias() ([]Categoria, error) {
	return s.repository.ListarCategorias()
}
