package main

import (
	"strings"
)

var MORSE_CODE = map[string]string{
	".-":        "A",
	"-...":      "B",
	"-.-.":      "C",
	"-..":       "D",
	".":         "E",
	"..-.":      "F",
	"--.":       "G",
	"....":      "H",
	"..":        "I",
	".---":      "J",
	"-.-":       "K",
	".-..":      "L",
	"--":        "M",
	"-.":        "N",
	"---":       "O",
	".--.":      "P",
	"--.-":      "Q",
	".-.":       "R",
	"...":       "S",
	"-":         "T",
	"..-":       "U",
	"...-":      "V",
	".--":       "W",
	"-..-":      "X",
	"-.--":      "Y",
	"--..":      "Z",
	".-.-":      "Ä",
	".--.-":     "Á",
	"..-..":     "É",
	"--.--":     "Ñ",
	"---.":      "Ö",
	"..--":      "Ü",
	"-----":     "0",
	".----":     "1",
	"..---":     "2",
	"...--":     "3",
	"....-":     "4",
	".....":     "5",
	"-....":     "6",
	"--...":     "7",
	"---..":     "8",
	"----":      "9",
	".-.-.-":    ".",
	"--..--":    ",",
	"---...":    ":",
	"..--..":    "?",
	".----.":    "'",
	"-....-":    "-",
	"-..-.":     "/",
	"-.--.-":    "(",
	".-..-.":    "\"",
	".--.-.":    "@",
	"-...-":     "=",
	"........":  "Error",
	"...−−−...": "SOS",
}

func DecodeMorse(morseCode string) string {
	undecoded := strings.Split(strings.ReplaceAll(morseCode, "  ", " | "), " ")

	var decoded = []string{}

	for _, item := range undecoded {
		if item == "|" {
			decoded = append(decoded, " ")
		} else {
			decoded = append(decoded, MORSE_CODE[item])
		}
	}

	return strings.TrimSpace(strings.Join(decoded, ""))
}

// Reference: https://www.codewars.com/kata/54b724efac3d5402db00065e
