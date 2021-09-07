package pandati

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExistsInSlice(t *testing.T) {
	type args struct {
		slice interface{}
		item  interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Item present in slice",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				item:  3,
			},
			want: true,
		},
		{
			name: "Item not present in slice",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				item:  99,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExistsInSlice(tt.args.slice, tt.args.item)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRemoveFromSlice(t *testing.T) {
	type args struct {
		slice interface{}
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Remove integer from slice",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				value: 3,
			},
			want: []int{1, 2, 0, 4, 5},
		},
		{
			name: "Remove string from slice",
			args: args{
				slice: []string{"a", "b", "c", "d", "e"},
				value: "c",
			},
			want: []string{"a", "b", "", "d", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveFromSlice(tt.args.slice, tt.args.value)
			assert.Equal(t, tt.want, tt.args.slice)
		})
	}
}

func TestRemoveFromSliceByIndex(t *testing.T) {
	type args struct {
		slice interface{}
		index int
	}
	tests := []struct {
		name  string
		args  args
		wants interface{}
	}{
		{
			name: "Remove value with index 3 from slice",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				index: 3,
			},
			wants: []int{1, 2, 3, 0, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveFromSliceByIndex(tt.args.slice, tt.args.index)
			assert.Equal(t, tt.wants, tt.args.slice)
		})
	}
}
