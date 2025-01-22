/**
 * Description：
 * FileName：delete.go
 * Author：CJiaの青姝
 * Create：2024/12/30 11:41:42
 * Remark：
 */

package slice

import "github.com/CJia-code/ekit/internal/errs"

// Delete 删除切片中的元素
// src: 原切片
// index: 要删除的元素的位置
// 返回值：删除元素后的切片，被删除的元素
func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		var zero T
		return nil, zero, errs.NewErrIndexOutOfRange(length, index)
	}
	result := src[index]
	// 从index位置开始，后面的元素依次往前挪1个位置
	for i := index; i+1 < length; i++ {
		src[i] = src[i+1]
	}
	// 去掉最后一个重复元素
	src = src[:length-1]
	return src, result, nil
}
