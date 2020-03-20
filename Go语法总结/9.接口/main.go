package main

import "fmt"

//定义一个接口：
// 一个或多个方法的集合，叫接口，如果一个方法都没有，叫空接口
//实现类：对于空接口，
type Animal3 interface {
	eat()
}

//空接口
type Animal4 interface {
}
type Dog3 struct {
	name string
}
type Cat3 struct {
	name string
}
type Rabbit struct {
	name string
}

func (d Dog3) eat() {
	fmt.Println("小狗，吃东西。。")
}
func (c Cat3) eat() {
	fmt.Println("小猫，吃东西。。")
}

func test(a Animal3) { // 接收dog3，cat3

}

// 空接口作为参数：代表可以接收任意类型
func test2(a Animal4) { // dog3,cat3,rabbit
	fmt.Println(a)
}

func main() {
	// Dog,Cat,是Animal的实现类,因为都实现类eat()

	d1 := Dog3{"啸天"}
	d2 := Dog3{"二黄"}
	c1 := Cat3{"咪咪"}
	c2 := Cat3{"豆豆"}

	as := [4]Animal3{d1, d2, c1, c2}
	fmt.Println(as)

	test2(d1)
	test2(c1)
	test2(Rabbit{"玉兔"})
	test2(100)

	//面向对象：面向接口编程--->面向实现类编程
	// 对象.方法()

	as2 := [4]Animal4{} //如果定义空接口的数组类型，那么实际上可以存入任意类型的数据。
	fmt.Println(as2)
	as2[0] = d1
	as2[1] = c1
	as2[3] = "memeda"
	fmt.Println(as2)
}
