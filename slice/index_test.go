/**
 * Description：
 * FileName：index_test.go
 * Author：CJiaの青姝
 * Create：2024/12/31 11:32:55
 * Remark：
 */

package slice

import "fmt"

func ExampleIndex() {
	res := Index[int]([]int{1, 2, 3}, 1)
	fmt.Println(res)
	res = Index[int]([]int{1, 2, 3}, 4)
	fmt.Println(res)
	// Output:
	// 0
	// -1
}

func ExampleIndexFunc() {
	res := IndexFunc[int]([]int{1, 2, 3}, func(src int) bool {
		return src == 1
	})
	fmt.Println(res)
	res = IndexFunc[int]([]int{1, 2, 3}, func(src int) bool {
		return src == 4
	})
	fmt.Println(res)
	// Output:
	// 0
	// -1
}

func ExampleIndexAll() {
	res := IndexAll[int]([]int{1, 2, 3, 4, 5, 3, 9}, 3)
	fmt.Println(res)
	res = IndexAll[int]([]int{1, 2, 3}, 4)
	fmt.Println(res)
	// Output:
	// [2 5]
	// []
}

func ExampleIndexAllFunc() {
	res := IndexAllFunc[int]([]int{1, 2, 3, 4, 5, 3, 9}, func(src int) bool {
		return src == 3
	})
	fmt.Println(res)
	res = IndexAllFunc[int]([]int{1, 2, 3}, func(src int) bool {
		return src == 4
	})
	fmt.Println(res)
	// Output:
	// [2 5]
	// []
}
