/**
 * Description：
 * FileName：symmetric_diff_test.go
 * Author：CJiaの青姝
 * Create：2024/12/31 14:54:32
 * Remark：
 */

package slice

import (
	"fmt"
	"sort"
)

func ExampleSymmetricDiffSet() {
	res := SymmetricDiffSet[int]([]int{1, 3, 4, 2}, []int{2, 5, 7, 3})
	sort.Ints(res)
	fmt.Println(res)
	// Output:
	// [1 4 5 7]
}

func ExampleSymmetricDiffSetFunc() {
	res := SymmetricDiffSetFunc[int]([]int{1, 3, 4, 2}, []int{2, 5, 7, 3}, func(src, dst int) bool {
		return src == dst
	})
	sort.Ints(res)
	fmt.Println(res)
	// Output:
	// [1 4 5 7]
}
