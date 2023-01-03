package password

import (
	"chatty/chatty/app/config"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestService_Generate(t *testing.T) {
	service := NewService(config.Password{
		Secret:      "test!@#123",
		Memory:      1024,
		Iterations:  3,
		SaltLength:  16,
		KeyLength:   32,
		Parallelism: 2,
	})

	convey.Convey("Should return particular length string with nil error for non-empty password", t, func() {
		password, wantLength := "qwerty123", 132

		got, err := service.Generate(password)

		convey.So(err, convey.ShouldBeNil)
		convey.So(got, convey.ShouldHaveLength, wantLength)

		convey.Convey("Should return a different string of the same length for the same password", func() {
			gotRegen, err := service.Generate(password)

			convey.So(err, convey.ShouldBeNil)
			convey.So(gotRegen, convey.ShouldHaveLength, wantLength)
			convey.So(gotRegen, convey.ShouldNotEqual, got)
		})
	})
}
