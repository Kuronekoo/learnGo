package main

import "fmt"

/*
	一、自定义类型是定义了一个全新的类型。我们可以基于内置的基本类型定义，也可以通过struct定义。例如：
	将MyInt定义为int类型
		type MyInt int
	通过type关键字的定义，MyInt就是一种新的类型，它具有int的特性。


	二、使用type和struct关键字来定义结构体，具体代码格式如下：
		type 类型名 struct {
		字段名 字段类型
		字段名 字段类型
		…
		}

	类型名：标识自定义结构体的名称，在同一个包内不能重复。
	字段名：表示结构体字段名。结构体中的字段名必须唯一。
	字段类型：表示结构体字段的具体类型。

	三、只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。

	结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。

	var 结构体实例 结构体类型

	四、内存相关
	结构体占用一块连续的内存。
	空结构体是不占用空间的。


*/
func main() {
	testStru()

}

func testStru() {
	var p1 person
	p1.name = "沙河娜扎"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("%T %v\n", p1, p1)  //p1={沙河娜扎 北京 18}
	fmt.Printf("%T %#v\n", p1, p1) //p1=main.person{name:"沙河娜扎", city:"北京", age:18}

	//匿名结构体
	//没有初始化的结构体，其成员变量都是对应其类型的零值。
	var user struct {
		Name string
		Age  int
	}
	user.Name = "小王子"
	user.Age = 18
	fmt.Printf("%T %#v\n", user, user)
	//p2是一个结构体指针
	var p2 = new(person)
	//Go语言中支持对结构体指针直接使用.来访问结构体的成员
	p2.name = "小王子"
	p2.age = 28
	p2.city = "上海"
	fmt.Printf("%T %#v\n", p2, p2)

	//使用&对结构体进行取地址操作，相当于对该结构体类型进行了一次new实例化操作。
	p3 := &person{}
	//p3.name = "七米"其实在底层是(*p3).name = "七米"，这是Go语言帮我们实现的语法糖。
	p3.name = "七米"
	p3.age = 30
	p3.city = "成都"

	fmt.Printf("%T %#v\n", p3, p3)

	//声明的同时对变量进行赋值，进行键值对初始化
	p5 := person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5)
	//使用&对结构体进行键值对初始化，可以只写部分字段
	p6 := &person{
		name: "小王子",
		city: "北京",
	}
	fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"小王子", city:"北京", age:18}
	//始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值：
	// 必须初始化结构体的所有字段。
	// 初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	// 该方式不能和键值初始化方式混用。
	p9 := newPerson("张三", "深圳", 90)
	fmt.Printf("%#v\n", p9) //&main.person{name:"张三", city:"沙河", age:90}

	testMethod()
}

//没有初始化的结构体，其成员变量都是对应其类型的零值。
type person struct {
	name, city string
	age        int8
}

//一个构造函数
//因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
