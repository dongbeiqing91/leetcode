package medium

import (
	"log"
	"testing"
)

func TestZConvert(t *testing.T) {
	res := convert("abcdefg", 3)
	log.Print(res)
}

func convert(s string, numRows int) string {
	str := ""
	arr := getArr(s, numRows)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			str += string(arr[i][j])
		}
	}
	return str
}

func getArr(s string, numRows int) [][]uint8 {
	arr := make([][]uint8, numRows)
	for i := 0; i < len(s); {
		for col := 0; ; col++ {
			if i == len(s) {
				return arr
			}
			if col%(numRows-1) == 0 {
				for row := 0; row < numRows; row++ {
					if len(arr[row]) == 0 {
						arr[row] = make([]uint8, 0)
					}
					arr[row][col] = s[i]
					i++
				}
			} else {
				offset := col % (numRows - 1)
				rowWithVal := numRows - 1 - offset
				for row := 0; row < numRows; row++ {
					if col == rowWithVal {
						arr[row][col] = s[i]
						i++
					} else {
						if i == len(s) {
							return arr
						}
						arr[row][col] = ' '
					}
				}
			}
		}
	}
	return arr
}
