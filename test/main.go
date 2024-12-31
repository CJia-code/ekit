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
	any1 := slice.ContainsAny([]int{1, 2, 3}, []int{3, 8})
	fmt.Println(any1)
}
