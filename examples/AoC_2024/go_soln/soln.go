package soln

import (
	"bytes"
	"os"
	"slices"
	"strconv"
)

func Parse(path string) (col1, col2 []int) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fileBytes = bytes.TrimSuffix(fileBytes, []byte{'\n'})

	for line := range bytes.SplitSeq(fileBytes, []byte{'\n'}) {
		splitLine := bytes.Split(line, bytes.Repeat([]byte{' '}, 3))

		num1, _ := strconv.Atoi(string(splitLine[0]))
		num2, _ := strconv.Atoi(string(splitLine[1]))

		col1 = append(col1, num1)
		col2 = append(col2, num2)
	}

	return col1, col2
}

func Abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}

func Part1(col1, col2 []int) int {
	slices.Sort(col1)
	slices.Sort(col2)

	soln := 0
	for i := range len(col1) {
		soln += Abs(col1[i] - col2[i])
	}

	return soln
}

func Part2(col1, col2 []int) int {
	// Could've done this with a map, but let's keep the solution as close as possible to Monkey's solution.
	slices.Sort(col2)

	soln := 0
	for _, num1 := range col1 {
		count := 0
		for _, num2 := range col2 {
			if num1 < num2 {
				break
			}

			if num1 != num2 {
				continue
			}

			count++
		}
		soln += num1 * count
	}

	return soln
}