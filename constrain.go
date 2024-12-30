/**
 * Description：泛型约束
 * FileName：constrain.go
 * Author：CJiaの青姝
 * Create：2024/12/30 14:47:39
 * Remark：
 */

package ekit

// RealNumber 实数
// 本类型以及它的衍生类型，一网打尽
// 绝大多数情况下，你都应该用这个来表达数字的含义
type RealNumber interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

type Number interface {
	RealNumber | ~complex64 | ~complex128
}
