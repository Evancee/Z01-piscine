package piscine
import github.com/01-edu/z01


func IsNegative(nb int) {
	if nb < 0 {
		z01.printrune('T')
	} else {
		z01.printrune('F')
	}
	z01.printrune('\n')

}
