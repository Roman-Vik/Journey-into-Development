package main

import "fmt"

func GetGrade(a, b, c int) rune {
	score := ((a + b + c) / 3)
	var res rune
	switch {
	case 90 <= score && score <= 100:
		res = 'A'
	case 80 <= score && score < 90:
		res = 'B'
	case 70 <= score && score < 80:
		res = 'C'
	case 60 <= score && score < 70:
		res = 'D'
	case 0 <= score && score < 60:
		res = 'F'
	}
	return res
}

func main(){
	fmt.Println(GetGrade(100,49,54))
}

/*
func GetGrade(a,b,c int) rune {
    return rune("FFFFFFDCBAA"[(a+b+c)/30])
}

package kata

const (
	gradeRunes = "FFFFFFDCBAA"
	divisor    = 30
)

// GetGrade calculates the average of three scores and returns the corresponding grade as a rune.
// The grades are based on the index within the gradeRunes string.
func GetGrade(a, b, c int) rune {
	index := (a + b + c) / divisor
	return rune(gradeRunes[index])
}

*/ 