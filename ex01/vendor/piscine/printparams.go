package piscine

import (
	"ft"
	"os"
)

func RevParams() {
	for _, s:= range os.Args[1:] {
		for _, c := range s {
			ft.PrintRune(c)
		}
		ft.PrintRune('\n')
	}
}
