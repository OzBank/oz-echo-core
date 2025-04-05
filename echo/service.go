package echo

type EchoService struct {
}

func NewEchoService() *EchoService {
	return &EchoService{}
}

func (s *EchoService) Echo(input string) string {
	return input
}
