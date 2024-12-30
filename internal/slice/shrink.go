/**
 * Description：
 * FileName：shrink.go
 * Author：CJiaの青姝
 * Create：2024/12/30 14:11:15
 * Remark：
 */

package slice

// Shrink 缩减切片容量
func Shrink[T any](src []T) []T {
	l, c := len(src), cap(src)
	n, changed := calCapacity(l, c)
	if !changed {
		return src
	}
	s := make([]T, 0, n)
	s = append(s, src...)
	return s
}

// calCapacity 计算容量
func calCapacity(l int, c int) (int, bool) {
	if c <= 64 {
		return c, false
	}
	if c > 2048 && (c/l >= 2) {
		factor := 0.625
		return int(float32(c) * float32(factor)), true
	}
	if c <= 2048 && (c/l >= 4) {
		return c / 2, true
	}
	return c, true
}
