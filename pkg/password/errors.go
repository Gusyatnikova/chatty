package password

import "github.com/pkg/errors"

var ErrInvalidHash = errors.New("the encoded hash is not in the correct format")
var ErrIncompatibleVersion = errors.New("incompatible version of argon2")
