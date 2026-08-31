package categoria

import "errors"

type Service struct {
	repository Repository
}

func NewService(repository *Repository) Service {
	return Service{
		repository: *repository,
	}
}

func (s *Service) NovaCategoria(nome string) error {

	if nome == "" {
		return errors.New("Nome de categoria é obrigatória")
	}

	return s.repository.AddCategoria(nome)
}

func (s *Service) ListarTodasCategorias() ([]Categoria, error) {
	return s.repository.ListarCategorias()
}
