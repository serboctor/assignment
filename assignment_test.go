package main

import "testing"

func TestCountTokens(t *testing.T) {

	// test empty list
	letters := []string{}
	if count, _ := CountTokens(letters); count != 0 {
		t.Errorf("Test failed for %+q", letters)
	}

	// test tokens with multiple spaces between, before, and after the tokens
	letters = []string{" ", " ", "t", "h", "i", "s", " ", "i", "s", " ", " ", "a", " ", "t", "e", "s", "t", " "}
	if count, _ := CountTokens(letters); count != 4 {
		t.Errorf("Test failed for %+q", letters)
	}

	// test tokens with multiple spaces between the tokens
	letters = []string{"t", "h", "i", "s", " ", "i", "s", " ", " ", "a", " ", "t", "e", "s", "t"}
	if count, _ := CountTokens(letters); count != 4 {
		t.Errorf("Test failed for %+q", letters)
	}

	// test list with no tokens and only spaces
	letters = []string{" ", " ", " "}
	if count, _ := CountTokens(letters); count != 0 {
		t.Errorf("Test failed for %+q", letters)
	}

	// test list with faulty tokens
	letters = []string{"more", " ", "t", "o", "k", "e", "n", "s"}
	if _, err := CountTokens(letters); err == nil {
		t.Errorf("Test failed for %+q", letters)
	}
}
