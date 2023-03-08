package middleware

import (
	"fmt"
	"math/rand"
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
		Sign: "secret",
		TTL:  2,
	}

	testCases := []struct {
		name               string
		expectedStatusCode int
		TTL                int64
		genInvalidToken    bool
		addTokenToCookie   bool
		addTokenToHeader   bool
	}{
		{
			name:               "ok, token is found in cookie",
			TTL:                2,
			addTokenToCookie:   true,
			addTokenToHeader:   false,
			expectedStatusCode: http.StatusOK,
		}, {
			name:               "ok, token is found in header",
			TTL:                2,
			addTokenToCookie:   false,
			addTokenToHeader:   true,
			expectedStatusCode: http.StatusOK,
		}, {
			name:               "ok, token is found in both header and cookie",
			TTL:                2,
			addTokenToCookie:   true,
			addTokenToHeader:   true,
			expectedStatusCode: http.StatusOK,
		}, {
			name:               "nok, token is not found in both cookie and header",
			TTL:                2,
			addTokenToCookie:   false,
			addTokenToHeader:   false,
			expectedStatusCode: http.StatusUnauthorized,
		}, {
			name:               "nok, token is expired",
			TTL:                -1,
			addTokenToCookie:   true,
			addTokenToHeader:   true,
			expectedStatusCode: http.StatusUnauthorized,
		}, {
			name:               "nok, token is invalid",
			TTL:                1,
			addTokenToCookie:   true,
			addTokenToHeader:   true,
			genInvalidToken:    true,
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
			if tc.genInvalidToken {
				token = randString(len(token))
			}
			if tc.addTokenToCookie {
				addCookie(req, JWTCfg, token, expAt)
			}
			if tc.addTokenToHeader {
				addHeader(req, JWTCfg, token)
			}

			res := httptest.NewRecorder()

			e.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatusCode, res.Code)
		})
	}
}

func addCookie(req *http.Request, cfg config.JWT, token string, expAt time.Time) {
	accessTokenCookie := &http.Cookie{
		Name:     AccessTokenCookieName,
		Value:    token,
		Expires:  expAt,
		HttpOnly: true,
	}
	req.AddCookie(accessTokenCookie)
}

func addHeader(req *http.Request, cfg config.JWT, token string) {
	req.Header.Set(AccessTokenHeaderName, fmt.Sprintf("%s %s", AuthScheme, token))
}

func randString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789."

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}
