/**
 * Description：
 * FileName：find_test.go
 * Author：CJiaの青姝
 * Create：2024/12/31 10:56:10
 * Remark：
 */

package slice

import "fmt"

func ExampleFind() {
	val, ok := Find[int]([]int{1, 2, 3}, func(src int) bool {
		return src == 2
	})
	fmt.Println(val, ok)
	val, ok = Find[int]([]int{1, 2, 3}, func(src int) bool {
		return src == 4
	})
	fmt.Println(val, ok)
	// Output:
	// 2 true
	// 0 false
}

func ExampleFindAll() {
	vals := FindAll[int]([]int{2, 3, 4}, func(src int) bool {
		return src%2 == 1
	})
	fmt.Println(vals)
	vals = FindAll[int]([]int{2, 3, 4}, func(src int) bool {
		return src > 5
	})
	fmt.Println(vals)
	// Output:
	// [3]
	// []
}
