/**
 * Description：
 * FileName：contains_test.go
 * Author：CJiaの青姝
 * Create：2024/12/30 15:06:01
 * Remark：
 */

package slice

import "fmt"

func ExampleContains() {
	res := Contains[int]([]int{1, 2, 3}, 3)
	fmt.Println(res)
	// Output:
	// true
}

func ExampleContainsFunc() {
	res := ContainsFunc[int]([]int{1, 2, 3}, func(src int) bool {
		return src == 3
	})
	fmt.Println(res)
	// Output:
	// true
}

func ExampleContainsAny() {
	res := ContainsAny[int]([]int{1, 2, 3}, []int{3, 6})
	fmt.Println(res)
	res = ContainsAny[int]([]int{1, 2, 3}, []int{4, 5, 9})
	fmt.Println(res)
	// Output:
	// true
	// false
}

func ExampleContainsAnyFunc() {
	res := ContainsAnyFunc[int]([]int{1, 2, 3}, []int{3, 1}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	res = ContainsAllFunc[int]([]int{1, 2, 3}, []int{4, 7, 6}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	// Output:
	// true
	// false
}

func ExampleContainsAll() {
	res := ContainsAll[int]([]int{1, 2, 3}, []int{3, 1})
	fmt.Println(res)
	res = ContainsAll[int]([]int{1, 2, 3}, []int{3, 1, 4})
	fmt.Println(res)
	// Output:
	// true
	// false
}

func ExampleContainsAllFunc() {
	res := ContainsAllFunc[int]([]int{1, 2, 3}, []int{3, 1}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	res = ContainsAllFunc[int]([]int{1, 2, 3}, []int{3, 1, 4}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	// Output:
	// true
	// false
}
