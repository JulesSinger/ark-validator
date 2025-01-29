package application

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) Start() error {
	return nil
}

func (s *Service) Stop() error {
	return nil
}
