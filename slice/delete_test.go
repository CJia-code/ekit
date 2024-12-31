/**
 * Description：
 * FileName：delete_test.go
 * Author：CJiaの青姝
 * Create：2024/12/30 15:48:56
 * Remark：
 */

package slice

import "fmt"

func ExampleDelete() {
	res, _ := Delete[int]([]int{1, 2, 3, 4}, 2)
	fmt.Println(res)
	_, err := Delete[int]([]int{1, 2, 3, 4}, -1)
	fmt.Println(err)
	// Output:
	// [1 2 4]
	// 【异常】索引 4 超出范围 -1
}