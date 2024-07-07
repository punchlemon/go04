package piscine

import (
	"ft"
	"os"
)

func PrintProgramName() {
	for _, c := range os.Args[0] {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}
