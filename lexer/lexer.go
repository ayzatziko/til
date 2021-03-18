package lexer

import (
	"fmt"
	"strconv"
)

type itemType int

const (
	eofItem itemType = iota
	errItem
	numItem
	opItem

	// test tokens
	charItem
)

type item struct {
	t itemType
	v string
}

type lexer struct {
	input      string
	start, pos int
	items      chan item
}

type stateFn func(*lexer) stateFn

func lex(input string) (*lexer, <-chan item) {
	l := lexer{
		input: input,
		items: make(chan item),
	}

	go l.run()
	return &l, l.items
}

func (l *lexer) run() {
	for state := lexStart; state != nil; {
		state = state(l)
	}

	close(l.items)
}

func lexStart(l *lexer) stateFn {
	_, ok := l.next()
	if !ok {
		l.emit(eofItem)
		return nil
	}

	l.backup()
	return lexNum
}

func (l *lexer) emit(t itemType) {
	s := l.input[l.start:l.pos]
	l.items <- item{t, s}
	l.start = l.pos
}

func lexNum(l *lexer) stateFn {
	for {
		c, ok := l.next()
		if !ok {
			l.emit(numItem)
			break
		}

		if c == "+" || c == "-" {
			l.backup()
			l.emit(numItem)

			return lexOp
		}

		if _, err := strconv.ParseInt(c, 10, 64); err != nil {
			l.errorf(fmt.Sprintf("%s is not int", c))
			return nil
		}
	}

	l.emit(eofItem)

	return nil
}

func (l *lexer) next() (string, bool) {
	if l.pos >= len(l.input) {
		return "", false
	}

	l.pos++
	return string(l.input[l.pos-1]), true
}

func (l *lexer) backup() {
	l.pos--
}

func (l *lexer) errorf(s string) {
	l.items <- item{errItem, s}
}

func lexOp(l *lexer) stateFn {
	l.next()
	l.emit(opItem)

	if _, ok := l.next(); !ok {
		l.errorf("unexpected eof")
	}
	l.backup()
	return lexNum
}
