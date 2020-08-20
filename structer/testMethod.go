package main

import "fmt"

/*
	Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的this或者 self。

	方法的定义格式如下：

		func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
			函数体
		}
	方法和函数对比，在方法名的前面多了一个接收的参数
		函数： func 函数(参数列表) (返回参数){}
		1.接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，而不是self、this之类的命名。例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
		2.接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
		3.方法名、参数列表、返回参数：具体格式与函数定义相同。

	什么时候应该使用指针类型接收者
		1.需要修改接收者中的值
		2.接收者是拷贝代价比较大的大对象
		3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。


	任意类型添加方法
		在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。
		举个例子，我们基于内置的int类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。
	注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。




*/
func testMethod() {
	p1 := NewPerson("小王子", 25)
	p1.show()
	p1.Dream()
	p1.SetAge(30)
	p1.show()
	//这里设置age不生效
	p1.SetAge2(50)
	p1.show()

	var m1 MyInt
	m1.SayHello() //Hello, 我是一个int。
	m1 = 100
	fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt

}

//Person 结构体
type Person struct {
	name string
	age  int8
}

// NewPerson Person 的构造函数
//因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//Dream 属于Person结构的Dream方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

// SetAge 设置p的年龄
// ******使用指针接收者******
//指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。
//这种方式就十分接近于其他语言中面向对象中的this或者self。
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// SetAge2 设置p的年龄
// ******使用值接收者******
// 使用值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。
// 在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。
// 这个方法中，改了age之后person的age不变
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

func (p *Person) show() {
	fmt.Printf("%v \n", p)
}

//MyInt 将int定义为自定义MyInt类型
type MyInt int

//SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}
