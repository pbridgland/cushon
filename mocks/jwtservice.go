package mocks

type JWTService struct {
	NextGenerateTokenResult string
	NextGenerateTokenErr    error
}

// GenerateToken implements interfaces.JWTService.
func (j *JWTService) GenerateToken(id int) (string, error) {
	return j.NextGenerateTokenResult, j.NextGenerateTokenErr
}
