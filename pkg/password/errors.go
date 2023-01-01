package password

import "github.com/pkg/errors"

var ErrInvalidHash = errors.New("The encoded hash is not in the correct format")
var ErrIncompatibleVersion = errors.New("Incompatible version of argon2")
