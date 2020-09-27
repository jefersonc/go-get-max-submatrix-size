package matrix

func main() {

}

func getMaxSubmatrix(matrix [][]int32) int32 {
	if(len(matrix) != len(matrix[0])) {
		panic("é esperada uma matrix quadrada")
	}

	var sizes []int

	for line := range matrix {
		for column := range matrix[line] {
			sizes = append(sizes, getMatrixSizeByPoints(matrix, line, column))
		}
	}

	var major = 0

	for index := range sizes {
		if sizes[index] > major {
			major = sizes[index]
		}
	}

	return int32(major)
}

func getMatrixSizeByPoints(matrix [][]int32, line int, column int) int {

	factor := 0

	for true {
		if !validRight(matrix, line, column, factor) || !validBottom(matrix, line, column, factor) {
			break
		}

		factor++
	}

	return factor
}

func validRight(matrix [][]int32, line int, column int, factor int) bool {
	if  column+factor >= len(matrix[line]) {
		return  false
	}

	if matrix[line][column+factor] == 1 {
		return true
	}

	return false
}

func validBottom(matrix [][]int32, line int, column int, factor int) bool {
	if column+factor >= len(matrix[line]) {
		return  false
	}
	if line+factor >= len(matrix) {
		return  false
	}

	for x := column; x <= column + factor; x++ {
		if matrix[line + factor][x] != 1 {
			return false
		}
	}

	return true
}
