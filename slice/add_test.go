/**
 * Description：
 * FileName：add_test.go
 * Author：CJiaの青姝
 * Create：2024/12/30 14:29:08
 * Remark：
 */

package slice

import "fmt"

func ExampleAdd() {
	res, _ := Add[int]([]int{1, 2, 3, 4}, 233, 2)
	fmt.Println(res)
	_, err := Add[int]([]int{1, 2, 3, 4}, 233, -1)
	fmt.Println(err)
	// Output:
	// [1 2 233 3 4]
	// 下标超出范围，长度 4, 下标 -1
}
