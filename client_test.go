package flaresolverr

import (
	"testing"

	"github.com/google/uuid"
)

func Test_handleSession(t *testing.T) {
	type args struct {
		session uuid.UUID
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid UUID",
			args: args{uuid.MustParse("47d0a203-a007-4a01-b8c1-0cf0156c3cc7")},
			want: "47d0a203-a007-4a01-b8c1-0cf0156c3cc7",
		},
		{
			name: "invalid valid UUID",
			args: args{uuid.Nil},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleSession(tt.args.session); got != tt.want {
				t.Errorf("handleSession() = %v, want %v", got, tt.want)
			}
		})
	}
}
