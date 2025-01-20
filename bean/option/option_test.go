/**
 * Description：
 * FileName：option_test.go
 * Author：CJiaの青姝
 * Create：2025/1/20 10:34:25
 * Remark：
 */

package option

import (
	"errors"
	"fmt"
)

type User struct {
	name string
	age  int
}

func WithName(name string) Option[User] {
	return func(u *User) {
		u.name = name
	}
}

func WithAge(age int) Option[User] {
	return func(u *User) {
		u.age = age
	}
}

func ExampleApply() {
	u := &User{}
	Apply[User](u, WithName("Tom"), WithAge(18))
	fmt.Println(u)
	// Output:
	// &{Tom 18}
}

func WithNameErr(name string) OptionErr[User] {
	return func(u *User) error {
		if name == "" {
			return errors.New("name 不能为空")
		}
		u.name = name
		return nil
	}
}

func WithAgeErr(age int) OptionErr[User] {
	return func(u *User) error {
		if age <= 0 {
			return errors.New("age 应该是正数")
		}
		u.age = age
		return nil
	}
}

func ExampleApplyErr() {
	u := &User{}
	err := ApplyErr[User](u, WithNameErr("Tom"), WithAgeErr(18))
	fmt.Println(err)
	fmt.Println(u)

	err = ApplyErr[User](u, WithNameErr(""), WithAgeErr(18))
	fmt.Println(err)
	// Output:
	// <nil>
	// &{Tom 18}
	// name 不能为空
}
