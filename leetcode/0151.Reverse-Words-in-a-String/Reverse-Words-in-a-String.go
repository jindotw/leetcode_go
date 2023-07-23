package main

import (
	"fmt"
	"strings"
	"unicode"
)

func reverseWords(s string) string {
	trimSpace := func() string {
		bytes := []byte(s)
		lft, rgt := 0, len(bytes)-1
		for bytes[lft] == ' ' {
			lft++
		}
		for bytes[rgt] == ' ' {
			rgt--
		}
		var prev byte
		builder := strings.Builder{}
		for i := lft; i <= rgt; i++ {
			ch := bytes[i]
			if ch != ' ' || prev != ' ' {
				prev = ch
				builder.WriteByte(ch)
			}
		}

		return builder.String()
	}

	reverse := func(str string) string {
		bytes := []byte(str)
		size := len(bytes)
		for lft := 0; lft < size/2; lft++ {
			rgt := size - 1 - lft
			bytes[lft], bytes[rgt] = bytes[rgt], bytes[lft]
		}

		return string(bytes)
	}

	trimmed := trimSpace()
	reversed := reverse(trimmed)
	builder := strings.Builder{}
	begin := 0
	for idx, v := range reversed {
		if v == ' ' {
			builder.WriteString(reverse(reversed[begin:idx]))
			builder.WriteRune(v)
			begin = idx + 1
		}
	}
	builder.WriteString(reverse(reversed[begin:]))

	return builder.String()
}

func reverseWords2(s string) string {
	words := strings.FieldsFunc(strings.TrimSpace(s), func(r rune) bool {
		return unicode.IsSpace(r)
	})
	size := len(words)
	for lft := 0; lft < size/2; lft++ {
		rgt := size - 1 - lft
		words[lft], words[rgt] = words[rgt], words[lft]
	}

	return strings.Join(words, " ")
}

func main() {
	fmt.Printf("`%s`\n", reverseWords2("  the   sky   is   blue  "))
}
