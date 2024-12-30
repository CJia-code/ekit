/**
 * Description：
 * FileName：types.go
 * Author：CJiaの青姝
 * Create：2024/12/30 15:28:21
 * Remark：
 */

package slice

// EqualFunc 比较两个元素是否相等
type EqualFunc[T any] func(src, dst T) bool

type MatchFunc[T any] func(src T) bool
