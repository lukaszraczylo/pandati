package pandati

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func Test_isZero(t *testing.T) {
// 	type dummyStruct struct {
// 		A int
// 	}

// 	type args struct {
// 		v reflect.Value
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		{
// 			name: "Non-zero int variable",
// 			args: args{
// 				v: reflect.ValueOf(31337),
// 			},
// 			want: false,
// 		},
// 		{
// 			name: "Zero int variable",
// 			args: args{
// 				v: 0,
// 			},
// 			want: true,
// 		},
// 		{
// 			name: "Zero struct",
// 			args: args{
// 				v: reflect.ValueOf(&dummyStruct{}),
// 			},
// 			want: true,
// 		},
// 		{
// 			name: "Struct with member",
// 			args: args{
// 				v: reflect.ValueOf(&dummyStruct{A: 31337}),
// 			},
// 			want: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := IsZero(tt.args.v)
// 			assert.Equal(t, got, tt.want)
// 		})
// 	}
// }

func TestIsZero(t *testing.T) {
	type dummyStruct struct {
		A int
	}

	type args struct {
		v interface{}
	}
	tests := []struct {
		args args
		name string
		want bool
	}{
		{
			name: "Non-zero int variable",
			args: args{
				v: 31337,
			},
			want: false,
		},
		{
			name: "Zero int variable",
			args: args{
				v: 0,
			},
			want: true,
		},
		{
			name: "Zero struct",
			args: args{
				v: dummyStruct{},
			},
			want: true,
		},
		{
			name: "Zero struct ptr",
			args: args{
				v: &dummyStruct{},
			},
			want: true,
		},
		{
			name: "Struct with member",
			args: args{
				v: dummyStruct{A: 31337},
			},
			want: false,
		},
		{
			name: "Struct ptr with member",
			args: args{
				v: &dummyStruct{A: 31337},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsZero(tt.args.v)
			assert.Equal(t, tt.want, got)
		})
	}
}
