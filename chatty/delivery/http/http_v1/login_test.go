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
	)

	t.Run("ok, successful login", func(t *testing.T) {
		rec, err := Login(t, reqBody, expectedUser, mockData{
			errUC:   nil,
			mockUC:  true,
			mockJWT: true,
		})

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
			name:        "nok, get error from usecase: unsuccessful authorization",
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
			rec, err := Login(t, tt.reqBody, entity.User{}, mockData{
				errUC:   tt.errFromUC,
				mockUC:  tt.mockUC,
				mockJWT: false,
			})

			require.EqualError(t, err, tt.expectedErr.Error())
			require.Nil(t, actualRespBody(rec))
		})
	}
}

type mockData struct {
	errUC   error
	mockUC  bool
	mockJWT bool
}

func Login(t *testing.T, rb map[string]interface{}, user entity.User, md mockData) (*httptest.ResponseRecorder, error) {
	body, _ := json.Marshal(rb)
	r := httptest.NewRequest(http.MethodGet, "/login", bytes.NewReader(body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	eCtx := getEchoContext(r, rec)

	h := newHandler(mockLogin(t, eCtx.Request().Context(), user, md.mockJWT, md.mockUC, md.errUC))

	return rec, h.Login(eCtx)
}

func mockLogin(t *testing.T, ctx context.Context, user entity.User, mockJWT bool, mockUC bool, ucErr error) (
	muc usecase.ChatUseCase, tm jwt.TokenManager) {
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
		tm = jwtManager
	}

	return
}
