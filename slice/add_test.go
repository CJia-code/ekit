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
	res, _ := Add[int]([]int{1, 2, 4}, 2, 2)
	fmt.Println(res)
}
