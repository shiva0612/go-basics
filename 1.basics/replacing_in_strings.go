package main

import (
	"fmt"
	"os"
	"strings"
)

func replacing_in_strings() {

	a := "shiva.surya.amma.nanna"
	replacer := strings.NewReplacer(".", "-")
	b := replacer.Replace(a)
	fmt.Println(b)
	replacer.WriteString(os.Stdout, a)

	c := strings.Map(func(r rune) rune {
		switch r {
		case '.':
			return '-'
		case 'a':
			return 'A'
		}
		return r
	}, a)
	fmt.Println(c)

}
