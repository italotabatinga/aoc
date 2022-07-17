package aoc

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Input4 []map[string]string

type Runner4 struct{}

func (r Runner4) FmtInput(input string) Input4 {
	lines := strings.Split(input, "\n\n")
	var result Input4
	count := len(lines)
	result = make(Input4, count)
	for i, s := range lines {
		fields := strings.Fields(s)
		passport := make(map[string]string)
		for _, field := range fields {
			if split := strings.Split(field, ":"); len(split) == 2 {
				passport[split[0]] = split[1]
			} else {
				panic(fmt.Errorf("unexpected field: %v", field))
			}
		}
		result[i] = passport
	}

	return result
}

func (r Runner4) Run1(input Input4) int {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	valids := 0
	for _, passport := range input {
		isValid := true
		for _, requiredField := range requiredFields {
			if _, ok := passport[requiredField]; !ok {
				isValid = false
				break
			}
		}
		if isValid {
			valids++
		}
	}
	fmt.Printf("Valids: %v\n", valids)
	return valids
}

func (r Runner4) Run2(input Input4) int {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	valids := 0
	for _, passport := range input {
		isValid := true
		for _, requiredField := range requiredFields {
			if value, ok := passport[requiredField]; ok {
				if !checkPassportField(requiredField, value) {
					// fmt.Printf("passport: %v\n", passport)
					// fmt.Printf("invalid rule! f: %v, v: %v\n\n", requiredField, value)
					isValid = false
					break
				}
			} else {
				// fmt.Printf("passport: %v\n", passport)
				// fmt.Printf("missing field! f: %v\n\n", requiredField)
				isValid = false
				break
			}
		}
		if isValid {
			valids++
		}
	}
	fmt.Printf("Valids: %v\n", valids)
	return valids
}

func checkPassportField(field string, value string) bool {
	switch field {
	case "byr":
		if year, err := strconv.Atoi(value); err == nil {
			return year >= 1920 && year <= 2002
		}
	case "iyr":
		if year, err := strconv.Atoi(value); err == nil {
			return year >= 2010 && year <= 2020
		}
	case "eyr":
		if year, err := strconv.Atoi(value); err == nil {
			return year >= 2020 && year <= 2030
		}
	case "hgt":
		unit := value[len(value)-2:]
		val := value[0 : len(value)-2]
		if year, err := strconv.Atoi(val); err == nil {
			switch unit {
			case "cm":
				return year >= 150 && year <= 193
			case "in":
				return year >= 59 && year <= 76
			}
		}
	case "hcl":
		reg := regexp.MustCompile(`^#[0-9a-f]{6}$`)
		return reg.MatchString(value)
	case "ecl":
		switch value {
		case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			return true
		}
	case "pid":
		reg := regexp.MustCompile(`^[0-9]{9}$`)
		return reg.MatchString(value)
	default:
		panic(fmt.Errorf("unexpected field: %v", field))
	}
	return false
}
