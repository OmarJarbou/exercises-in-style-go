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
				{Word: "A", Indexes: []int{1, 1, 3}},
				{Word: "AI", Indexes: []int{1}},
				{Word: "ApiKey", Indexes: []int{2}},
				{Word: "Available", Indexes: []int{1}},
				{Word: "Chapter", Indexes: []int{2}},
				{Word: "Common", Indexes: []int{2}},
				{Word: "Concatenating", Indexes: []int{2}},
				{Word: "G", Indexes: []int{2}},
				{Word: "Groups", Indexes: []int{1}},
				{Word: "I", Indexes: []int{1}},
				{Word: "IN", Indexes: []int{3}},
				{Word: "LISP", Indexes: []int{1, 2}},
				{Word: "Language", Indexes: []int{2}},
				{Word: "MAC", Indexes: []int{1}},
				{Word: "MIT", Indexes: []int{1}},
				{Word: "Memo", Indexes: []int{1}},
				{Word: "No", Indexes: []int{1}},
				{Word: "PROGRAM", Indexes: []int{3}},
				{Word: "STYLE", Indexes: []int{3}},
				{Word: "Steele", Indexes: []int{2}},
				{Word: "THIS", Indexes: []int{3}},
				{Word: "a", Indexes: []int{2, 2, 2}},
				{Word: "at", Indexes: []int{1}},
				{Word: "b", Indexes: []int{2, 2, 2}},
				{Word: "bin", Indexes: []int{3, 4}},
				{Word: "c", Indexes: []int{2, 2, 2}},
				{Word: "clm", Indexes: []int{1, 1}},
				{Word: "cltl", Indexes: []int{1}},
				{Word: "cmu", Indexes: []int{1}},
				{Word: "cs", Indexes: []int{1}},
				{Word: "d", Indexes: []int{2, 2}},
				{Word: "e", Indexes: []int{2, 2}},
				{Word: "edu", Indexes: []int{1}},
				{Word: "env", Indexes: []int{3, 4}},
				{Word: "f", Indexes: []int{2, 2, 2}},
				{Word: "html", Indexes: []int{1, 1}},
				{Word: "http", Indexes: []int{1}},
				{Word: "python", Indexes: []int{3, 4}},
				{Word: "sk", Indexes: []int{2}},
				{Word: "test", Indexes: []int{2}},
				{Word: "the", Indexes: []int{2}},
				{Word: "usr", Indexes: []int{3, 4}},
				{Word: "www", Indexes: []int{1}},
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
				{Word: "AI", Indexes: []int{1}},
				{Word: "Groups", Indexes: []int{1}},
				{Word: "clm", Indexes: []int{1, 1}},
				{Word: "cltl", Indexes: []int{1}},
				{Word: "cmu", Indexes: []int{1}},
				{Word: "cs", Indexes: []int{1}},
				{Word: "edu", Indexes: []int{1}},
				{Word: "html", Indexes: []int{1, 1}},
				{Word: "http", Indexes: []int{1}},
				{Word: "www", Indexes: []int{1}},
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
				{Word: "hello", Indexes: []int{1}},
			},
			err: nil,
		},
		{
			words_indexes: []WordIndex{
				{Word: "hello", Indexes: []int{1}},
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
			if wi.Word != exp.Word || !slices.Equal(wi.Indexes, exp.Indexes) {
				t.Fatalf("Test %d, word %d failed:\nexpected: %+v\ngot: %+v",
					i+1, j+1, exp, wi)
			}
		}
	}
}
