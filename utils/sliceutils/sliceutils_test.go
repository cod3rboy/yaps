// Package sliceutils contains common utility functions for go slices.

package sliceutils

import "testing"

func TestContainsString(t *testing.T) {
	type args struct {
		slice       []string
		searchValue string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty Slice",
			args: args{[]string{}, "value"},
			want: false,
		},
		{
			name: "Contains Match",
			args: args{[]string{"value", "value2"}, "value"},
			want: true,
		},
		{
			name: "Does Not Contain Match",
			args: args{[]string{"value1", "value2"}, "value"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsString(tt.args.slice, tt.args.searchValue); got != tt.want {
				t.Errorf("ContainsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
