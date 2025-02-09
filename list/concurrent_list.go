/**
 * Description：
 * FileName：concurrent_list.go
 * Author：CJiaの青姝
 * Create：2025/1/22 15:41:51
 * Remark：
 */

package list

import "sync"

var (
	_ List[any] = &ConcurrentList[any]{}
)

// ConcurrentList 用读写锁封装了对 List 的操作
// 达到线程安全的目标
type ConcurrentList[T any] struct {
	List[T]
	lock sync.RWMutex
}

// NewConcurrentListOfSlice 初始化一个ArrayList，并延续接口的方法
func NewConcurrentListOfSlice[T any](ts []T) *ConcurrentList[T] {
	var list List[T] = NewArrayListOf(ts)
	return &ConcurrentList[T]{List: list}
}

func (c *ConcurrentList[T]) Get(index int) (T, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.List.Get(index)
}

func (c *ConcurrentList[T]) Append(ts ...T) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.List.Append(ts...)
}

func (c *ConcurrentList[T]) Add(index int, t T) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.List.Add(index, t)
}

func (c *ConcurrentList[T]) Set(index int, t T) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.List.Set(index, t)
}

func (c *ConcurrentList[T]) Delete(index int) (T, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.List.Delete(index)
}

func (c *ConcurrentList[T]) Len() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.List.Len()
}

func (c *ConcurrentList[T]) Cap() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.List.Cap()
}

func (c *ConcurrentList[T]) Range(fn func(index int, t T) error) error {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.List.Range(fn)
}

func (c *ConcurrentList[T]) AsSlice() []T {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.List.AsSlice()
}
