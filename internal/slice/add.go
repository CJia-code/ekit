/**
 * Description：
 * FileName：add.go
 * Author：CJiaの青姝
 * Create：2024/12/30 11:23:48
 * Remark：
 */

package slice

import "github.com/CJia-code/ekit/internal/errs"

// Add 添加元素到切片
// src: 原切片
// element: 要添加的元素
// index: 要添加的元素的位置
// 返回值：添加元素后的切片
func Add[T any](src []T, element T, index int) ([]T, error) {
	length := len(src)
	if length < 0 || index < 0 || index > length {
		return nil, errs.NewErrIndexOutOfRange(index, length)
	}
	// 声明一个比原切片长度多1的切片
	dst := make([]T, length+1)
	copy(dst, src[:index])
	dst[index] = element
	copy(dst[index+1:], src[index:])
	return dst, nil
}
