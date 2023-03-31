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
		args args
		name string
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
		args args
		want interface{}
		name string
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
		wants interface{}
		args  args
		name  string
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

func TestUniqueSlice(t *testing.T) {
	type args struct {
		slice interface{}
	}
	tests := []struct {
		args args
		want interface{}
		name string
	}{
		{
			name: "Unique slice of integers",
			args: args{
				slice: []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Unique slice of strings",
			args: args{
				slice: []string{"a", "b", "c", "d", "e", "a", "b", "c", "d", "e", "g"},
			},
			want: []string{"a", "b", "c", "d", "e", "g"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, UniqueSlice(tt.args.slice))
		})
	}
}
