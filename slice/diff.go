/**
 * Description：
 * FileName：diff.go
 * Author：CJiaの青姝
 * Create：2024/12/31 10:21:47
 * Remark：
 */

package slice

// DiffSet 差集，只支持 comparable 类型
// 已去重
// 并且返回值的顺序是不确定的
func DiffSet[T comparable](src []T, dst []T) []T {
	srcMap := toMap[T](src)
	for _, v := range dst {
		delete(srcMap, v)
	}
	var ret = make([]T, 0, len(srcMap))
	for key := range srcMap {
		ret = append(ret, key)
	}
	return ret
}

// DiffSetFunc 差集，已去重
// 优先使用 DiffSet
func DiffSetFunc[T any](src, dst []T, equal EqualFunc[T]) []T {
	var ret = make([]T, 0, len(src))
	for _, v := range src {
		if ContainsFunc[T](dst, func(src T) bool {
			return equal(v, src)
		}) {
			ret = append(ret, v)
		}
	}
	return ret
}
