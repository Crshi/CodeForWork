package main

import (
	"fmt"
	"reflect"
)

// 函数在开始执行的时候，先初始化了它的返回值，如果未声明，随机一个tempS,如果声明了，即s
// 在执行到return（返回标志）的时候赋值。
// sum3函数先在执行到return时候，将tempS赋值为s，即3。由于s和tempS3是不同的变量，defer中的s+2就与tempS3无关。
// sum4函数由于声明了返回变量s，return时候，如果是 return x+y，其实将x+y复制给s，如果直接return，可以理解为跳过了复制操作。
// 在return之后的defer，s与sum3所谓的tempS3其实是一个，所以defer的操作是有效的。

func sum3(x, y int) int {
	s := x + y
	defer func() {
		s += 2 // 第二步，对s进行自增运算，不影响s1的值
	}()
	return s // 第一步，把值赋予给s1
}

func sum4(x, y int) (s int) {
	s = x + y
	defer func() {
		s += 2 // 第二步，对s进行自增运算，影响返回值s
	}()
	return // 第一步，把值赋予给s本身(s2:=x+y;return s2是一样的)
}

func main() {
	// s3 := sum3(1, 2)
	// s4 := sum4(1, 2)
	// println(s3) // 3
	// println(s4) // 5

	var a int
	// fn1
	defer func() {
		a = 3
		if err := recover(); err != nil {
			a = 4
			fmt.Println("++++")
			f := err.(func() string)
			fmt.Println(err, f(), reflect.TypeOf(err).Kind().String())
		} else {
			fmt.Println("fatal")
		}
	}()
	// fn2
	defer func() {
		a = 2
		if r := recover(); r != nil { // 这里的recover()去掉感受一下效果
			panic(r)
		}
		panic(func() string {
			return "defer panic"
		})
	}()
	a = 1
	panic("panic1") // 这里的panic去掉感受一下效果
	// panic("panic2")
}
