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
	testCases := []struct {
		name               string
		expectedStatusCode int
		TTL                int64
		addToCookie        bool
		addToHeader        bool
		isInvalid          bool
	}{
		{
			name:               "ok, token is found in cookie",
			addToCookie:        true,
			addToHeader:        false,
			expectedStatusCode: http.StatusOK,
		}, {
			name:               "ok, token is found in header",
			addToCookie:        false,
			addToHeader:        true,
			expectedStatusCode: http.StatusOK,
		}, {
			name:               "ok, token is found in both header and cookie",
			addToCookie:        true,
			addToHeader:        true,
			expectedStatusCode: http.StatusOK,
		}, {
			name:               "nok, token is not found in both cookie and header",
			addToCookie:        false,
			addToHeader:        false,
			expectedStatusCode: http.StatusUnauthorized,
		}, {
			name:               "nok, token is expired",
			TTL:                -1,
			addToCookie:        true,
			addToHeader:        true,
			expectedStatusCode: http.StatusUnauthorized,
		}, {
			name:               "nok, token is invalid",
			addToCookie:        true,
			addToHeader:        true,
			isInvalid:          true,
			expectedStatusCode: http.StatusUnauthorized,
		},
	}

	baseUser := entity.NewUser(entity.UserCreds{
		Login:    "John Deer",
		Password: "12345",
	}, entity.UserContacts{})

	JWTCfg := config.JWT{
		Sign: "secret",
		TTL:  2,
	}
	e := initEcho(&JWTCfg)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			JWTCfg.TTL = tc.TTL
			token, expAt := generateToken(JWTCfg, *baseUser, tc.isInvalid)

			req := newRequest(token, expAt, tc.addToCookie, tc.addToHeader)
			res := httptest.NewRecorder()

			e.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatusCode, res.Code)
		})
	}
}

func newRequest(token string, expAt time.Time, addToCookie, addToHeader bool) *http.Request {
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	if addToCookie {
		addCookie(req, token, expAt)
	}
	if addToHeader {
		addHeader(req, token)
	}

	return req
}

func initEcho(jwtCfg *config.JWT) *echo.Echo {
	e := echo.New()

	e.Use(JWTHandlerMiddleware(*jwtCfg))

	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	return e
}

func generateToken(JWTCfg config.JWT, user entity.User, isInvalid bool) (string, time.Time) {
	manager := jwtmanager.NewJWTManager(JWTCfg)

	token, expAt, _ := manager.GenerateAccessToken(user)
	if isInvalid {
		token = randString(len(token))
	}

	return token, expAt
}

func addCookie(req *http.Request, token string, expAt time.Time) {
	accessTokenCookie := &http.Cookie{
		Name:     AccessTokenCookieName,
		Value:    token,
		Expires:  expAt,
		HttpOnly: true,
	}
	req.AddCookie(accessTokenCookie)
}

func addHeader(req *http.Request, token string) {
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
