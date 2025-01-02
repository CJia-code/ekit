/**
 * Description：
 * FileName：intersect_test.go
 * Author：CJiaの青姝
 * Create：2024/12/31 14:01:14
 * Remark：
 */

package slice

import (
	"fmt"
	"sort"
)

func ExampleIntersectSet() {
	res := IntersectSet[int]([]int{1, 2, 3, 3, 4}, []int{1, 1, 3})
	sort.Ints(res)
	fmt.Println(res)
	res = IntersectSet[int]([]int{1, 2, 3, 3, 4}, []int{5, 7})
	fmt.Println(res)
	// Output:
	// [1 3]
	// []
}

func ExampleIntersectSetFunc() {
	res := IntersectSetFunc[int]([]int{1, 2, 3, 3, 4}, []int{1, 1, 3}, func(src, dst int) bool {
		return src == dst
	})
	sort.Ints(res)
	fmt.Println(res)
	res = IntersectSetFunc[int]([]int{1, 2, 3, 3, 4}, []int{5, 7}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	// Output:
	// [1 3]
	// []
}
