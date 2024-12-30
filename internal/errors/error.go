/**
 * Description：
 * FileName：error.go
 * Author：CJiaの青姝
 * Create：2024/12/30 11:16:08
 * Remark：
 */

package errors

import (
	"fmt"
	"time"
)

// NewErrIndexOutOfRange 创建异常：索引越界
func NewErrIndexOutOfRange(index int, length int) error {
	return fmt.Errorf("【异常】索引 %d 超出范围 %d", index, length)
}

// NewErrInvalidType 创建异常：类型转换失败
func NewErrInvalidType(want string, got any) error {
	return fmt.Errorf("【异常】 类型转换失败，预期类型 %s, 实际值 %v", want, got)
}

// NewErrInvalidIntervalValue 创建异常：无效的间隔时间
func NewErrInvalidIntervalValue(interval time.Duration) error {
	return fmt.Errorf("ekit: 无效的间隔时间 %d, 预期值应大于 0", interval)
}

// NewErrInvalidMaxIntervalValue 创建异常：无效的最大重试间隔时间
func NewErrInvalidMaxIntervalValue(maxInterval, initialInterval time.Duration) error {
	return fmt.Errorf("ekit: 最大重试间隔的时间 [%d] 应大于等于初始重试的间隔时间 [%d] ", maxInterval, initialInterval)
}

// NewErrRetryExhausted 创建异常：超过最大重试次数
func NewErrRetryExhausted(lastErr error) error {
	return fmt.Errorf("ekit: 超过最大重试次数，业务返回的最后一个 error %w", lastErr)
}
