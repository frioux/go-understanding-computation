package regex

import "testing"

func TestP91(t *testing.T) {
	pattern := Repeat{
		Concatenate{
			Literal{'a'},
			Choose{
				Empty{}, Literal{'b'},
			},
		},
	}

	shouldMatch := []string{
		"", "a", "ab", "aba", "abab", "abaab",
	}

	for _, v := range shouldMatch {
		if !Matches(pattern, v) {
			t.Errorf("«"+v+"»", "should match")
		}
	}
	if Matches(pattern, "abba") {
		t.Errorf("«abba» should not match")
	}
}
