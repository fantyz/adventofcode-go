#!/bin/env bash

set -e

if [ "x$1" == "x" ] ||  [ "x$2" == "x" ]; then
    echo "Usage:   ./generate-empty-files <path> <name>"
    echo "Example: ./generate-empty-files 2020/ 25 would generate 2020/25.go, 2020/25_test.go and 2020/25_input.go"
    exit 0
fi

DAY="${1%/}/${2}.go"
TEST="${1%/}/${2}_test.go"
INPUT="${1%/}/${2}_input.go"

if [ -f "$DAY" ] || [ -f "$TEST" ] || [ -f "$INPUT" ]; then
    echo "${DAY}, ${TEST} and/or ${INPUT} already exist!"
    exit 1
fi

# <day>.go
echo "package main

import (
       \"fmt\"
)

func init() { days[\"${2}\"] = Day${2} }

/*
 */

func Day${2}() {
	fmt.Println(\"\")
}
" > "${DAY}"

# <day>_test.go
echo "package main

import (
	\"testing\"


        \"github.com/stretchr/testify/assert\"
)

func TestDay${2}Pt1(t *testing.T) {
}

func TestDay${2}Pt2(t *testing.T) {
}
" > "${TEST}"

# <day>_input.go
echo "package main

const day${2}Input = \`\`
" > "${INPUT}"
