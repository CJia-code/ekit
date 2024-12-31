/**
 * Description：
 * FileName：aggregate_test.go
 * Author：CJiaの青姝
 * Create：2024/12/30 15:05:39
 * Remark：
 */

package slice

import "fmt"

func ExampleSum() {
	res := Sum[int]([]int{1, 2, 3})
	fmt.Println(res)
	res = Sum[int](nil)
	fmt.Println(res)
	// Output:
	// 6
	// 0
}

func ExampleMin() {
	res := Min[int]([]int{1, 2, 3})
	fmt.Println(res)
	// Output:
	// 1
}

func ExampleMax() {
	res := Max[int]([]int{1, 2, 3})
	fmt.Println(res)
	// Output:
	// 3
}
