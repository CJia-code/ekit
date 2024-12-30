/**
 * Description：
 * FileName：delete.go
 * Author：CJiaの青姝
 * Create：2024/12/30 15:48:51
 * Remark：
 */

package slice

import (
	"fmt"
	"github.com/CJia-code/ekit/internal/slice"
)

// Delete 删除index处的元素
func Delete[T any](src []T, index int) ([]T, error) {
	res, _, err := slice.Delete[T](src, index)
	return res, err
}

func FilterDelete[T any](src []T, m func(index int, val T) bool) []T {
	// 记录被删除的元素位置，也称空缺的位置
	emptyPos := 0
	for index := range src {
		fmt.Println(index)
		// 判断是否满足删除的条件
		if m(index, src[index]) {
			continue
		}
		// 移动元素
		src[emptyPos] = src[index]
		emptyPos++
	}
	return src[:emptyPos]
}
