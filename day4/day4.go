package day4

import (
	"io"
	"regexp"
	"strings"
)

type Field struct {
	name string
	rule *regexp.Regexp
}

func Compute(r io.Reader, runPartTwo bool) (int, error) {
	fields := map[string]*regexp.Regexp{}
	fields["byr"] = regexp.MustCompile(`^[1-2](9|0)([2-9][0-9]|0[1-2])$`)                   // 1920-2002
	fields["iyr"] = regexp.MustCompile(`^20(1[0-9]|20)$`)                                   // 2010-2020
	fields["eyr"] = regexp.MustCompile(`^20(2[0-9]|30)$`)                                   // 2020-2030
	fields["hgt"] = regexp.MustCompile(`^((59|6[0-9]|7[0-6])in)|(1([5-8][0-9]|9[0-3])cm)$`) // 59in-76in OR 150cm-193cm
	fields["ecl"] = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)                   // colors
	fields["hcl"] = regexp.MustCompile(`^#[a-f0-9]{6}$`)                                    // # followed by 6 chars a-f or 0-9
	fields["pid"] = regexp.MustCompile(`^[0-9]{9}$`)                                        // 9 digit

	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		return 0, err
	}
	validPassports := 0
	ps := strings.Split(buf.String(), "\n\n")
	for _, p := range ps {
		fs := strings.Split(strings.ReplaceAll(p, "\n", " "), " ")
		validFieldCount := 0
		for _, f := range fs {
			fieldPair := strings.Split(f, ":")
			if _, ok := fields[fieldPair[0]]; ok {
				if runPartTwo {
					if fields[fieldPair[0]].MatchString(fieldPair[1]) {
						validFieldCount++
					}
				} else {
					validFieldCount++
				}
			}
		}
		if validFieldCount == len(fields) {
			validPassports++
		}
	}
	return validPassports, nil
}
