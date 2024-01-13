package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {

	//test1()
	//test2()
	//test3()
	//test4(100)() //函数作为返回值的一种写法
	//
	//x := 100 //函数作为返回值的另一种写法
	//f := test4(x)
	//f()
	//
	//test6()
	//test7()
	//test8()
	//test9()

	//test11()

	test12()

}

// 变量
func test1() {
	var x int32
	var s = "hello hello!"
	println(x, s)
}

// 表达式： if switch for
func test2() {
	x := 100

	println("if test")

	if x > 0 {
		println("x")
	} else if x < 0 {
		println("-x")
	} else {
		println("0")
	}

	println("switch test")

	switch {
	case x > 0:
		println("x")
	case x < 0:
		println("-x")
	default:
		println("0")

	}

	for i := 0; i < 5; i++ {
		print(i)
	}

	for i := 4; i >= 0; i-- {
		print(i)
	}

	println()
	println("for other usage")

	y := 0
	for y < 5 { // 类似于 while(x < 5) {...} ;  go 里面没有while
		print(y)
		y++
	}

	z := 4
	for { //类似于 while(true){...} ; go  里面没有while
		print(z)
		z--

		if z < 0 {
			break
		}
	}

	println()
	println("for ... range的用法")

	m := []int{100, 101, 102}

	for i, n := range m {
		println(i, ":", n)
	}

	println()
}

// 函数：
// （1）函数可以定义多个返回值，并且对其命名
// （2）函数是第一类型，可以作为参数或者返回值
// （3）用defer 定义延迟调用/ 也就是不管函数是否有出错，他都确保结束之前被调用

func test3() {
	a, b := 10, 0       //定义多个变量
	c, err := div(a, b) // 接收多个返回值

	fmt.Println(c, err)

	test5(a, b)
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

func test4(x int) func() { // 返回函数类型
	return func() { // 匿名函数
		println(x) // 闭包： 函数和其引用环境组合的一个整体（实体）
	}
}

// defer 用法
func test5(a, b int) {
	defer println("dispose...") // 延迟调用，函数结束前调用
	//div(a, b)
	println(a / b)
}

// 数据： slice  map  struct
func test6() {
	x := make([]int, 0, 5) //创建一个切片，长度为0，容量为5

	for i := 0; i < 8; i++ { //追加数据。 当超出容量限制时，自动分配更大的存储空间
		x = append(x, i)
	}

	fmt.Println(x)
}

func test7() {
	m := make(map[string]int) //创建字典类型对象

	m["a"] = 1       //添加或者设置
	x, ok1 := m["a"] //获取元素，第二个返回值表示是否存在

	y, ok2 := m["b"]

	println(x, ok1)
	println(y, ok2)
	delete(m, "a") //删除元素
	println(m["a"])
}

type user struct { //定义结构类型
	name string
	age  byte
}

type manager struct {
	user  //匿名嵌入其他类型
	title string
}

func test8() {
	var m manager

	m.name = "Tom" //直接访问匿名字段的成员
	m.age = 29

	m.title = "CTO" //访问嵌入字段的成员

	fmt.Println(m)
	m.ToString() //继承了user的方法！！！

	println("接口的用法")
	var p Pringer = m //接口类型变量
	println(m.ToString())
	println(p.ToString())
}

//方法  接口  并发

// 1. 方法: 可以为当前包内任意类型定义方法，不仅仅是结构体，但是不能为其他包内的类型定义方法

type X int

func (x *X) inc() { // 为X类型定义一个方法
	*x++
}

func test9() {
	// 对某个类重命名，相当于是类型的别名
	var x X
	x.inc()
	println(x)
}

// 实现与继承类似的功能
func (u user) ToString() string { // 为user类型定义一个方法
	return fmt.Sprintf("%+v", u)
}

// 接口
type Pringer interface { //定义接口类型
	ToString() string //只定义方法，不定义成员变量
}

// 并发
func task10(id int) {

	for id := 0; id < 10; id++ {
		println(id)
		time.Sleep(time.Second)
	}
}

func test11() {
	go task10(1) //启动goroutine
	go task10(2)

	time.Sleep(time.Second * 5)
}

// 通道 : channel 与 goroutine 搭配使用， 实现用通信代替内存共享的 CSP 模型

func consumer(data chan int, done chan bool) {
	for x := range data { //接收数据，直到通道被关闭
		println("recv:", x)
	}

	done <- true //通知main，消费结束
}

func producer(data chan int) {
	for i := 0; i < 4; i++ {
		data <- i //发送数据
	}

	close(data) //生产结束，关闭通道)  {

}
func test12() {

	done := make(chan bool) //用于接收消费结束信号
	data := make(chan int)  //数据管道

	go consumer(data, done) //启动消费者
	go producer(data)       //启动生产者

	<-done //阻塞，直到消费者发回结束信号

}
