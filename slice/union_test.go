/**
 * Description：
 * FileName：union_test.go
 * Author：CJiaの青姝
 * Create：2024/12/31 14:54:40
 * Remark：
 */

package slice

import (
	"fmt"
	"sort"
)

func ExampleUnionSet() {
	res := UnionSet[int]([]int{1, 3, 4, 5}, []int{1, 4, 7})
	sort.Ints(res)
	fmt.Println(res)
	// Output:
	// [1 3 4 5 7]
}

func ExampleUnionSetFunc() {
	res := UnionSetFunc[int]([]int{1, 3, 4, 5}, []int{1, 4, 7}, func(src, dst int) bool {
		return src == dst
	})
	sort.Ints(res)
	fmt.Println(res)
	// Output:
	// [1 3 4 5 7]
}
