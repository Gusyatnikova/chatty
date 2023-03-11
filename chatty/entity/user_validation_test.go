package entity_test

import (
	"testing"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"

	"chatty/chatty/entity"
)

func TestUser_validate(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		user    entity.User
	}{
		{
			name:    "ok, user is valid",
			wantErr: false,
			user: entity.User{
				ID: ulid.Make(),
				Creds: entity.UserCreds{
					Login:    "TestLogin1",
					Password: "password",
				},
				Contacts: entity.UserContacts{
					Email:       "test@testmail.com",
					PhoneNumber: "+71111111111",
				},
			},
		},
		{
			name:    "nok, user creds are empty",
			wantErr: true,
			user: entity.User{
				ID: ulid.Make(),
				Creds: entity.UserCreds{
					Login:    "",
					Password: "",
				},
				Contacts: entity.UserContacts{
					Email:       "test@testmail.com",
					PhoneNumber: "+71111111111",
				},
			},
		}, {
			name:    "nok, creds and contacts are nil",
			wantErr: true,
			user:    entity.User{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.user.Validate()

			assert.Equal(t, err != nil, tc.wantErr)
		})
	}
}

func TestUserCreds_Validate(t *testing.T) {
	tests := []struct {
		name      string
		wantErr   bool
		userCreds entity.UserCreds
	}{
		{
			name:    "Login with length between 1 and 256 and containing English letters and digits only should be valid",
			wantErr: false,
			userCreds: entity.UserCreds{
				Login:    "TestLogin1",
				Password: "password",
			},
		},
		{
			name:    "Login with zero length should not be valid",
			wantErr: true,
			userCreds: entity.UserCreds{
				Login:    "",
				Password: "password",
			},
		},
		{
			name:    "Login with length greater than 256 should not be valid",
			wantErr: true,
			userCreds: entity.UserCreds{
				Login:    entity.UserLogin(make([]byte, 257)),
				Password: "password",
			},
		},
		{
			name:    "Login contained not only English letters and digits (a-zA-Z0-9) should not be valid",
			wantErr: true,
			userCreds: entity.UserCreds{
				Login:    "Test!@Login1",
				Password: "password",
			},
		},
		{
			name:    "Password with zero length should not be valid",
			wantErr: true,
			userCreds: entity.UserCreds{
				Login:    "TestLogin1",
				Password: "",
			},
		},
		{
			name:    "Password with length greater than 128 should not be valid",
			wantErr: true,
			userCreds: entity.UserCreds{
				Login:    "TestLogin1",
				Password: entity.UserPassword(make([]byte, 129)),
			},
		},
		{
			name:    "Password contained not only printable ASCII characters should not be valid",
			wantErr: true,
			userCreds: entity.UserCreds{
				Login:    "Test1Login1",
				Password: "\n\tpassword",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.userCreds.Validate()
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}
}

func TestUserContacts_Validate(t *testing.T) {
	tests := []struct {
		name         string
		wantErr      bool
		userContacts entity.UserContacts
	}{
		{
			name:    "Phone number in E164 format should be valid",
			wantErr: false,
			userContacts: entity.UserContacts{
				Email:       "test@testmail.com",
				PhoneNumber: "+71111111111",
			},
		},
		{
			name:    "Phone number not in E164 format should not be valid",
			wantErr: true,
			userContacts: entity.UserContacts{
				Email:       "test@testmail.com",
				PhoneNumber: "++71111111111",
			},
		},
		{
			name:    "Correct email address should be valid",
			wantErr: false,
			userContacts: entity.UserContacts{
				Email:       "test@testmail.com",
				PhoneNumber: "+71111111111",
			},
		},
		{
			name:    "Incorrect email address should not be valid",
			wantErr: true,
			userContacts: entity.UserContacts{
				Email:       "usertest@test@mail.com",
				PhoneNumber: "+71111111111",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.userContacts.Validate()
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}
}
