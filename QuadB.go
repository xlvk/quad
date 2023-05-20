package piscine

import (
	"github.com/01-edu/z01"
)

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	i := 0
	if i < x {
		for c := 0; c < 1; c++ {
			z01.PrintRune('/')
		}
		for c := 1; c < x-1; c++ {
			z01.PrintRune('*')
		}
		for c := x; c <= x; c++ {
			if x > 1 {
				z01.PrintRune('\\')
			} else {
				break
			}
		}
	}
	z01.PrintRune('\n')
	for i := 0; i < y-2; i++ {
		z01.PrintRune('*')
		if x == 1 {
			z01.PrintRune('\n')
		}
		if y >= 2 {
			for n := 0; n < x-2; n++ {
				z01.PrintRune(' ')
			}
			if x >= 2 {
				z01.PrintRune('*')
				z01.PrintRune('\n')
			}
		}
	}
	if y >= 2 {
		for c := 0; c < 1; c++ {
			z01.PrintRune('\\')
		}
		for c := 1; c < x-1; c++ {
			z01.PrintRune('*')
		}
		for c := x; c <= x; c++ {
			if x > 1 {
				z01.PrintRune('/')
			} else {
				break
			}
		}
	}
}
