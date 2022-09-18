// Package utils contains common utility functions used across the application.

package utils

import "testing"

func TestGetRGBComponents(t *testing.T) {
	type args struct {
		color uint64
	}
	tests := []struct {
		name      string
		args      args
		wantRed   uint8
		wantGreen uint8
		wantBlue  uint8
	}{
		{
			name:      "Red component value",
			args:      args{0x00230000},
			wantRed:   0x23,
			wantGreen: 0,
			wantBlue:  0,
		},
		{
			name:      "Green component value",
			args:      args{0x00002300},
			wantRed:   0,
			wantGreen: 0x23,
			wantBlue:  0,
		},
		{
			name:      "Blue component value",
			args:      args{0x00000023},
			wantRed:   0,
			wantGreen: 0,
			wantBlue:  0x23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRed, gotGreen, gotBlue := GetRGBComponents(tt.args.color)
			if gotRed != tt.wantRed {
				t.Errorf("GetRGBComponents() gotRed = %v, want %v", gotRed, tt.wantRed)
			}
			if gotGreen != tt.wantGreen {
				t.Errorf("GetRGBComponents() gotGreen = %v, want %v", gotGreen, tt.wantGreen)
			}
			if gotBlue != tt.wantBlue {
				t.Errorf("GetRGBComponents() gotBlue = %v, want %v", gotBlue, tt.wantBlue)
			}
		})
	}
}

func TestScaleDimension(t *testing.T) {
	type args struct {
		size  int
		scale float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Half Scale",
			args: args{2, 0.5},
			want: 1,
		},
		{
			name: "Double Scale",
			args: args{2, 2},
			want: 4,
		},
		{
			name: "One and Half Scale",
			args: args{2, 1.5},
			want: 3,
		},
		{
			name: "Three times Scale",
			args: args{2, 3},
			want: 6,
		},
		{
			name: "Rounded result with fraction part > .5",
			args: args{2, 1.3},
			want: 3,
		},
		{
			name: "Rounded result with fraction part < .5",
			args: args{2, 1.6},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ScaleDimension(tt.args.size, tt.args.scale); got != tt.want {
				t.Errorf("ScaleDimension() = %v, want %v", got, tt.want)
			}
		})
	}
}
