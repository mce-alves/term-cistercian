package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// The drawings are not dynamic, so changing the dimensions will break the results
const WIDTH = 9
const HEIGHT = 7

type Cistercian struct {
	X     int       // the integer associated with the cistercian representation
	Shape [7]string // the shape of the cistercian that will be drawn
}

func pad_line(s *string, l int) {
	for len([]rune(*s)) < l {
		*s += " "
	}
}

func (c *Cistercian) pad_shape() {
	for i := range c.Shape {
		pad_line(&(c.Shape[i]), WIDTH)
	}
}

// The function is a bit ugly, but at the moment I don't see a way to dynamically generate all the shapes
func create_base_cistercians() [4][10]Cistercian {
	var ret = [4][10]Cistercian{}

	ret[0][0].Shape = [HEIGHT]string{"     ", "    |", "    |", "    |", "    |", "    |", "    |"}
	ret[0][1].Shape = [HEIGHT]string{"     ___", "    |", "    |", "    |", "    |", "    |", "    |"}
	ret[0][2].Shape = [HEIGHT]string{"", "    |", "    |___", "    |", "    |", "    |", "    |"}
	ret[0][3].Shape = [HEIGHT]string{"", "    |\\", "    | \\", "    |", "    |", "    |", "    |"}
	ret[0][4].Shape = [HEIGHT]string{"", "    | /", "    |/", "    |", "    |", "    |", "    |"}
	ret[0][5].Shape = [HEIGHT]string{"     ___", "    |  /", "    | /", "    |", "    |", "    |", "    |"}
	ret[0][6].Shape = [HEIGHT]string{"", "    |   |", "    |   |", "    |", "    |", "    |", "    |"}
	ret[0][7].Shape = [HEIGHT]string{"     ___", "    |   |", "    |   |", "    |", "    |", "    |", "    |"}
	ret[0][8].Shape = [HEIGHT]string{"", "    |   |", "    |___|", "    |", "    |", "    |", "    |"}
	ret[0][9].Shape = [HEIGHT]string{"     ___", "    |   |", "    |___|", "    |", "    |", "    |", "    |"}

	ret[1][0].Shape = [HEIGHT]string{"     ", "    |", "    |", "    |", "    |", "    |", "    |"}
	ret[1][1].Shape = [HEIGHT]string{" ___", "    |", "    |", "    |", "    |", "    |", "    |"}
	ret[1][2].Shape = [HEIGHT]string{"", "    |", " ___|", "    |", "    |", "    |", "    |"}
	ret[1][3].Shape = [HEIGHT]string{"", "   /|", "  / |", "    |", "    |", "    |", "    |"}
	ret[1][4].Shape = [HEIGHT]string{"", "  \\ |", "   \\|", "    |", "    |", "    |", "    |"}
	ret[1][5].Shape = [HEIGHT]string{" ___", " \\  |", "  \\ |", "    |", "    |", "    |", "    |"}
	ret[1][6].Shape = [HEIGHT]string{"", "|   |", "|   |", "    |", "    |", "    |", "    |"}
	ret[1][7].Shape = [HEIGHT]string{" ___", "|   |", "|   |", "    |", "    |", "    |", "    |"}
	ret[1][8].Shape = [HEIGHT]string{"", "|   |", "|___|", "    |", "    |", "    |", "    |"}
	ret[1][9].Shape = [HEIGHT]string{" ___", "|   |", "|___|", "    |", "    |", "    |", "    |"}

	ret[2][0].Shape = [HEIGHT]string{"     ", "    |", "    |", "    |", "    |", "    |", "    |"}
	ret[2][1].Shape = [HEIGHT]string{"", "    |", "    |", "    |", "    |", "    |", "    |___"}
	ret[2][2].Shape = [HEIGHT]string{"", "    |", "    |", "    |", "    |___", "    |", "    |"}
	ret[2][3].Shape = [HEIGHT]string{"", "    |", "    |", "    |", "    |", "    | /", "    |/"}
	ret[2][4].Shape = [HEIGHT]string{"", "    |", "    |", "    |", "    |", "    |\\", "    | \\"}
	ret[2][5].Shape = [HEIGHT]string{"", "    |", "    |", "    |", "    |", "    | \\", "    |__\\"}
	ret[2][6].Shape = [HEIGHT]string{"", "    |", "    |", "    |", "    |", "    |   |", "    |   |"}
	ret[2][7].Shape = [HEIGHT]string{"", "    |", "    |", "    |", "    |", "    |   |", "    |___|"}
	ret[2][8].Shape = [HEIGHT]string{"", "    |", "    |", "    |", "    |___", "    |   |", "    |   |"}
	ret[2][9].Shape = [HEIGHT]string{"", "    |", "    |", "    |", "    |___", "    |   |", "    |___|"}

	ret[3][0].Shape = [HEIGHT]string{"     ", "    |", "    |", "    |", "    |", "    |", "    |"}
	ret[3][1].Shape = [HEIGHT]string{"", "    |", "    |", "    |", "    |", "    |", " ___|"}
	ret[3][2].Shape = [HEIGHT]string{"", "    |", "    |", "    |", " ___|", "    |", "    |"}
	ret[3][3].Shape = [HEIGHT]string{"", "    |", "    |", "    |", "    |", "  \\ |", "   \\|"}
	ret[3][4].Shape = [HEIGHT]string{"", "    |", "    |", "    |", "    |", "   /|", "  / |"}
	ret[3][5].Shape = [HEIGHT]string{"", "    |", "    |", "    |", "    |", "  / |", " /__|"}
	ret[3][6].Shape = [HEIGHT]string{"", "    |", "    |", "    |", "    |", "|   |", "|   |"}
	ret[3][7].Shape = [HEIGHT]string{"", "    |", "    |", "    |", "    |", "|   |", "|___|"}
	ret[3][8].Shape = [HEIGHT]string{"", "    |", "    |", "    |", " ___|", "|   |", "|   |"}
	ret[3][9].Shape = [HEIGHT]string{"", "    |", "    |", "    |", " ___|", "|   |", "|___|"}

	for i := range ret {
		for j := range ret[i] {
			ret[i][j].X = j
			ret[i][j].pad_shape()
		}
	}

	return ret
}

func (c *Cistercian) combine(c2 Cistercian) {
	i := 0

	for i < HEIGHT {
		newline := ""
		c2r := []rune(c2.Shape[i])

		for j, char := range c.Shape[i] {
			if unicode.IsSpace(char) {
				newline += string(c2r[j])
			} else {
				newline += string(char)
			}
		}
		c.Shape[i] = newline
		i++
	}
}

func Convert(number int) Cistercian {
	cistercians := create_base_cistercians()
	numstr := strconv.Itoa(number)
	digits := len([]rune(numstr))

	d, err := strconv.Atoi(string(numstr[0]))

	if err != nil {
		panic(err) // should never happen, because we know it is a number
	}

	result := &cistercians[digits-1][d]

	for i, digit := range numstr {
		if i == 0 {
			continue
		}

		d, err := strconv.Atoi(string(digit))

		if err != nil {
			panic(err) // should never happen, because we know it is a number
		}

		result.combine(cistercians[digits-i-1][d])
	}

	return *result
}

func (c Cistercian) Draw() {
	for _, line := range c.Shape {
		fmt.Println(line)
	}
}
