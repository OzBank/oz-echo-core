package echo

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Echo(input string) string {
	return input
}
