/**
 * Description：
 * FileName：map_test.go
 * Author：CJiaの青姝
 * Create：2024/12/31 14:02:19
 * Remark：
 */

package slice

import (
	"fmt"
	"strconv"
)

func ExampleFilterMap() {
	src := []int{1, -2, 3}
	dst := FilterMap[int, string](src, func(idx int, src int) (string, bool) {
		return strconv.Itoa(src), src > 0
	})
	fmt.Println(dst)
	// Output: [1 3]
}

func ExampleMap() {
	src := []int{1, 2, 3}
	dst := Map(src, func(idx int, src int) string {
		return strconv.Itoa(src)
	})
	fmt.Println(dst)
	// Output: [1 2 3]
}

func ExampleToMap() {
	elements := []string{"1", "2", "3", "4", "5"}
	resMap := ToMap(elements, func(str string) int {
		num, _ := strconv.Atoi(str)
		return num
	})
	fmt.Println(resMap)
	// Output: map[1:1 2:2 3:3 4:4 5:5]
}

func ExampleToMapV() {
	type eleType struct {
		A string
		B string
		C int
	}
	type eleTypeOut struct {
		A string
		B string
	}
	elements := []eleType{
		{
			A: "a",
			B: "b",
			C: 1,
		},
		{
			A: "c",
			B: "d",
			C: 2,
		},
	}
	resMap := ToMapV(elements, func(ele eleType) (string, eleTypeOut) {
		return ele.A, eleTypeOut{
			A: ele.A,
			B: ele.B,
		}
	})
	fmt.Println(resMap)
	// Output: map[a:{a b} c:{c d}]
}
