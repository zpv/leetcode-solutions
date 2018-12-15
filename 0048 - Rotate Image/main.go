package main

func rotate(matrix [][]int) {
	// reverse rows
	for i := 0; i < len(matrix); i++ {
		reverse(matrix[i])
	}

	// reverse diagonal y-start
	for y := 0; y < len(matrix); y++ {
		for yi := 0; yi < (len(matrix)-y)/2; yi++ {
			a := matrix[y+yi][yi]
			matrix[y+yi][yi] = matrix[len(matrix)-1-yi][len(matrix)-1-yi-y]
			matrix[len(matrix)-1-yi][len(matrix)-1-yi-y] = a
		}
	}

	// reverse diagonal x-start
	for x := 1; x < len(matrix); x++ {
		for xi := 0; xi < (len(matrix)-x)/2; xi++ {
			a := matrix[xi][x+xi]
			matrix[xi][x+xi] = matrix[len(matrix)-1-xi-x][len(matrix)-1-xi]
			matrix[len(matrix)-1-xi-x][len(matrix)-1-xi] = a
		}
	}
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}

func main() {

}
