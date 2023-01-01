package password

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
)

func (e Service) Generate(password string) (string, error) {
	salt, err := generateRandomBytes(e.params.saltLength)
	if err != nil {
		return "", err
	}
	secretSalt := append(salt, e.params.secret...)

	hash := argon2.IDKey([]byte(password), secretSalt,
		e.params.iterations, e.params.memory, e.params.parallelism, e.params.keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d,k=%d$%s$%s",
		argon2.Version, e.params.memory, e.params.iterations, e.params.parallelism, e.params.keyLength, b64Salt, b64Hash)

	return encodedHash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
