package password

import (
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
