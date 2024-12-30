/**
 * Description：
 * FileName：map.go
 * Author：CJiaの青姝
 * Create：2024/12/30 15:21:42
 * Remark：
 */

package slice

// toMap 构造map
func toMap[T comparable](src []T) map[T]struct{} {
	var dataMap = make(map[T]struct{}, len(src))
	for _, v := range src {
		// 使用空结构体,减少内存消耗
		dataMap[v] = struct{}{}
	}
	return dataMap
}
