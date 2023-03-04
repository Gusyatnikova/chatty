package jwt

type TokenManager interface {
	Generate(sub string) (string, error)
	ExtractSub(token string) (string, error)
}
