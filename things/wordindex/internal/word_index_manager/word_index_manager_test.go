package wordindexmanager

import (
	"errors"
	"slices"
	"testing"
)

type testForm struct {
	word_index_manager WordindexManager
	lines              []string
	lines_per_page     int
}

type resultForm struct {
	words_indexes []WordIndex
	err           error
}

func Test_ExtractWordsAndIndexes_SortListAlphabetically(t *testing.T) {
	tests := []testForm{
		{
			word_index_manager: WordindexManager{
				words_indexes: []WordIndex{},
			},
			lines: []string{
				"http www cs cmu edu Groups AI html cltl clm clm html",
				"MAC LISP MIT A I Memo No A Available at ",
				"Steele G Common LISP the Language Chapter Concatenating ",
				"sk test f b c d a e b f a c d e f a b c ApiKey",
				"A PROGRAM IN THIS STYLE",
				"usr bin env python",
				"usr bin env python ",
			},
			lines_per_page: 2,
		},
		{
			word_index_manager: WordindexManager{
				words_indexes: []WordIndex{},
			},
			lines: []string{
				"http www cs cmu edu Groups AI html cltl clm clm html",
				"MAC LISP MIT A I Memo No A Available at ",
				"Steele G Common LISP the Language Chapter Concatenating ",
				"sk test f b c d a e b f a c d e f a b c ApiKey",
				"A PROGRAM IN THIS STYLE",
				"usr bin env python",
				"usr bin env python ",
			},
			lines_per_page: -1,
		},
		{
			word_index_manager: WordindexManager{
				words_indexes: []WordIndex{},
			},
			lines: []string{
				"http www cs cmu edu Groups AI html cltl clm clm html",
				"MAC LISP MIT A I Memo No A Available at ",
				"Steele G Common LISP the Language Chapter Concatenating ",
				"sk test f b c d a e b f a c d e f a b c ApiKey",
				"A PROGRAM IN THIS STYLE",
				"usr bin env python",
				"usr bin env python ",
			},
			lines_per_page: 0,
		},
		{
			word_index_manager: WordindexManager{
				words_indexes: []WordIndex{},
			},
			lines: []string{
				"http www cs cmu edu Groups AI html cltl clm clm html",
			},
			lines_per_page: 2,
		},
		{
			word_index_manager: WordindexManager{
				words_indexes: []WordIndex{},
			},
			lines: []string{
				"",
			},
			lines_per_page: 2,
		},
		{
			word_index_manager: WordindexManager{
				words_indexes: []WordIndex{},
			},
			lines:          []string{},
			lines_per_page: 2,
		},
		{
			word_index_manager: WordindexManager{
				words_indexes: []WordIndex{},
			},
			lines: []string{
				"hello",
			},
			lines_per_page: 2,
		},
		{
			word_index_manager: WordindexManager{
				words_indexes: []WordIndex{},
			},
			lines: []string{
				"hello ",
			},
			lines_per_page: 2,
		},
	}
	expected_results := []resultForm{
		{
			words_indexes: []WordIndex{
				{Word: "A", Occurrences: 3, Indexes: []int{1, 3}},
				{Word: "AI", Occurrences: 1, Indexes: []int{1}},
				{Word: "ApiKey", Occurrences: 1, Indexes: []int{2}},
				{Word: "Available", Occurrences: 1, Indexes: []int{1}},
				{Word: "Chapter", Occurrences: 1, Indexes: []int{2}},
				{Word: "Common", Occurrences: 1, Indexes: []int{2}},
				{Word: "Concatenating", Occurrences: 1, Indexes: []int{2}},
				{Word: "G", Occurrences: 1, Indexes: []int{2}},
				{Word: "Groups", Occurrences: 1, Indexes: []int{1}},
				{Word: "I", Occurrences: 1, Indexes: []int{1}},
				{Word: "IN", Occurrences: 1, Indexes: []int{3}},
				{Word: "LISP", Occurrences: 2, Indexes: []int{1, 2}},
				{Word: "Language", Occurrences: 1, Indexes: []int{2}},
				{Word: "MAC", Occurrences: 1, Indexes: []int{1}},
				{Word: "MIT", Occurrences: 1, Indexes: []int{1}},
				{Word: "Memo", Occurrences: 1, Indexes: []int{1}},
				{Word: "No", Occurrences: 1, Indexes: []int{1}},
				{Word: "PROGRAM", Occurrences: 1, Indexes: []int{3}},
				{Word: "STYLE", Occurrences: 1, Indexes: []int{3}},
				{Word: "Steele", Occurrences: 1, Indexes: []int{2}},
				{Word: "THIS", Occurrences: 1, Indexes: []int{3}},
				{Word: "a", Occurrences: 3, Indexes: []int{2}},
				{Word: "at", Occurrences: 1, Indexes: []int{1}},
				{Word: "b", Occurrences: 3, Indexes: []int{2}},
				{Word: "bin", Occurrences: 2, Indexes: []int{3, 4}},
				{Word: "c", Occurrences: 3, Indexes: []int{2}},
				{Word: "clm", Occurrences: 2, Indexes: []int{1}},
				{Word: "cltl", Occurrences: 1, Indexes: []int{1}},
				{Word: "cmu", Occurrences: 1, Indexes: []int{1}},
				{Word: "cs", Occurrences: 1, Indexes: []int{1}},
				{Word: "d", Occurrences: 2, Indexes: []int{2}},
				{Word: "e", Occurrences: 2, Indexes: []int{2}},
				{Word: "edu", Occurrences: 1, Indexes: []int{1}},
				{Word: "env", Occurrences: 2, Indexes: []int{3, 4}},
				{Word: "f", Occurrences: 3, Indexes: []int{2}},
				{Word: "html", Occurrences: 2, Indexes: []int{1}},
				{Word: "http", Occurrences: 1, Indexes: []int{1}},
				{Word: "python", Occurrences: 2, Indexes: []int{3, 4}},
				{Word: "sk", Occurrences: 1, Indexes: []int{2}},
				{Word: "test", Occurrences: 1, Indexes: []int{2}},
				{Word: "the", Occurrences: 1, Indexes: []int{2}},
				{Word: "usr", Occurrences: 2, Indexes: []int{3, 4}},
				{Word: "www", Occurrences: 1, Indexes: []int{1}},
			},
			err: nil,
		},
		{
			words_indexes: []WordIndex{},
			err:           errors.New("lines per page cannot be less than 1"),
		},
		{
			words_indexes: []WordIndex{},
			err:           errors.New("lines per page cannot be less than 1"),
		},
		{
			words_indexes: []WordIndex{
				{Word: "AI", Occurrences: 1, Indexes: []int{1}},
				{Word: "Groups", Occurrences: 1, Indexes: []int{1}},
				{Word: "clm", Occurrences: 2, Indexes: []int{1}},
				{Word: "cltl", Occurrences: 1, Indexes: []int{1}},
				{Word: "cmu", Occurrences: 1, Indexes: []int{1}},
				{Word: "cs", Occurrences: 1, Indexes: []int{1}},
				{Word: "edu", Occurrences: 1, Indexes: []int{1}},
				{Word: "html", Occurrences: 2, Indexes: []int{1}},
				{Word: "http", Occurrences: 1, Indexes: []int{1}},
				{Word: "www", Occurrences: 1, Indexes: []int{1}},
			},
			err: nil,
		},
		{
			words_indexes: []WordIndex{},
			err:           nil,
		},
		{
			words_indexes: []WordIndex{},
			err:           nil,
		},
		{
			words_indexes: []WordIndex{
				{Word: "hello", Occurrences: 1, Indexes: []int{1}},
			},
			err: nil,
		},
		{
			words_indexes: []WordIndex{
				{Word: "hello", Occurrences: 1, Indexes: []int{1}},
			},
			err: nil,
		},
	}

	for i, test := range tests {
		err := test.word_index_manager.ExtractWordsAndIndexes(test.lines, test.lines_per_page)
		if (err == nil && expected_results[i].err != nil) ||
			(err != nil && expected_results[i].err == nil) ||
			(err != nil && expected_results[i].err != nil && err.Error() != expected_results[i].err.Error()) {
			t.Fatalf("Test %d failed:\nexpected err: %v\ngot err: %v", i+1, expected_results[i].err, err)
		}
		test.word_index_manager.SortListAlphabetically()

		if len(test.word_index_manager.words_indexes) != len(expected_results[i].words_indexes) {
			t.Fatalf("Test %d failed:\nexpected length: %d\ngot length: %d\nexpected: %+v\ngot: %+v",
				i+1, len(expected_results[i].words_indexes), len(test.word_index_manager.words_indexes),
				expected_results[i].words_indexes, test.word_index_manager.words_indexes)
		}
		for j, wi := range test.word_index_manager.words_indexes {
			exp := expected_results[i].words_indexes[j]
			if wi.Word != exp.Word || !slices.Equal(wi.Indexes, exp.Indexes) || wi.Occurrences != exp.Occurrences {
				t.Fatalf("Test %d, word %d failed:\nexpected: %+v\ngot: %+v",
					i+1, j+1, exp, wi)
			}
		}
	}
}
