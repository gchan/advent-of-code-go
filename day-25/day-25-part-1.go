package main

func main() {
	row := 1
	col := 1

	count := 0

	for !(row == 2978 && col == 3083) {
		col++
		row--

		if row == 0 {
			row = col
			col = 1
		}

		count++
	}

	prevCode := 20151125
	code := 0

	for i := 0; i < count; i++ {
		code = (prevCode * 252533) % 33554393
		prevCode = code
	}

	println(code)
}
