package main

func main() {
	row := 1
	col := 1

	count := 0

	for !(row == 2978 && col == 3083) {
		col += 1
		row -= 1

		if row == 0 {
			row = col
			col = 1
		}

		count++
	}

	prev_code := 20151125
	code := 0

	for i := 0; i < count; i++ {
		code = (prev_code * 252533) % 33554393
		prev_code = code
	}

	println(code)
}
