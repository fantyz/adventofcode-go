package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadPassports(t *testing.T) {
	testData := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

	assert.Equal(t, 2, len(LoadPassports(LoadPassportBatch(testData))))
}

func TestDay4Pt1(t *testing.T) {
	assert.Equal(t, 216, len(LoadPassportBatch(day4Input)))
}

func TestDay4Pt2(t *testing.T) {
	assert.Equal(t, 150, len(LoadPassports(LoadPassportBatch(day4Input))))
}
