package main

import (
	"fmt"
	"github.com/hradecek/programming_in_go/ch01/exercises/01/bigdigits"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type programFlags struct {
	bar bool
}

func main () {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s [-b|--bar] <whole-number>\n\n", filepath.Base(os.Args[0]))
		fmt.Println("-b --bar\tdraw an underbar and an overbar")
		os.Exit(1)
	}

	flags, stringOfDigits := processArgs()

	bigDigits, err := bigdigits.New(stringOfDigits)
	if err != nil {
		log.Fatal(err)
	}

	bigDigitsStr := bigDigits.String()
	if flags.bar {
		bigDigitsStr = addBars(bigDigits.String(), bigDigits.Len())
	}
	fmt.Println(bigDigitsStr)
}

func processArgs() (flags *programFlags, number string) {
	bar := false
	for _, arg := range os.Args {
		switch arg {
		case "-b", "--b":
			bar = true
		default:
			number = arg
		}
	}

	flags = &programFlags{
		bar: bar,
	}

	return flags, number
}

func addBars(number string, width int) string {
	bar := strings.Repeat("*", width)
	return fmt.Sprintf("%s\n%s%s", bar, number, bar)
}
