package pandati

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatMap(t *testing.T) {
	type args struct {
		nested map[string]interface{}
		opts   *FlattenOptions
	}
	tests := []struct {
		args    args
		wantM   map[string]interface{}
		name    string
		wantErr bool
	}{
		{
			name: "flattening nested map",
			args: args{
				nested: map[string]interface{}{
					"a": map[string]interface{}{
						"b": map[string]interface{}{
							"c": "d",
						},
					},
				},
				opts: &FlattenOptions{
					Separator: ".",
				},
			},
			wantM: map[string]interface{}{
				"a.b.c": "d",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotM, err := FlattenMap(tt.args.nested, tt.args.opts)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.wantM, gotM)
		})
	}
}

func Test_flattenmap(t *testing.T) {
	type args struct {
		nested interface{}
		opts   *FlattenOptions
		prefix string
		depth  int
	}
	tests := []struct {
		args        args
		wantFlatmap map[string]interface{}
		name        string
		wantErr     bool
	}{
		{
			name: "flattening nested map",
			args: args{
				prefix: "",
				depth:  0,
				nested: map[string]interface{}{
					"a": map[string]interface{}{
						"b": map[string]interface{}{
							"c": "d",
						},
					},
				},
				opts: &FlattenOptions{
					Separator: ".",
				},
			},
			wantFlatmap: map[string]interface{}{
				"a.b.c": "d",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFlatmap, err := flattenmap(tt.args.prefix, tt.args.depth, tt.args.nested, tt.args.opts)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.wantFlatmap, gotFlatmap)
		})
	}
}

func Test_updatemap(t *testing.T) {
	type args struct {
		to   map[string]interface{}
		from map[string]interface{}
	}
	tests := []struct {
		args   args
		result map[string]interface{}
		name   string
	}{
		{
			name: "updating nested map",
			args: args{
				to: map[string]interface{}{
					"hi": "there",
				},
				from: map[string]interface{}{
					"foo": "bar",
				},
			},
			result: map[string]interface{}{
				"hi":  "there",
				"foo": "bar",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updatemap(tt.args.to, tt.args.from)
			assert.Equal(t, tt.result, tt.args.to)
		})
	}
}

func TestUnflattenMap(t *testing.T) {
	type args struct {
		flat map[string]interface{}
		opts *FlattenOptions
	}
	tests := []struct {
		args       args
		wantNested map[string]interface{}
		name       string
		wantErr    bool
	}{
		{
			name: "unflattening nested map",
			args: args{
				flat: map[string]interface{}{
					"a.b.c": "d",
				},
				opts: &FlattenOptions{
					Separator: ".",
				},
			},
			wantNested: map[string]interface{}{
				"a": map[string]interface{}{
					"b": map[string]interface{}{
						"c": "d",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNested, err := UnflattenMap(tt.args.flat, tt.args.opts)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.wantNested, gotNested)
		})
	}
}
