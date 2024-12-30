/**
 * Description：
 * FileName：add.go
 * Author：CJiaの青姝
 * Create：2024/12/30 14:27:37
 * Remark：
 */

package slice

import "github.com/CJia-code/ekit/internal/slice"

// Add 在index处添加元素
// index 必须小于等于len(src)
// 若 index == len(src) 则表示在末尾添加
func Add[T any](src []T, element T, index int) ([]T, error) {
	res, err := slice.Add[T](src, element, index)
	return res, err
}
