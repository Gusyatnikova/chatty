package password

import (
	"chatty/chatty/app/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Compare(t *testing.T) {
	type args struct {
		password    string
		getHashFunc func(string) (string, error)
	}

	service := NewService(config.Password{
		Secret:      "test!@#123",
		Memory:      1024,
		Iterations:  3,
		SaltLength:  16,
		KeyLength:   32,
		Parallelism: 2,
	})
	staticHash := func(string) (string, error) {
		return "$argon2id$v=19$m=1024,t=3,p=2,k=32$V+ZWY8CBYXxIioFhD/x82Q$HDgcrDEr3AWPYf0KFukF0YxvMz6f+XjDwJU4EjCwFmM", nil
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Should return true and nil err if hash is generated from password",
			args: args{
				password:    "testPwd(1)9=+",
				getHashFunc: service.Generate,
			},
			want:    true,
			wantErr: false,
		}, {
			name: "Should return false and NOT nil err if hash is NOT generated from password",
			args: args{
				password:    "testPwd(1)9=+",
				getHashFunc: staticHash,
			},
			want:    false,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedHash, _ := tt.args.getHashFunc(tt.args.password)

			got, err := service.Compare(tt.args.password, encodedHash)

			assert.Equal(t, err != nil, tt.wantErr)
			assert.Equal(t, got, tt.want)
		})
	}
}
