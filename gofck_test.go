package main

import (
	"testing"

	assert "gopkg.in/go-playground/assert.v1"
)

func TestHelloWorldNoNested(t *testing.T) {
	c := []byte("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.")
	r := interpret(c)

	assert.Equal(t, []byte("Hello World!\n"), r)
}

func TestHelloNested(t *testing.T) {
	c := []byte("+[-[<<[+[--->]-[<<<]]]>>>-]>-.---.>..>.<<<<-.<+.>>>>>.>.<<.<-.")
	r := interpret(c)

	assert.Equal(t, []byte("hello world"), r)
}

func TestEmail(t *testing.T) {
	c := []byte("+[----->+++<]>+.---.+++++++..+++.+[--->++++<]>.----[->++<]>-.--------.+++.------.--------.[->+++<]>++.[--->++<]>-.++++[->+++<]>.")
	r := interpret(c)

	assert.Equal(t, []byte("hello@world.se"), r)
}
