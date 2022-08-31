package main

import "testing"

func TestExtract(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "test-0",
			arg:  "",
			want: "",
		},
		{
			name: "test-1",
			arg:  "a4bc2d5e",
			want: "aaaabccddddde",
		},
		{
			name: "test-2",
			arg:  "45",
			want: "некорректная строка",
		},
		{
			name: "test-3",
			arg:  "abcd",
			want: "abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Extract(tt.arg); got != tt.want {
				t.Errorf("Extract() = %v, want %v", got, tt.want)
			}
		})
	}
}
