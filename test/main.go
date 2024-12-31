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
	vals := slice.FindAll[int]([]int{2, 3, 4}, func(src int) bool {
		return src%2 == 1
	})
	fmt.Println(vals)
	fmt.Println(3 >> 8)
}
