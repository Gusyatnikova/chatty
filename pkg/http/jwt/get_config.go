package jwt

import "chatty/chatty/app/config"

func (e *JWTManager) GetConfig() config.JWT {
	return e.cfg
}
