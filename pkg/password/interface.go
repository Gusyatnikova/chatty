package password

type Passworder interface {
	Generate(password string) (string, error)
	Compare(password, hash string) (bool, error)
}
