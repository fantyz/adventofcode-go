package main

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["19"] = Day19 }

/*
--- Day 19: Monster Messages ---
You land in an airport surrounded by dense forest. As you walk to your high-speed train, the Elves at the Mythical Information Bureau contact you again. They think their satellite has collected an image of a sea monster! Unfortunately, the connection to the satellite is having problems, and many of the messages sent back from the satellite have been corrupted.

They sent you a list of the rules valid messages should obey and a list of received messages they've collected so far (your puzzle input).

The rules for valid messages (the top part of your puzzle input) are numbered and build upon each other. For example:

0: 1 2
1: "a"
2: 1 3 | 3 1
3: "b"
Some rules, like 3: "b", simply match a single character (in this case, b).

The remaining rules list the sub-rules that must be followed; for example, the rule 0: 1 2 means that to match rule 0, the text being checked must match rule 1, and the text after the part that matched rule 1 must then match rule 2.

Some of the rules have multiple lists of sub-rules separated by a pipe (|). This means that at least one list of sub-rules must match. (The ones that match might be different each time the rule is encountered.) For example, the rule 2: 1 3 | 3 1 means that to match rule 2, the text being checked must match rule 1 followed by rule 3 or it must match rule 3 followed by rule 1.

Fortunately, there are no loops in the rules, so the list of possible matches will be finite. Since rule 1 matches a and rule 3 matches b, rule 2 matches either ab or ba. Therefore, rule 0 matches aab or aba.

Here's a more interesting example:

0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"
Here, because rule 4 matches a and rule 5 matches b, rule 2 matches two letters that are the same (aa or bb), and rule 3 matches two letters that are different (ab or ba).

Since rule 1 matches rules 2 and 3 once each in either order, it must match two pairs of letters, one pair with matching letters and one pair with different letters. This leaves eight possibilities: aaab, aaba, bbab, bbba, abaa, abbb, baaa, or babb.

Rule 0, therefore, matches a (rule 4), then any of the eight options from rule 1, then b (rule 5): aaaabb, aaabab, abbabb, abbbab, aabaab, aabbbb, abaaab, or ababbb.

The received messages (the bottom part of your puzzle input) need to be checked against the rules so you can determine which are valid and which are corrupted. Including the rules and the messages together, this might look like:

0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb
Your goal is to determine the number of messages that completely match rule 0. In the above example, ababbb and abbbab match, but bababa, aaabbb, and aaaabbb do not, producing the answer 2. The whole message must match all of rule 0; there can't be extra unmatched characters in the message. (For example, aaaabbb might appear to match rule 0 above, but it has an extra unmatched b on the end.)

How many messages completely match rule 0?

Your puzzle answer was 115.

--- Part Two ---
As you look over the list of messages, you realize your matching rules aren't quite right. To fix them, completely replace rules 8: 42 and 11: 42 31 with the following:

8: 42 | 42 8
11: 42 31 | 42 11 31
This small change has a big impact: now, the rules do contain loops, and the list of messages they could hypothetically match is infinite. You'll need to determine how these changes affect which messages are valid.

Fortunately, many of the rules are unaffected by this change; it might help to start by looking at which rules always match the same set of values and how those rules (especially rules 42 and 31) are used by the new versions of rules 8 and 11.

(Remember, you only need to handle the rules you have; building a solution that could handle any hypothetical combination of rules would be significantly more difficult.)

For example:

42: 9 14 | 10 1
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
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
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
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba
Without updating rules 8 and 11, these rules only match three messages: bbabbbbaabaabba, ababaaaaaabaaab, and ababaaaaabbbaba.

However, after updating rules 8 and 11, a total of 12 messages match:

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
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba
After updating rules 8 and 11, how many messages completely match rule 0?

Your puzzle answer was 237.
*/

