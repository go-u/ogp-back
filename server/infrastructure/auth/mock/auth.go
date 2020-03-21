package mock

type AuthService struct {
	OnVerify func(string) (uid string, err error)
}

func (s *AuthService) Verify(jwt string) (string, error) {
	return s.OnVerify(jwt)
}
