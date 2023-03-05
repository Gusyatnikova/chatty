package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"chatty/chatty/app/config"
	"chatty/chatty/entity"
	jwtmanager "chatty/pkg/http/jwt"
)

func TestJWT(t *testing.T) {
	JWTCfg := config.JWT{
		Sign:                  "secret",
		TTL:                   2,
		AccessTokenCookieName: "access-token",
		AccessTokenHeaderName: "Authorization",
		AuthScheme:            "Bearer ",
	}

	testCases := []struct {
		name               string
		TTL                int64
		Delay              time.Duration
		addTokenToCookie   bool
		addTokenToHeader   bool
		expectedStatusCode int
	}{
		{
			name:               "ok, token is found in cookie",
			TTL:                2,
			Delay:              0,
			addTokenToCookie:   true,
			addTokenToHeader:   false,
			expectedStatusCode: http.StatusOK,
		}, {
			name:               "ok, token is found in header",
			TTL:                2,
			Delay:              0,
			addTokenToCookie:   false,
			addTokenToHeader:   true,
			expectedStatusCode: http.StatusOK,
		}, {
			name:               "ok, token is found in both header and cookie",
			TTL:                2,
			Delay:              0,
			addTokenToCookie:   true,
			addTokenToHeader:   true,
			expectedStatusCode: http.StatusOK,
		}, {
			name:               "nok, token is not found in both cookie and header",
			TTL:                2,
			Delay:              0,
			addTokenToCookie:   false,
			addTokenToHeader:   false,
			expectedStatusCode: http.StatusUnauthorized,
		}, {
			name:               "nok, token is expired",
			TTL:                1,
			Delay:              2,
			addTokenToCookie:   true,
			addTokenToHeader:   true,
			expectedStatusCode: http.StatusUnauthorized,
		},
	}

	baseUser := entity.NewUser(entity.UserCreds{
		Login:    "John Deer",
		Password: "12345",
	}, entity.UserContacts{})

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			JWTCfg.TTL = tc.TTL
			e.Use(JWTHandlerMiddleware(JWTCfg))

			jwtManager := jwtmanager.NewJWTManager(JWTCfg)

			req := httptest.NewRequest(http.MethodGet, "/", nil)

			token, expAt, _ := jwtManager.GenerateAccessToken(*baseUser)
			if tc.addTokenToCookie {
				addCookie(req, JWTCfg, token, expAt)
			}
			if tc.addTokenToHeader {
				addHeader(req, JWTCfg, token)
			}

			res := httptest.NewRecorder()

			time.Sleep(time.Second * tc.Delay)

			e.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatusCode, res.Code)
		})
	}
}

func addCookie(req *http.Request, cfg config.JWT, token string, expAt time.Time) {
	accessTokenCookie := &http.Cookie{
		Name:     cfg.AccessTokenCookieName,
		Value:    token,
		Expires:  expAt,
		HttpOnly: true,
	}
	req.AddCookie(accessTokenCookie)
}

func addHeader(req *http.Request, cfg config.JWT, token string) {
	req.Header.Set(cfg.AccessTokenHeaderName, fmt.Sprintf("%s %s", cfg.AuthScheme, token))
}
