/**
 * Description：
 * FileName：main.go
 * Author：CJiaの青姝
 * Create：2024/12/30 15:09:45
 * Remark：
 */

package main

import (
	"fmt"
	"github.com/CJia-code/ekit/slice"
)

func main() {
	val := []int{1, 2, 3, 4, 5}
	fmt.Println(slice.FilterDelete(val, func(index int, val int) bool {
		return index == 2
	}))
	fmt.Println(slice.Contains(val, 13))
	fmt.Println(slice.ContainsAny(val, []int{1, 2, 3}))
}
