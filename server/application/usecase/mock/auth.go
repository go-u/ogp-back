package mock

type AuthUsecase struct {
	OnVerify func(jwt string) (uid string, err error)
}

func (u *AuthUsecase) Verify(jwt string) (string, error) {
	return u.OnVerify(jwt)
}
