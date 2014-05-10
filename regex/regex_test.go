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

	for i := 0; i < len(shouldMatch); i++ {
		str := shouldMatch[i]
		if !Matches(pattern, str) {
			t.Errorf("«"+str+"»", "should match")
		}
	}
	if Matches(pattern, "abba") {
		t.Errorf("«abba» should not match")
	}
}
