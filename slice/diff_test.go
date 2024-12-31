/**
 * Description：
 * FileName：diff_test.go
 * Author：CJiaの青姝
 * Create：2024/12/31 10:21:55
 * Remark：
 */

package slice

import (
	"fmt"
	"sort"
)

func ExampleDiffSet() {
	res := DiffSet[int]([]int{1, 3, 2, 2, 4}, []int{3, 4, 5, 6})
	sort.Ints(res)
	fmt.Println(res)
	// Output:
	// [1 2]
}

func ExampleDiffSetFunc() {
	res := DiffSetFunc[int]([]int{1, 3, 2, 2, 4}, []int{3, 4, 5, 6}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	// Output:
	// [1 2]
}
