package piscine

import (
	"ft"
	"os"
)

func Invert(s []string) []string {
	len := 0
	for range s {
		len++
	}
	for i := 0; i < len/2; i++ {
		s[i], s[len-1-i] = s[len-1-i], s[i]
	}
	return s
}

func RevParams() {
	strs := Invert(os.Args[1:])
	for _, s:= range strs {
		for _, c := range s {
			ft.PrintRune(c)
		}
		ft.PrintRune('\n')
	}
}
