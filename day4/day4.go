package day4

import (
	"io"
	"strings"
)

var validFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func Compute(r io.Reader, runPartTwo bool) (int, error) {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		return 0, err
	}
	validPassports := 0
	ps := strings.Split(buf.String(), "\n\n")
	for _, p := range ps {
		p = strings.ReplaceAll(p, "\n", " ")
		p = strings.TrimSpace(p)
		fs := strings.Split(p, " ")
		if len(fs) == 8 {
			validPassports++
			continue
		}
		validFieldCount := 0
		for _, f := range fs {
			for _, validField := range validFields {
				if strings.Contains(f, validField) {
					validFieldCount++
				}
			}
		}
		if validFieldCount >= 7 {
			validPassports++
		}
	}
	return validPassports, nil
}
