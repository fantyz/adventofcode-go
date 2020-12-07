package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["4"] = Day4 }

/*
--- Day 4: Passport Processing ---
You arrive at the airport only to realize that you grabbed your North Pole Credentials instead of your passport. While these documents are extremely similar, North Pole Credentials aren't issued by a country and therefore aren't actually valid documentation for travel in most of the world.

It seems like you're not the only one having problems, though; a very long line has formed for the automatic passport scanners, and the delay could upset your travel itinerary.

Due to some questionable network security, you realize you might be able to solve both of these problems at the same time.

The automatic passport scanners are slow because they're having trouble detecting which passports have all required fields. The expected fields are as follows:

byr (Birth Year)
iyr (Issue Year)
eyr (Expiration Year)
hgt (Height)
hcl (Hair Color)
ecl (Eye Color)
pid (Passport ID)
cid (Country ID)
Passport data is validated in batch files (your puzzle input). Each passport is represented as a sequence of key:value pairs separated by spaces or newlines. Passports are separated by blank lines.

Here is an example batch file containing four passports:

ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
The first passport is valid - all eight fields are present. The second passport is invalid - it is missing hgt (the Height field).

The third passport is interesting; the only missing field is cid, so it looks like data from North Pole Credentials, not a passport at all! Surely, nobody would mind if you made the system temporarily ignore missing cid fields. Treat this "passport" as valid.

The fourth passport is missing two fields, cid and byr. Missing cid is fine, but missing any other field is not, so this passport is invalid.

According to the above rules, your improved system would report 2 valid passports.

Count the number of valid passports - those that have all required fields. Treat cid as optional. In your batch file, how many passports are valid?

Your puzzle answer was 216.

--- Part Two ---
The line is moving more quickly now, but you overhear airport security talking about how passports with invalid data are getting through. Better add some data validation, quick!

You can continue to ignore the cid field, but each other field has strict rules about what values are valid for automatic validation:

byr (Birth Year) - four digits; at least 1920 and at most 2002.
iyr (Issue Year) - four digits; at least 2010 and at most 2020.
eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
hgt (Height) - a number followed by either cm or in:
If cm, the number must be at least 150 and at most 193.
If in, the number must be at least 59 and at most 76.
hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
pid (Passport ID) - a nine-digit number, including leading zeroes.
cid (Country ID) - ignored, missing or not.
Your job is to count the passports where all required fields are both present and valid according to the above rules. Here are some example values:

byr valid:   2002
byr invalid: 2003

hgt valid:   60in
hgt valid:   190cm
hgt invalid: 190in
hgt invalid: 190

hcl valid:   #123abc
hcl invalid: #123abz
hcl invalid: 123abc

ecl valid:   brn
ecl invalid: wat

pid valid:   000000001
pid invalid: 0123456789
Here are some invalid passports:

eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007
Here are some valid passports:

pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719
Count the number of valid passports - those that have all required fields and valid values. Continue to treat cid as optional. In your batch file, how many passports are valid?

Your puzzle answer was 150.
*/

func Day4() {
	fmt.Println("--- Day 4: Passport Processing ---")
	batch := LoadPassportBatch(day4Input)
	fmt.Println("  Passports containing required fields:", len(batch))
	fmt.Println("    Passports containing required data:", len(LoadPassports(batch)))
}

// NOTE: Normally it would make more sense to combine the data validation into one step. My
// solution here has it split up into multiple steps to allow both part 1 and part 2 to run.

// LoadPassportBatch takes a passport batch and returns a slice of strings where each string
// contains a single line with all fields for a single passport.
func LoadPassportBatch(in string) []map[string]string {
	var batch []map[string]string

	// make sure the  input ends with an empty line to ease reading it
	in += "\n\n"

	passportData := ""
	for _, s := range strings.Split(in, "\n") {
		if s == "" && passportData != "" {
			f, err := LoadPassportFields(passportData)
			if err != nil {
				// ignore invalid passport data
				passportData = ""
				continue
			}
			batch = append(batch, f)
			passportData = ""
			continue
		}
		// append the passport data to a single space separated line
		if passportData != "" {
			passportData += " "
		}
		passportData += s
	}

	return batch
}

