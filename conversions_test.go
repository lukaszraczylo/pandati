package pandati

import (
	"reflect"
	"testing"
)

func TestStringToBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Test string to bytes",
			args: args{
				s: "Test string",
			},
			want: []byte{84, 101, 115, 116, 32, 115, 116, 114, 105, 110, 103},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToBytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesToString(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		want string
		args args
	}{
		{
			name: "Test bytes to string",
			args: args{
				b: []byte{84, 101, 115, 116, 32, 115, 116, 114, 105, 110, 103},
			},
			want: "Test string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToString(tt.args.b); got != tt.want {
				t.Errorf("BytesToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
