package bigdigits

import (
	"errors"
	"fmt"
	"strings"
)

var bigDigits = [][]string{
	{"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  "},
	{" 1 ",
		"11 ",
		" 1 ",
		" 1 ",
		" 1 ",
		" 1 ",
		"111"},
	{" 222 ",
		"2   2",
		"   2 ",
		"  2  ",
		" 2   ",
		"2    ",
		"22222"},
	{" 333 ",
		"3   3",
		"    3",
		"  33 ",
		"    3",
		"3   3",
		" 333 "},
	{"   4  ",
		"  44  ",
		" 4 4  ",
		"4  4  ",
		"444444",
		"   4  ",
		"   4  "},
	{"55555",
		"5    ",
		"5    ",
		" 555 ",
		"    5",
		"5   5",
		" 555 "},
	{" 666 ",
		"6    ",
		"6    ",
		"6666 ",
		"6   6",
		"6   6",
		" 666 "},
	{"777777",
		"     7",
		"    7 ",
		"   7  ",
		"  7   ",
		" 7    ",
		"7     "},
	{" 888 ",
		"8   8",
		"8   8",
		" 888 ",
		"8   8",
		"8   8",
		" 888 "},
	{" 9999",
		"9   9",
		"9   9",
		" 9999",
		"    9",
		"    9",
		"    9"}}

type BigDigits struct {
	digits string
}

func New(stringOfDigits string) (*BigDigits, error) {
	line := ""
	for row := range bigDigits[0] {
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if digit <= 9 {
				line += bigDigits[digit][row] + " "
			} else {
				return nil, errors.New(fmt.Sprintf("invalid whole number: %s", stringOfDigits))
			}
		}
		line += "\n"
	}

	return &BigDigits{line}, nil
}

func (b *BigDigits) Len() int {
	return len(strings.Split(b.digits, "\n")[0]) - 1
}

func (b *BigDigits) String() string {
	return b.digits
}
