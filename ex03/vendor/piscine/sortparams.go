package piscine

import (
	"ft"
	"os"
)

func Sort(s []string) []string {
	n := 0
	for range s {
		n++
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}

func SortParams() {
	strs := Sort(os.Args[1:])
	for _, s:= range strs {
		for _, c := range s {
			ft.PrintRune(c)
		}
		ft.PrintRune('\n')
	}
}
