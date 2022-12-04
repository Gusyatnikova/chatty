package entity_test

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"

	"chatty/chatty/entity"
)

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
		convey.Convey(tt.name, t, func() {
			err := tt.userCreds.Validate()
			convey.So(err != nil, convey.ShouldEqual, tt.wantErr)
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
		convey.Convey(tt.name, t, func() {
			err := tt.userContacts.Validate()
			convey.So(err != nil, convey.ShouldEqual, tt.wantErr)
		})
	}
}
