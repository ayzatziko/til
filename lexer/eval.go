package lexer

import "strconv"

func eval(s string) int {
	v := 0
	f := func(i int) {
		v = v + i
	}

	_, items := lex(s)

	for it := range items {
		switch it.t {
		case numItem:
			n, _ := strconv.ParseInt(it.v, 10, 64)
			f(int(n))
		case opItem:
			switch it.v {
			case "+":
				f = func(i int) {
					v = v + i
				}
			case "-":
				f = func(i int) {
					v = v - i
				}
			}
		}
	}

	return v
}
