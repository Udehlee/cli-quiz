package main

import (
	"testing"
)

func TestSort(t *testing.T) {
	var tests = []struct {
		records  [][]string
		expected []Quiz
	}{
		{
			records: [][]string{
				{"Question", "Answer"},
			},
			expected: []Quiz{
				{question: "Question", answer: "Answer"},
			},
		},
	}

	for _, tc := range tests {
		t.Run("TestSort", func(t *testing.T) {
			got, err := sort(tc.records)

			// Check for error
			if err != nil {
				t.Errorf("sort() error = %v", err)
				return
			}

			// Check the length of the result
			if len(got) != len(tc.expected) {
				t.Errorf("got= %d, want %d", len(got), len(tc.expected))
				return
			}

			// Check each Quiz struct in the result
			for i := 0; i < len(got); i++ {
				if got[i].question != tc.expected[i].question || got[i].answer != tc.expected[i].answer {
					t.Errorf("sort() result[%d]: got %v, want %v", i, got[i], tc.expected[i])
				}
			}
		})
	}
}
