package password

import (
	"chatty/chatty/app/config"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestService_Generate(t *testing.T) {
	tests := []struct {
		name       string
		password   string
		wantLength int
		wantErr    bool
	}{
		{
			name:       "Should return particular length string with nil error for non-empty password",
			password:   "qwerty123",
			wantLength: 101,
			wantErr:    false,
		},
	}

	service := NewService(config.Password{
		Secret:      "test!@#123",
		Memory:      1024,
		Iterations:  3,
		SaltLength:  16,
		KeyLength:   32,
		Parallelism: 2,
	})

	for _, tt := range tests {
		convey.Convey(tt.name, t, func() {
			{
				got, err := service.Generate(tt.password)

				convey.So(err != nil, convey.ShouldEqual, tt.wantErr)
				convey.So(got, convey.ShouldHaveLength, tt.wantLength)
			}
		})
	}
}
