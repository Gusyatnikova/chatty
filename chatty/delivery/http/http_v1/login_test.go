package http_v1

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"chatty/chatty/delivery"
	"chatty/chatty/entity"
	"chatty/chatty/usecase"
	ucmocks "chatty/chatty/usecase/mocks"
	"chatty/pkg/http/jwt"
	jwtmocks "chatty/pkg/http/jwt/mocks"
)

func TestLogin_Positive(t *testing.T) {
	var (
		expectedUser = entity.User{
			Creds: entity.UserCreds{
				Login: "TestLogin",
			},
			Contacts: entity.UserContacts{
				Email:       "test22@gmail.com",
				PhoneNumber: "+713333333333",
			},
		}
		reqBody = map[string]interface{}{
			"login":    "TestLogin",
			"password": "!@@@key123",
		}
		body, _ = json.Marshal(reqBody)
	)

	req := httptest.NewRequest(http.MethodGet, "/login", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	eCtx := getEchoContext(req, rec)

	t.Run("ok, successful registration", func(t *testing.T) {
		h := newRegisterHandler(mockLogin(
			t, eCtx.Request().Context(), true, true, expectedUser, nil))

		err := h.Login(eCtx)

		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userToRespBody(expectedUser), *actualRespBody(rec))
	})
}

func TestLogin_Negative(t *testing.T) {
	tests := []struct {
		name        string
		expectedErr error
		mockUC      bool
		errFromUC   error
		reqBody     map[string]interface{}
	}{
		{
			name:        "nok, invalid request body: wrong type",
			expectedErr: delivery.ErrBadRequestBody,
			reqBody: map[string]interface{}{
				"login": 12345,
			},
		}, {
			name:        "nok, invalid request body: password don't passed",
			expectedErr: errors.New("Password: cannot be blank."),
			reqBody: map[string]interface{}{
				"login": "TestLogin",
			},
		}, {
			name:        "nok, invalid request body: login don't passed",
			expectedErr: errors.New("Login: cannot be blank."),
			reqBody: map[string]interface{}{
				"password": "!@@@key123",
			},
		}, {
			name:        "nok, get error from usecase: user not found",
			expectedErr: usecase.ErrNoUser,
			errFromUC:   usecase.ErrNoUser,
			mockUC:      true,
			reqBody: map[string]interface{}{
				"login":        "TestLogin",
				"password":     "!@@@key123",
				"email":        "test22@gmail.com",
				"phone_number": "+713333333333",
			},
		}, {
			name:        "nok, get error from usecase: wrong password",
			expectedErr: usecase.ErrUserUnauthorized,
			errFromUC:   usecase.ErrUserUnauthorized,
			mockUC:      true,
			reqBody: map[string]interface{}{
				"login":    "TestLogin",
				"password": "000000",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.reqBody)
			req := httptest.NewRequest(http.MethodGet, "/login", bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()

			eCtx := getEchoContext(req, rec)

			h := newRegisterHandler(mockLogin(
				t, eCtx.Request().Context(), tt.mockUC, false, entity.User{}, tt.errFromUC))

			err := h.Login(eCtx)

			require.EqualError(t, err, tt.expectedErr.Error())
			require.Nil(t, actualRespBody(rec))
		})
	}
}

func mockLogin(t *testing.T, ctx context.Context, mockUC bool, mockJWT bool, user entity.User, ucErr error) (
	muc usecase.ChatUseCase, mjwt jwt.TokenManager) {
	if !mockUC && !mockJWT {
		return nil, nil
	}

	if mockUC {
		uc := ucmocks.NewChatUseCase(t)
		uc.On("Login", ctx, mock.Anything, mock.Anything).Return(ucErr)

		if ucErr == nil {
			uc.On("GetUserByLogin", ctx, mock.Anything).Return(user, nil)
		}

		muc = uc
	}

	if mockJWT {
		jwtManager := jwtmocks.NewTokenManager(t)
		jwtManager.On("GenerateAccessToken", mock.Anything).Return("", time.Now(), nil)
		mjwt = jwtManager
	}

	return
}
