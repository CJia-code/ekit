/**
 * Description：
 * FileName：reverse_test.go
 * Author：CJiaの青姝
 * Create：2024/12/31 14:55:01
 * Remark：
 */

package slice

import "fmt"

func ExampleReverse() {
	res := Reverse[int]([]int{1, 3, 2, 2, 4})
	fmt.Println(res)
	res2 := Reverse[string]([]string{"a", "b", "c", "d", "e"})
	fmt.Println(res2)
	// Output:
	// [4 2 2 3 1]
	// [e d c b a]
}

func ExampleReverseSelf() {
	src := []int{1, 3, 2, 2, 4}
	ReverseSelf[int](src)
	fmt.Println(src)
	src2 := []string{"a", "b", "c", "d", "e"}
	ReverseSelf[string](src2)
	fmt.Println(src2)
	// Output:
	// [4 2 2 3 1]
	// [e d c b a]
}
