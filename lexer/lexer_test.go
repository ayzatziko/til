package lexer

import (
	"fmt"
	"testing"
)

// func TestLexer(t *testing.T) {
// 	tests := []struct {
// 		input string
// 		items []item
// 	}{}

// 	for _, test := range tests {
// 		t.Run(test.input, func(t *testing.T) {
// 			_, items := lex(test.input)

// 			skipFail := ""
// 			for it := range items {
// 				if skipFail != "" {
// 					continue
// 				}

// 				if len(test.items) == 0 {
// 					skipFail = "got more items than expected"
// 					continue
// 				}

// 				if test.items[0] != it {
// 					skipFail = fmt.Sprintf("want item %v, got %v", test.items[0], it)
// 					continue
// 				}

// 				test.items = test.items[1:]
// 			}
// 		})
// 	}
// }

func TestLexerV2(t *testing.T) {
	tests := []string{
		// "",
		// "4",
		// "123",
		"3+2",
		"123+321",
		// "+",
		// "-",
		// "2+",
		// "2-",
		// "1+1",
	}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			_, items := lex(tt)

			for it := range items {
				fmt.Printf("%+v", it)
			}
			fmt.Println()
		})
	}
}
