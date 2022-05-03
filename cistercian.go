package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

// The drawings are not dynamic, so changing the dimensions will break the results
const WIDTH = 9
const HEIGHT = 7

// The struct representing a cistercian number
// Note that these numbers range from 0 to 9999 ( [0, 9999] )
type Cistercian struct {
	X     int       // the integer associated with the cistercian representation
	Shape [7]string // the shape of the cistercian that will be drawn
}

// Adds padding to ensure a line as length = WIDTH
func pad_line(s *string, l int) {
	for len([]rune(*s)) < l {
		*s += " "
	}
}

// Adds padding to all lines in a shape to ensure all rows have the same width
func (c *Cistercian) pad_shape() {
	for i := range c.Shape {
		pad_line(&(c.Shape[i]), WIDTH)
	}
}

// Returns the zero cistercian
func zero() Cistercian {
	ret := Cistercian{}

	ret.X = 0
	ret.Shape = [HEIGHT]string{"     ", "    |", "    |", "    |", "    |", "    |", "    |"}

	ret.pad_shape()

	return ret
}

// Returns cistercians from 1-9
// The remaining cistercians can be obtained by flipping and/or combining these
func create_base_cistercians() [10]Cistercian {
	var ret = [10]Cistercian{}

	ret[0].Shape = [HEIGHT]string{"     ", "    |", "    |", "    |", "    |", "    |", "    |"}
	ret[1].Shape = [HEIGHT]string{"     ___", "    |", "    |", "    |", "    |", "    |", "    |"}
	ret[2].Shape = [HEIGHT]string{"", "    |", "    |___", "    |", "    |", "    |", "    |"}
	ret[3].Shape = [HEIGHT]string{"", "    |\\", "    | \\", "    |", "    |", "    |", "    |"}
	ret[4].Shape = [HEIGHT]string{"", "    | /", "    |/", "    |", "    |", "    |", "    |"}
	ret[5].Shape = [HEIGHT]string{"     ___", "    |  /", "    | /", "    |", "    |", "    |", "    |"}
	ret[6].Shape = [HEIGHT]string{"", "    |   |", "    |   |", "    |", "    |", "    |", "    |"}
	ret[7].Shape = [HEIGHT]string{"     ___", "    |   |", "    |   |", "    |", "    |", "    |", "    |"}
	ret[8].Shape = [HEIGHT]string{"", "    |   |", "    |___|", "    |", "    |", "    |", "    |"}
	ret[9].Shape = [HEIGHT]string{"     ___", "    |   |", "    |___|", "    |", "    |", "    |", "    |"}

	for i := range ret {
		ret[i].X = i
		ret[i].pad_shape()
	}

	return ret
}

//Takes a cistercian number and flips it horizontally (which translates to multiplying by 10)
func flip_horizontal(c Cistercian) Cistercian {
	rc := Cistercian{}

	rc.X = c.X * 10

	for i, line := range c.Shape {
		rune_line := []rune(line)
		new_line := ""
		j := len(rune_line) - 1

		for j >= 0 {
			if rune_line[j] == []rune("/")[0] {
				new_line += "\\" // invert / to \
			} else if rune_line[j] == []rune("\\")[0] {
				new_line += "/" // invert \ to /
			} else {
				new_line += string(rune_line[j]) // just copy symbol
			}

			j--
		}

		rc.Shape[i] = new_line
	}

	return rc
}

// Takes a cistercian number and flips it vertically (which translates to multiplying by 100)
func flip_vertical(c Cistercian) Cistercian {
	rc := Cistercian{}

	rc.X = c.X * 100

	i := len(c.Shape) - 1
	x := 0
	rc.Shape[x] = ""
	x++

	// a bit more complex that horizontal rotation, because <underscores> must be handled differently
	// otherwise the shapes get deformed
	for i >= 1 {
		prev_rune_line := []rune(c.Shape[i-1])
		rune_line := []rune(c.Shape[i])
		new_line := ""
		j := 0

		for j < len(rune_line) {
			if prev_rune_line[j] == []rune("_")[0] {
				new_line += "_" // _ is moved one line
			} else if rune_line[j] == []rune("/")[0] {
				new_line += "\\" // invert / to \
			} else if rune_line[j] == []rune("\\")[0] {
				new_line += "/" // invert \ to /
			} else if rune_line[j] == []rune("_")[0] {
				new_line += " " // _ is moved one line
			} else {
				new_line += string(rune_line[j]) // just copy symbol
			}

			j++
		}

		rc.Shape[x] = new_line
		i--
		x++
	}

	rc.pad_shape()
	return rc
}

// Combines two cistercian numbers
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

// Main function to convert an integer between [0, 9999] to its Cistercian representation
func Convert(number int) Cistercian {
	if number < 0 || number > 9999 {
		panic(errors.New("Function <Convert> received a number outside of valid Cistercian range."))
	}

	cistercians := create_base_cistercians()
	numstr := strconv.Itoa(number)
	digits := len([]rune(numstr))

	zero := zero()

	result := &zero

	for i, digit := range numstr {
		d, err := strconv.Atoi(string(digit))
		if err != nil {
			panic(err) // should never happen, because we know it is a number
		}

		var cist Cistercian

		switch digits - i {
		case 4: // thousands
			cist = flip_horizontal(flip_vertical(cistercians[d]))
		case 3: // hundreds
			cist = flip_vertical(cistercians[d])
		case 2: // tens
			cist = flip_horizontal(cistercians[d])
		default: // units
			cist = cistercians[d]
		}

		result.combine(cist)
	}

	result.X = number
	return *result
}

func (c Cistercian) Draw() {
	for _, line := range c.Shape {
		fmt.Println(line)
	}
}