// LoadPassportFields takes raw passport data and returns a map containing the individual raw fields.
// LoadPassportFields returns an error if any of the required fields are missing.
func LoadPassportFields(data string) (map[string]string, error) {
	fields := map[string]string{}
	for _, field := range strings.Split(data, " ") {
		if len(field) < 4 {
			return nil, fmt.Errorf("invalid field (field=%s)", field)
		}
		fields[field[0:3]] = field[4:]
	}

	// validate that required fields exist
	if _, ok := fields["byr"]; !ok {
		return nil, errors.New("byr missing")
	}
	if _, ok := fields["iyr"]; !ok {
		return nil, errors.New("iyr missing")
	}
	if _, ok := fields["eyr"]; !ok {
		return nil, errors.New("eyr missing")
	}
	if _, ok := fields["hgt"]; !ok {
		return nil, errors.New("hgt missing")
	}
	if _, ok := fields["hcl"]; !ok {
		return nil, errors.New("hcl missing")
	}
	if _, ok := fields["ecl"]; !ok {
		return nil, errors.New("ecl missing")
	}
	if _, ok := fields["pid"]; !ok {
		return nil, errors.New("pid missing")
	}

	return fields, nil
}

// LoadPassports takes a passport batch validate the data and returns the valid passports it contains.
func LoadPassports(data []map[string]string) []Passport {
	var passports []Passport
	for _, entry := range data {
		p, err := LoadPassport(entry)
		if err != nil {
			// ignore invalid passports
			continue
		}
		passports = append(passports, p)
	}
	return passports
}

// LoadPassport takes a passport batch entry and returns a passport. LoadPassport will return an error
// if the batch entry contains any invalid data.
func LoadPassport(entry map[string]string) (Passport, error) {
	var err error
	p := Passport{}
	for k, v := range entry {
		switch k {
		case "byr":
			p.BirthYear, err = strconv.Atoi(v)
			if err != nil {
				return Passport{}, errors.Wrapf(err, "invalid byr (value=%s)", v)
			}
			if p.BirthYear < 1920 || p.BirthYear > 2002 {
				return Passport{}, errors.Errorf("invalid byr period (byr=%d)", p.BirthYear)
			}
		case "iyr":
			p.IssueYear, err = strconv.Atoi(v)
			if err != nil {
				return Passport{}, errors.Wrapf(err, "invalid iyr (value=%s)", v)
			}
			if p.IssueYear < 2010 || p.IssueYear > 2020 {
				return Passport{}, errors.Errorf("invalid iyr period (iyr=%d)", p.IssueYear)
			}
		case "eyr":
			p.ExpirationYear, err = strconv.Atoi(v)
			if err != nil {
				return Passport{}, errors.Wrapf(err, "invalid eyr (value=%s)", v)
			}
			if p.ExpirationYear < 2020 || p.ExpirationYear > 2030 {
				return Passport{}, errors.Errorf("invalid eyr period (eyr=%d)", p.ExpirationYear)
			}
		case "hgt":
			if len(v) < 3 {
				return Passport{}, errors.Errorf("invalid hgt length (value=%s)", v)
			}
			p.Height, err = strconv.Atoi(v[:len(v)-2])
			if err != nil {
				return Passport{}, errors.Wrapf(err, "invalid hgt (value=%s)", v)
			}
			p.HeightUnit = v[len(v)-2:]
			switch p.HeightUnit {
			case "cm":
				if p.Height < 150 || p.Height > 193 {
					return Passport{}, errors.Errorf("invalid hgt range (height=%d, unit=%s)", p.Height, p.HeightUnit)
				}
			case "in":
				if p.Height < 59 || p.Height > 76 {
					return Passport{}, errors.Errorf("invalid hgt range (height=%d, unit=%s)", p.Height, p.HeightUnit)
				}
			default:
				return Passport{}, errors.Errorf("invalid hgt unit (unit=%s)", p.HeightUnit)
			}
		case "hcl":
			hclExp := regexp.MustCompile(`^#[0-9a-f]{6}$`)
			p.HairColor = v
			if !hclExp.MatchString(p.HairColor) {
				return Passport{}, errors.Errorf("invalid hcl (hcl=%s)", p.HairColor)
			}
		case "ecl":
			eclExp := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
			p.EyeColor = v
			if !eclExp.MatchString(p.EyeColor) {
				return Passport{}, errors.Errorf("invalid ecl (ecl=%s)", p.EyeColor)
			}
		case "pid":
			pidExp := regexp.MustCompile(`^[0-9]{9}$`)
			if !pidExp.MatchString(v) {
				return Passport{}, errors.Errorf("invalid pid (pid=%s)", v)
			}
			p.PassportID, err = strconv.Atoi(v)
			if err != nil {
				panic("should never happen")
			}
		case "cid":
			p.CountryID, err = strconv.Atoi(v)
			if err != nil {
				// invalid cid is ok
				p.CountryID = -1
			}
		}
	}
	return p, nil
}

type Passport struct {
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	Height         int
	HeightUnit     string
	HairColor      string
	EyeColor       string
	PassportID     int
	CountryID      int
}
