package pandati

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ComapreStructsReturnDiff(t *testing.T) {
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		args     args
		wantDiff map[string]interface{}
		name     string
	}{
		{
			name: "Identical structs",
			args: args{
				a: struct {
					A int
				}{
					A: 1,
				},
				b: struct {
					A int
				}{
					A: 1,
				},
			},
			wantDiff: map[string]interface{}{},
		},
		{
			name: "Replaced value",
			args: args{
				a: struct {
					A int
				}{
					A: 1,
				},
				b: struct {
					A int
				}{
					A: 2,
				},
			},
			wantDiff: map[string]interface{}{"/A": float64(2)},
		},
		{
			name: "Different structs",
			args: args{
				a: struct {
					A int
				}{
					A: 1,
				},
				b: struct {
					A int
					C int
				}{
					A: 1,
					C: 2,
				},
			},
			wantDiff: map[string]interface{}{},
		},
		{
			name: "Different structs pt 2",
			args: args{
				a: struct {
					A int
				}{
					A: 1,
				},
				b: struct {
					A int
					C int
				}{
					A: 2,
					C: 2,
				},
			},
			wantDiff: map[string]interface{}{"/A": float64(2)},
		},
		{
			name: "Different structs - nested ",
			args: args{
				a: struct {
					A int
					C struct {
						D int
						E bool
					}
				}{
					A: 1,
				},
				b: struct {
					A int
					C struct {
						D int
						E bool
					}
				}{
					A: 1,
					C: struct {
						D int
						E bool
					}{
						D: 3,
						E: true,
					},
				},
			},
			wantDiff: map[string]interface{}{"/C/D": float64(3), "/C/E": true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diff, zero, err := CompareStructsReplaced(tt.args.a, tt.args.b)
			assert.NoError(t, err)
			if !zero {
				mapStringInterface := map[string]interface{}{}
				for _, v := range diff {
					mapStringInterface[v.Key] = v.Value
				}

				assert.Equalf(t, tt.wantDiff, mapStringInterface, "diff: %v", mapStringInterface)
			}
		})
	}
}
