package http_v1

import "github.com/labstack/echo/v4"

// HealthCheck godoc
// @Summary Return 200 and empty json if service is OK
// @Tags    Health check
// @Produce json
// @Success 200
// @Router  /health [get]
func (e *ServerHandler) HealthCheck(eCtx echo.Context) error {
	return e.uc.HealthCheck(eCtx.Request().Context())
}
