package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

type pair struct {
	R rune
	M int
}

const (
	DIGIT = iota
	CHAR
	ESCAPE
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var (
		prev    = DIGIT
		buffer  []pair
		sb      strings.Builder
		p       pair
		isEsc   bool
		isDigit bool
	)

	if len(s) == 0 {
		return "", nil
	}

	buffer = make([]pair, 0)
	for _, r := range s {
		isEsc = r == '\\'
		isDigit = unicode.IsDigit(r)

		if prev == ESCAPE {
			if !isEsc && !isDigit {
				return "", ErrInvalidString
			}
			p = pair{R: r, M: 1}
			prev = CHAR
			continue
		}

		if isDigit {
			if prev == DIGIT {
				return "", ErrInvalidString
			}
			p.M = int(r - '0')
			prev = DIGIT
			continue
		}

		buffer = append(buffer, p)
		if isEsc {
			prev = ESCAPE
			continue
		}

		p = pair{R: r, M: 1}
		prev = CHAR
	}
	buffer = append(buffer, p)

	if prev == ESCAPE {
		return "", ErrInvalidString
	}

	for _, p = range buffer {
		sb.WriteString(strings.Repeat(string(p.R), p.M))
	}

	return sb.String(), nil
}
