/**
 * Description：
 * FileName：union.go
 * Author：CJiaの青姝
 * Create：2024/12/31 14:54:12
 * Remark：
 */

package slice

// UnionSet 并集，只支持 comparable
// 已去重
// 返回值的元素顺序是不定的
func UnionSet[T comparable](src, dst []T) []T {
	srcMap, dstMap := toMap[T](src), toMap[T](dst)
	for key := range srcMap {
		dstMap[key] = struct{}{}
	}

	var ret = make([]T, 0, len(dstMap))
	for key := range dstMap {
		ret = append(ret, key)
	}

	return ret
}

// UnionSetFunc 并集，支持任意类型
// 优先使用 UnionSet
// 已去重
func UnionSetFunc[T any](src, dst []T, equal EqualFunc[T]) []T {
	var ret = make([]T, 0, len(src)+len(dst))
	ret = append(ret, dst...)
	ret = append(ret, src...)

	return deduplicateFunc[T](ret, equal)
}