func Day19() {
	fmt.Println("--- Day 19: Monster Messages ---")
	matcher, err := NewMessageRuleMatcher(day19InputRules)
	if err != nil {
		fmt.Println(err)
		return
	}
	matches, err := MessagesMatching(matcher, strings.Split(day19InputMessages, "\n"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Messages matching the message rules:", matches)
	matcher.ReplaceRules()
	matches, err = MessagesMatching(matcher, strings.Split(day19InputMessages, "\n"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Messages matching the message rules (after replace):", matches)
}

// MessagesMatching takes a MessageRuleMatcher and a list of messages and return the number of
// messages that match the MessageRuleMatcher.
func MessagesMatching(matcher *MessageRuleMatcher, messages []string) (int, error) {
	sum := 0
	for _, msg := range messages {
		match, err := matcher.Match(msg)
		if err != nil {
			return 0, errors.Wrapf(err, "failed to match message (message=%s)", msg)
		}
		if match {
			sum++
		}
	}
	return sum, nil
}

// NewMessageRuleMatcher takes a raw set of rules and creates a NewMessageRuleMatcher. It returns
// an error if the rules are invalid.
func NewMessageRuleMatcher(rules string) (*MessageRuleMatcher, error) {
	r := map[string]MessageRule{}

	for _, rule := range strings.Split(rules, "\n") {
		idx := strings.Index(rule, ": ")
		if idx < 0 {
			return nil, errors.Errorf("rule has no id (rule=%s)", rule)
		}
		id := rule[:idx]
		rule = rule[idx+2:]

		// check if rule is a literal
		if len(rule) == 3 && rule[0] == '"' && rule[2] == '"' {
			r[id] = LiteralMessageRule(rule[1])
			continue
		}

		// create either a SequenceMessageRule or an OrMessageRule containing two injected SequenceMessageRules
		var finalRule MessageRule
		seqRule := &SequenceMessageRule{Rules: r}
		finalRule = seqRule
		for _, ruleID := range strings.Split(rule, " ") {
			if ruleID == "|" {
				// create OrMessageRule with two extra SequenceMessageRule (left and right)
				r[id] = &OrMessageRule{
					Rules: r,
					Rule1: id + "l",
					Rule2: id + "r",
				}

				// finish initial left SequenceMessageRule and prepare the following right SequenceMessageRule
				r[id+"l"] = seqRule
				id = id + "r"
				seqRule = &SequenceMessageRule{Rules: r}
				finalRule = seqRule
				continue
			}
			seqRule.Seq = append(seqRule.Seq, ruleID)
		}
		r[id] = finalRule
	}

	return &MessageRuleMatcher{rules: r}, nil
}

type MessageRuleMatcher struct {
	rules    map[string]MessageRule
	replaced bool
}

// Match takes a message and match it. It returns true if the match is exact. An error is returned if
// the MessageRuleMatcher is missing rules (this is not caught by NewMessageRuleMatcher).
func (r *MessageRuleMatcher) Match(message string) (bool, error) {
	// We won't get loops working with this approach easily - the problem is that a given rule can successfully match
	// the same messsage but with different idx (how many characters that it has matched).
	// This can cause it to match, but use "too many" characters causing the rest of the matcher to fail.
	//
	// As per the hints given in the puzzle, we can use the fact that the only place rule 8 and 11 is used is either
	// in themselves or from rule 0. This allow us to manually adjust for the matchers inability to handle loops by
	// splitting the message into two messages repeatedly with all possible lengths (eg. aabb would be a+abb, aa+bb,
	// aab+b) and use rule 8 for the first part and rule 11 for the second part. If both match in any of the
	// combinations, the string matches.

	if !r.replaced {
		// without loops
		rule0, found := r.rules["0"]
		if !found {
			return false, errors.New("no MessageRule with id 0 found")
		}

		i, err := rule0.Match(0, message)
		if err != nil {
			return false, err
		}
		return len(message) == i, nil
	} else {
		// hacked support for loops - depends entirely on rule 0 being "8 11" and the loops being contained
		// within rule 8 and 11 themselves.
		rule8, foundR8 := r.rules["8"]
		rule11, foundR11 := r.rules["11"]
		if !foundR8 || !foundR11 {
			return false, errors.New("no MessageRule with id 8 and/or 11 found")
		}

		for i := 1; i < len(message); i++ {
			part1, part2 := message[:i], message[i:]
			idx, err := rule8.Match(0, part1)
			if err != nil {
				return false, err
			}
			if idx == len(part1) {
				// part1 matched
				idx, err = rule11.Match(0, part2)
				if err != nil {
					return false, err
				}
				if idx == len(part2) {
					// part2 matched as well
					return true, nil
				}
			}
		}
		return false, nil
	}
}

// ReplaceRules will replace rule 8 and 11 with hardcoded rules provided in the puzzle description.
func (r *MessageRuleMatcher) ReplaceRules() {
	r.replaced = true

	// flip the expression from whats given in the puzzle to let it be as greedy as possible to
	// avoid the expression matching a smaller string than it could causing it to complete the
	// match before it has matched all the characters thus not be an actual match.
	r.rules["8"] = &OrMessageRule{
		Rules: r.rules,
		Rule1: "8l",
		Rule2: "42",
	}
	r.rules["8l"] = &SequenceMessageRule{
		Rules: r.rules,
		Seq:   []string{"42", "8"},
	}
	r.rules["11"] = &OrMessageRule{
		Rules: r.rules,
		Rule1: "11l",
		Rule2: "11r",
	}
	r.rules["11l"] = &SequenceMessageRule{
		Rules: r.rules,
		Seq:   []string{"42", "11", "31"},
	}
	r.rules["11r"] = &SequenceMessageRule{
		Rules: r.rules,
		Seq:   []string{"42", "31"},
	}
}

type MessageRule interface {
	Match(int, string) (int, error)
}

// OrMessageRule contains two different MessageRules and match using the either of the two.
type OrMessageRule struct {
	Rules        map[string]MessageRule
	Rule1, Rule2 string
}

func (r *OrMessageRule) Match(idx int, message string) (int, error) {
	if idx == -1 {
		return -1, nil
	}
	r1, found := r.Rules[r.Rule1]
	if !found {
		return -1, errors.Errorf("rule not found (id=%s)", r.Rule1)
	}

	r1Idx, err := r1.Match(idx, message)
	if err != nil {
		return -1, err
	}
	if r1Idx >= 0 {
		// left expression matched
		return r1Idx, nil
	}

	r2, found := r.Rules[r.Rule2]
	if !found {
		return -1, errors.Errorf("rule not found (id=%s)", r.Rule2)
	}
	return r2.Match(idx, message)
}

// SequenceMessageRule is a sequence of message rules that much be matched one by one
// in the order they appear in the sequence.
type SequenceMessageRule struct {
	Rules map[string]MessageRule
	Seq   []string
}

func (r *SequenceMessageRule) Match(idx int, message string) (int, error) {
	var err error
	if idx == -1 {
		return -1, nil
	}

	for _, ruleID := range r.Seq {
		rule, found := r.Rules[ruleID]
		if !found {
			return -1, errors.Errorf("rule not found (id=%s)", ruleID)
		}
		idx, err = rule.Match(idx, message)
		if err != nil {
			return -1, err
		}
		if idx < 0 {
			return -1, nil
		}
	}
	return idx, nil
}

// LiteralMessageRule is a MessageRule that match a single character in the message.
type LiteralMessageRule byte

func (r LiteralMessageRule) Match(idx int, message string) (int, error) {
	if idx < 0 || idx >= len(message) || message[idx] != byte(r) {
		return -1, nil
	}
	return idx + 1, nil
}
