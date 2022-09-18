// Package stringutils contains common utility functions for strings.

package stringutils

import "testing"

func TestIsNumber(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Negative Integer",
			args: args{"-100"},
			want: false,
		},
		{
			name: "Positive Integer",
			args: args{"10"},
			want: true,
		},
		{
			name: "Zero",
			args: args{"0"},
			want: true,
		},
		{
			name: "Any Letter",
			args: args{"A"},
			want: false,
		},
		{
			name: "Symbol",
			args: args{"+"},
			want: false,
		},
		{
			name: "Digits Contain Letter",
			args: args{"123A4"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumber(tt.args.str); got != tt.want {
				t.Errorf("IsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseColorHex(t *testing.T) {
	type args struct {
		colorHexValue string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name:    "Hex Color Large",
			args:    args{"A1B7D9"},
			want:    0x00A1B7D9,
			wantErr: false,
		},
		{
			name:    "Hex Color Short",
			args:    args{"F19"},
			want:    0x00FF1199,
			wantErr: false,
		},
		{
			name:    "Hex Large Contains Invalid Digit",
			args:    args{"F2169Z"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "Hex Short Contains Invalid Digit",
			args:    args{"F2Z"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "Invalid Number Of Digits",
			args:    args{"F2A1"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "No 0x Prefix",
			args:    args{"0xF367A1"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "No # Prefix",
			args:    args{"#F367A1"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "Exceeded Maximum Number Of Digits",
			args:    args{"F367A167"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseColorHex(tt.args.colorHexValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseColorHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseColorHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
