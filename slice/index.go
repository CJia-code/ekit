/**
 * Description：
 * FileName：index.go
 * Author：CJiaの青姝
 * Create：2024/12/31 11:32:48
 * Remark：
 */

package slice

// Index 返回和 target 相等的第一个元素下标
// -1 表示没找到
func Index[T comparable](src []T, target T) int {
	return IndexFunc[T](src, func(val T) bool {
		return val == target
	})
}

// IndexFunc 返回 match 返回 true 的第一个下标
// -1 表示没找到
// 优先使用 Index
func IndexFunc[T any](src []T, match MatchFunc[T]) int {
	for k, v := range src {
		if match(v) {
			return k
		}
	}
	return -1
}

// LastIndex 返回和 dst 相等的最后一个元素下标
// -1 表示没找到
func LastIndex[T comparable](src []T, dst T) int {
	return LastIndexFunc[T](src, func(val T) bool {
		return val == dst
	})
}

// LastIndexFunc 返回和 dst 相等的最后一个元素下标
// -1 表示没找到
// 优先使用 LastIndex
func LastIndexFunc[T any](src []T, match MatchFunc[T]) int {
	// 反向遍历
	for i := len(src) - 1; i >= 0; i-- {
		if match(src[i]) {
			return i
		}
	}
	return -1
}

// IndexAll 返回和 dst 相等的所有元素的下标
func IndexAll[T comparable](src []T, dst T) []int {
	return IndexAllFunc[T](src, func(val T) bool {
		return val == dst
	})
}

// IndexAllFunc 返回和 match 返回 true 的所有元素的下标
// 你应该优先使用 IndexAll
func IndexAllFunc[T any](src []T, match MatchFunc[T]) []int {
	var indexes = make([]int, 0, len(src))
	for k, v := range src {
		if match(v) {
			indexes = append(indexes, k)
		}
	}
	return indexes
}
