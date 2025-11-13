package datastorage

import (
	"testing"
)

func TestFilterLinesCharacters(t *testing.T) {
	tests := []DataStorageManager{
		{
			lines: []string{
				"http://www.cs.cmu.edu/Groups/AI/html/cltl/clm/clm.html",
				"MAC LISP (1967). MIT A.I. Memo No.116A. Available at:",
				"Steele, G. (1984). Common LISP the Language. Chapter 14.2: Concatenating,",
				"sk_test_4f8b1c2d9a6e7b3f5a0c8d2e1f4a7b9c ApiKey",
			},
		},
		{
			lines: []string{
				"28.2 A PROGRAM IN THIS STYLE",
				"#!/usr/bin/env python",
				"#!/usr/bin/env python##",
			},
		},
		{
			lines: []string{
				"",
				"sk_test_4f8b1c2d9a6e7b3f5a0c8d2e1f4a7b9c ApiKey",
			},
		},
		{
			lines: []string{""},
		},
		{
			lines: []string{},
		},
		{
			lines: []string{"hello"},
		},
	}
	expected_results := [][]string{
		{
			"http www cs cmu edu Groups AI html cltl clm clm html",
			"MAC LISP MIT A I Memo No A Available at ",
			"Steele G Common LISP the Language Chapter Concatenating ",
			"sk test f b c d a e b f a c d e f a b c ApiKey",
		},
		{
			"A PROGRAM IN THIS STYLE",
			"usr bin env python",
			"usr bin env python ",
		},
		{
			"",
			"sk test f b c d a e b f a c d e f a b c ApiKey",
		},
		{
			"",
		},
		{},
		{
			"hello",
		},
	}

	for i, test := range tests {
		test.FilterLinesCharacters()
		for j, line := range test.lines {
			if line != expected_results[i][j] {
				t.Fatalf("For test %v\nexpected: %v\ngot: %v", i+1, expected_results[i], test.lines)
			}
		}
	}
}
