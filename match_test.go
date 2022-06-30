package main

import "testing"

func TestMatch(t *testing.T) {
	tt := []struct {
		name, state, deny, word string
		exp                     bool
		err                     error
	}{
		{"AnyWord", "_____", "", "HELLO", true, nil},
		{"UnequalLen", "______", "", "JELLO", false, ErrUnequalLength},
		{"InvalidChar", "?____", "", "MELLO", false, ErrInvalidChar},
		{"MixedMatch", "_ELo_", "", "YELLO", true, nil},
		{"AllYellow", "lozel", "", "ZELLO", true, nil},
		{"AllYellowNoMatch", "hello", "", "HELLO", false, nil},
		{"DenySuccess", "__a_Y", "lernsop", "GAWKY", true, nil},
		{"DenyFailure", "__a_Y", "lernsop", "SPLAY", false, nil},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := IsMatch(tc.state, tc.word, tc.deny)
			if tc.err != nil && err == nil {
				t.Fatalf("expected error %q, got nil", tc.err)
			}
			if tc.err != nil && err != tc.err {
				t.Fatalf("expected error %q, got %q", tc.err, err)
			}
			if err != nil && tc.err == nil {
				t.Fatalf("unexpected error %q", err)
			}
			if got != tc.exp {
				t.Errorf("expected %t got %t", tc.exp, got)
			}
		})
	}
}
