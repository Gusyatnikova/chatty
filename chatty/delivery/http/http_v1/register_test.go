package http_v1

import (
	"bytes"
	"context"
	"encoding/json"

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

func TestRegister(t *testing.T) {
	testCases := []struct {
		name         string
		expectedErr  error
		expectedUser entity.User
		mockUC       bool
		mockJWT      bool
		errFromUC    error
		reqBody      map[string]interface{}
	}{
		{
			name:        "ok, successful registration",
			expectedErr: nil,
			expectedUser: entity.User{
				Creds: entity.UserCreds{
					Login:    "TestLogin",
					Password: "!@@@key123",
				},
				Contacts: entity.UserContacts{
					Email:       "test22@gmail.com",
					PhoneNumber: "+713333333333",
				},
			},
			mockUC:    true,
			mockJWT:   true,
			errFromUC: nil,
			reqBody: map[string]interface{}{
				"login":        "TestLogin",
				"password":     "!@@@key123",
				"email":        "test22@gmail.com",
				"phone_number": "+713333333333",
			},
		}, {
			name:        "nok, invalid request body: wrong type",
			expectedErr: delivery.ErrBadRequestBody,
			errFromUC:   nil,
			reqBody: map[string]interface{}{
				"login": 12345,
			},
		}, {
			name:        "nok, get error from usecase: user duplication",
			expectedErr: usecase.ErrUserDuplication,
			mockUC:      true,
			errFromUC:   usecase.ErrUserDuplication,
			reqBody: map[string]interface{}{
				"login":        "TestLogin",
				"password":     "!@@@key123",
				"email":        "test22@gmail.com",
				"phone_number": "+713333333333",
			},
		}, {
			name:        "nok, get error from usecase: user not found for response",
			expectedErr: usecase.ErrNoUser,
			mockUC:      true,
			errFromUC:   usecase.ErrNoUser,
			reqBody: map[string]interface{}{
				"login":        "TestLogin",
				"password":     "!@@@key123",
				"email":        "test22@gmail.com",
				"phone_number": "+713333333333",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.reqBody)
			req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()

			eCtx := getEchoContext(req, rec)
			h := newRegisterHandler(mockRegister(
				t, eCtx.Request().Context(), tc.mockUC, tc.mockJWT, tc.expectedUser, tc.errFromUC))

			err := h.Register(eCtx)

			if err != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
				return
			}

			assert.Equal(t, userToRespBody(tc.expectedUser), *actualRespBody(rec))
		})
	}
}

func mockRegister(t *testing.T, ctx context.Context, mockUC bool, mockJWT bool, user entity.User, ucErr error) (
	muc usecase.ChatUseCase, mjwt jwt.TokenManager) {
	if !mockUC && !mockJWT {
		return nil, nil
	}

	if mockUC {
		uc := ucmocks.NewChatUseCase(t)
		uc.On("Register", ctx, mock.Anything, mock.Anything).Return(ucErr)

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

func newRegisterHandler(uc usecase.ChatUseCase, jwtManager jwt.TokenManager) *ServerHandler {
	return &ServerHandler{
		uc:         uc,
		jwtManager: jwtManager,
	}
}

func actualRespBody(r *httptest.ResponseRecorder) *userRespBody {
	rBody := userRespBody{}

	err := json.NewDecoder(r.Body).Decode(&rBody)
	if err != nil {
		return nil
	}

	return &rBody
}

func getEchoContext(r *http.Request, w http.ResponseWriter) echo.Context {
	e := echo.New()

	eCtx := e.NewContext(r, w)
	eCtx.SetPath(r.URL.Path)

	return eCtx
}
