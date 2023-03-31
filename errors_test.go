package pandati

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckForError(t *testing.T) {
	type args struct {
		err error
		msg []string
	}
	tests := []struct {
		name         string
		wantToReturn string
		args         args
	}{
		{
			name: "Test no error",
			args: args{
				err: nil,
				msg: []string{"Test error"},
			},
			wantToReturn: "",
		},
		{
			name: "Test error",
			args: args{
				err: errors.New("Test error"),
				msg: []string{"Test error"},
			},
			wantToReturn: "Test error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToReturn := CheckForError(tt.args.err, tt.args.msg...)
			assert.Contains(t, gotToReturn, tt.wantToReturn)
		})
	}
}

func TestTrace(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test trace",
			want: "TRACE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Trace()
			assert.Contains(t, got, tt.want)
		})
	}
}
