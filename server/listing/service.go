package listing

// Repository ...
type Repository interface {
	GetTodos() []Todo
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
func (s *Service) GetTodos() []Todo {
	return s.r.GetTodos()
}
