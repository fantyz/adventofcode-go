package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageRuleMatcherEx1(t *testing.T) {
	testRules := `0: 1 2
1: "a"
2: 1 3 | 3 1
3: "b"`
	r, err := NewMessageRuleMatcher(testRules)
	if err != nil {
		assert.NoError(t, err, "testRules")
		t.FailNow()
	}

	testCases := []struct {
		Message string
		Match   bool
	}{
		{"aab", true},
		{"aba", true},
		{"abz", false},
	}

	for _, c := range testCases {
		match, err := r.Match(c.Message)
		if assert.NoError(t, err, c.Message) {
			assert.Equal(t, c.Match, match, c.Message)
		}
	}
}

func TestMessageRuleMatcherEx2(t *testing.T) {
	testRules := `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"`
	matcher, err := NewMessageRuleMatcher(testRules)
	if err != nil {
		assert.NoError(t, err, "testRules")
		t.FailNow()
	}

	testCases := []struct {
		Message string
		Match   bool
	}{
		{"ababbb", true},
		{"bababa", false},
		{"abbbab", true},
		{"aaabbb", false},
		{"aaaabbb", false},
	}

	for _, c := range testCases {
		match, err := matcher.Match(c.Message)
		if assert.NoError(t, err, c.Message) {
			assert.Equal(t, c.Match, match, c.Message)
		}
	}
}

func TestMessageRuleMatcherWithReplace(t *testing.T) {
	testRules := `42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1`

	testMessages := `abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba`

	matcher, err := NewMessageRuleMatcher(testRules)
	if err != nil {
		assert.NoError(t, err, "testRules")
		t.FailNow()
	}
	matches, err := MessagesMatching(matcher, strings.Split(testMessages, "\n"))
	if assert.NoError(t, err, "MessagesMatching wo replace") {
		assert.Equal(t, 3, matches, "MessagesMatching wo replace")
	} else {
		t.FailNow()
	}

	matcher.ReplaceRules()
	matches, err = MessagesMatching(matcher, strings.Split(testMessages, "\n"))
	if assert.NoError(t, err, "MessagesMatching w replace") {
		assert.Equal(t, 12, matches, "MessagesMatching w replace")
	}
}

func TestDay19Pt1(t *testing.T) {
	matcher, err := NewMessageRuleMatcher(day19InputRules)
	if assert.NoError(t, err, "NewMessageRuleMacher") {
		matches, err := MessagesMatching(matcher, strings.Split(day19InputMessages, "\n"))
		if assert.NoError(t, err, "MessagesMatching") {
			assert.Equal(t, 115, matches)
		}

	}
}

func TestDay19Pt2(t *testing.T) {
	matcher, err := NewMessageRuleMatcher(day19InputRules)
	if assert.NoError(t, err, "NewMessageRuleMacher") {
		matcher.ReplaceRules()
		matches, err := MessagesMatching(matcher, strings.Split(day19InputMessages, "\n"))
		if assert.NoError(t, err, "MessagesMatching") {
			assert.Equal(t, 237, matches)
		}
	}
}
