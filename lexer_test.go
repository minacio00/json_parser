package main

import (
	"regexp"
	"testing"
)

func TestStringPattern(t *testing.T) {
	re := regexp.MustCompile(String)
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "simple alpha numeric", input: `"abc123"`, want: true},
		{name: "escaped newline", input: `"line\nbreak"`, want: true},
		{name: "escaped quote", input: `"quote\"here"`, want: true},
		{name: "escaped slash", input: `"slash\/here"`, want: true},
		{name: "unicode escape", input: `"uni\u00A9ok"`, want: true},
		{name: "empty string", input: `""`, want: true},
		{name: "invalid escape", input: `"bad\c"`, want: false},
		{name: "invalid unicode hex", input: `"bad\u12GZ"`, want: false},
		{name: "missing closing quote", input: `"unterminated`, want: false},
		{name: "no surrounding quotes", input: `abc`, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := re.MatchString(tt.input)
			if got != tt.want {
				t.Fatalf("MatchString(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
