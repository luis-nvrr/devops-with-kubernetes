package main

import "testing"

func TestRandomString(t *testing.T) {
	firstRandom := GetRandomHash()
	secondRandom := GetRandomHash()

	if firstRandom == secondRandom {
		t.Errorf("error got equals, first %q, second %q", firstRandom, secondRandom)
	}
}
