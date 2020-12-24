package listing

import "todo/domain"

// Repository ...
type Repository interface {
	GetTodos() []domain.Todo
}

// Service ...
type Service struct {
	r Repository
}

// NewService ...
func NewService(r Repository) *Service {
	return &Service{r}
}

// GetTodos ...
func (s *Service) GetTodos() []domain.Todo {
	return s.r.GetTodos()
}
