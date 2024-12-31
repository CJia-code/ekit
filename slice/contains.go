/**
 * Description：
 * FileName：contains.go
 * Author：CJiaの青姝
 * Create：2024/12/30 15:05:21
 * Remark：
 */

package slice

// Contains 判断切片中是否包含某个元素
func Contains[T comparable](src []T, val T) bool {
	return ContainsFunc[T](src, func(compare T) bool {
		return compare == val
	})
}

// ContainsFunc 判断切片中是否包含某个元素
func ContainsFunc[T any](src []T, equal func(compare T) bool) bool {
	// 遍历调用equal函数进行判断
	for _, v := range src {
		if equal(v) {
			return true
		}
	}
	return false
}

// ContainsAny 判断切片中是否存在val中的任何一个元素
func ContainsAny[T comparable](src, val []T) bool {
	srcMap := toMap[T](src)
	for _, v := range val {
		if _, exist := srcMap[v]; exist {
			return true
		}
	}
	return false
}

// ContainsAnyFunc 判断切片中是否存在val中的任何一个元素
// 优先使用 ContainsAny
func ContainsAnyFunc[T any](src, val []T, equal EqualFunc[T]) bool {
	for _, valDst := range val {
		for _, valSrc := range src {
			if equal(valSrc, valDst) {
				return true
			}
		}
	}
	return false
}

// ContainsAll 判断切片中是否存在val中的所有元素
func ContainsAll[T comparable](src, val []T) bool {
	srcMap := toMap[T](src)
	for _, v := range val {
		if _, exist := srcMap[v]; !exist {
			return false
		}
	}
	return true
}

// ContainsAllFunc 判断切片中是否存在val中的所有元素
// 优先使用 ContainsAll
func ContainsAllFunc[T any](src, val []T, equal EqualFunc[T]) bool {
	for _, valDst := range val {
		if !ContainsFunc[T](src, func(src T) bool {
			return equal(src, valDst)
		}) {
			return false
		}
	}
	return true
}
