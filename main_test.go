package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	tests := []struct {
		input string
		exp   int
	}{
		{input: "word1 word2 word3 word4\n", exp: 4},
		{input: "line1\nline2\nline3\n", exp: 3},
		{input: "hello\nworld\n", exp: 2},
	}

	for _, test := range tests {
		b := bytes.NewBufferString(test.input)
		res := count(b, false, false)
		if res != test.exp {
			t.Errorf("Expected %d, got %d instead.\n", test.exp, res)
		}
	}
}

// TestCountLines tests the count function set to count lines
func TestCountLines(t *testing.T) {
	tests := []struct {
		input string
		exp   int
	}{
		{input: "word1 word2 word3 word4\n", exp: 1},
		{input: "line1\nline2\nline3\n", exp: 3},
		{input: "hello\nworld\n", exp: 2},
	}

	for _, test := range tests {
		b := bytes.NewBufferString(test.input)
		res := count(b, true, false)
		if res != test.exp {
			t.Errorf("Expected %d, got %d instead.\n", test.exp, res)
		}
	}
}

// TestCountBytes tests the count function set to count bytes
func TestCountBytes(t *testing.T) {
	tests := []struct {
		input string
		exp   int
	}{
		{input: "word1 word2 word3 word4\n", exp: 24},
		{input: "line1\nline2\nline3\n", exp: 18},
		{input: "hello\nworld\n", exp: 12},
	}

	for _, test := range tests {
		b := bytes.NewBufferString(test.input)
		res := count(b, false, true)
		if res != test.exp {
			t.Errorf("Expected %d, got %d instead.\n", test.exp, res)
		}
	}
}
