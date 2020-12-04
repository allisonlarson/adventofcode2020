package day4

import (
	"io"
	"regexp"
	"strings"
)

func Compute(r io.Reader, runPartTwo bool) (int, error) {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		return 0, err
	}

	validPassports := 0
	ps := strings.Split(buf.String(), "\n\n")
	for _, p := range ps {
		passport := newPassport(p)
		if passport.IsValid(runPartTwo) {
			validPassports++
		}
	}
	return validPassports, nil
}

type passport struct {
	data map[string]string
}

func newPassport(line string) *passport {
	data := map[string]string{}
	fs := strings.Fields(line)
	for _, f := range fs {
		fieldPair := strings.Split(f, ":")
		data[fieldPair[0]] = fieldPair[1]
	}
	return &passport{data}
}

func (p *passport) IsValid(strictValidation bool) bool {
	requiredFields := map[string]*regexp.Regexp{}
	requiredFields["byr"] = regexp.MustCompile(`^[1-2](9|0)([2-9][0-9]|0[1-2])$`)                   // 1920-2002
	requiredFields["iyr"] = regexp.MustCompile(`^20(1[0-9]|20)$`)                                   // 2010-2020
	requiredFields["eyr"] = regexp.MustCompile(`^20(2[0-9]|30)$`)                                   // 2020-2030
	requiredFields["hgt"] = regexp.MustCompile(`^((59|6[0-9]|7[0-6])in)|(1([5-8][0-9]|9[0-3])cm)$`) // 59in-76in OR 150cm-193cm
	requiredFields["ecl"] = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)                   // colors
	requiredFields["hcl"] = regexp.MustCompile(`^#[a-f0-9]{6}$`)                                    // # followed by 6 chars a-f or 0-9
	requiredFields["pid"] = regexp.MustCompile(`^[0-9]{9}$`)                                        // 9 digit

	for field, rule := range requiredFields {
		val, ok := p.data[field]
		if !ok {
			return false
		}

		if strictValidation && !rule.MatchString(val) {
			return false
		}
	}
	return true
}
