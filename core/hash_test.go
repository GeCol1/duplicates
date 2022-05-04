package core

import (
	"testing"
)

func TestGetHash(t *testing.T) {
	var tests = []struct {
		desc string
		what []byte
		want string
	}{
		{desc: "Hash of simple string", what: []byte("a simple string"), want: "4b26fb3e86544790ce6cc70897df747e4dd12d90"},
		{desc: "Nil input", what: nil, want: "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
	}

	for _, test := range tests {
		got := GetHash(test.what)
		if got != test.want {
			t.Errorf("%s >>> %s expected but got %s\n", test.desc, test.want, got)
		}
	}
}
