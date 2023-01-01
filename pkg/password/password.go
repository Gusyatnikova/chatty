package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"

	"chatty/chatty/app/config"
)

type params struct {
	secret      []byte
	memory      uint32
	iterations  uint32
	saltLength  uint32
	keyLength   uint32
	parallelism uint8
}

type Service struct {
	params params
}

func NewService(cfg config.Password) Passworder {
	return &Service{params: params{
		secret:      []byte(cfg.Secret),
		memory:      cfg.Memory,
		iterations:  cfg.Iterations,
		parallelism: cfg.Parallelism,
		saltLength:  cfg.SaltLength,
		keyLength:   cfg.KeyLength,
	}}
}

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

func (e Service) Compare(password, encodedHash string) (bool, error) {
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	secretSalt := append(salt, e.params.secret...)
	pwdHash := argon2.IDKey([]byte(password), secretSalt, p.iterations, p.memory, p.parallelism, p.keyLength)
	if subtle.ConstantTimeCompare(hash, pwdHash) == 1 {
		return true, nil
	}

	return false, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func decodeHash(encodedHash string) (p *params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d,k=%d", &p.memory, &p.iterations, &p.parallelism, &p.keyLength)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}
