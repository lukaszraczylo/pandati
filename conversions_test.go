package pandati

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestConvertReplyType(t *testing.T) {
	type args struct {
		desired_reply_type interface{}
		reply              []byte
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "TestConvertReplyType - slice of int",
			args: args{
				desired_reply_type: []int{},
				reply:              []byte("[1,2,3]"),
			},
			want: []int{1, 2, 3},
		},
		{
			name: "TestConvertReplyType - struct",
			args: args{
				desired_reply_type: struct {
					Name string
					Age  int
				}{},
				reply: []byte("{\"Name\":\"test\",\"Age\":10}"),
			},
			want: struct {
				Name string
				Age  int
			}{
				Name: "test",
				Age:  10,
			},
		},
		{
			name: "TestConvertReplyType - string",
			args: args{
				desired_reply_type: "",
				reply:              []byte("\"test\""),
			},
			want: "test",
		},
		{
			name: "TestConvertReplyType - []float64",
			args: args{
				desired_reply_type: []float64{},
				reply:              []byte("[1.1,2.2,3.3]"),
			},
			want: []float64{1.1, 2.2, 3.3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertReplyType(tt.args.desired_reply_type, tt.args.reply)
			assert.Equal(t, got, tt.want)
		})
	}
}
