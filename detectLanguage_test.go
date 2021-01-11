package main

import (
	"io/ioutil"
	"testing"
)

type testCase struct {
	language string
	code     string
	score    float64
}

var testCases []testCase

func init() {
	testCases = []testCase{
		{"c", "testFiles/c.txt", 0.4},
		{"csharp", "testFiles/csharp.txt", 0.6},
		{"go", "testFiles/go.txt", 0.5},
	}
}

func TestCheckAndDetection(t *testing.T) {
	for _, test := range testCases {
		content, err := ioutil.ReadFile(test.code)
		if err != nil {
			t.Errorf("Error reading file %s: %s", test.code, err)
		}
		code := string(content)

		result := check(code, checks[test.language])
		if result < test.score {
			t.Errorf("Certainty should be at least %.1f, but is %.1f", test.score, result)
		}

		language := detectLanguage(code)
		if language != test.language {
			t.Errorf("Language should be %s, but is %s", test.language, language)
		}
	}
}
