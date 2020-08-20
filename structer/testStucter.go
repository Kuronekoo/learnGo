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

	结构体字段的可见性
	结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

	三、只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。

	结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。

	var 结构体实例 结构体类型

	四、内存相关
	结构体占用一块连续的内存。
	空结构体是不占用空间的。

	在Go语言中，访问结构体指针的成员变量时可以继续使用 '.' 。
	这是因为Go语言为了方便开发者访问结构体指针的成员变量，使用了语法糖（Syntactic sugar）技术，将 ins.Name 形式转换为 (*ins).Name。

	五、结构体与JSON序列化

	六、结构体标签
	Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。 Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：

	`key1:"value1" key2:"value2"`
	结构体tag由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。同一个结构体字段可以设置多个键值对tag，不同的键值对之间使用空格分隔。
	注意事项： 为结构体编写Tag时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。例如不要在key和value之间添加空格。

*/
func main() {
	fmt.Println("--------------testStru--------------")
	testStru()
	fmt.Println("--------------testMethod--------------")
	testMethod()
	fmt.Println("--------------testJson--------------")
	testJson()
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
	//创建一个空结构
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

	//长的和json一毛一样
	user1 := User{
		Name:   "小王子",
		Gender: "男",
		Address: Address{
			Province: "山东",
			City:     "威海",
		},
	}
	fmt.Printf("user1=%#v\n", user1) //user1=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}
	d1 := &Dog{
		Feet: 4,
		//注意嵌套的是结构体指针
		Animal: &Animal{
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！

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

//一个构造函数
//可以按照参数的声明顺序直接创建一个persoon实例
func newPersonSimple(name, city string, age int8) *person {
	return &person{name, city, age}
}

/*
	Alien struct
	结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。
	注意：这里匿名字段的说法并不代表没有字段名，而是默认会采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。
	比如下面的两个字段的名称就依次叫做 string int
*/
type Alien struct {
	string
	int
}

//Address 地址结构体
type Address struct {
	Province string
	City     string
}

//User 用户结构体
type User struct {
	Name   string
	Gender string
	//这个字段可以匿名
	//即只写一个Address
	//匿名字段可以省略，可以直接通过上级结构给嵌套结构体赋值 user2.City = "威海"
	//当访问结构体成员时会先在结构体中查找该字段，找不到再去嵌套的匿名字段中查找。
	Address Address
}

//实现继承
//Animal 动物
type Animal struct {
	name string
}

// 定义属于Animal方法的move
func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

//调用嵌套匿名结构的属性
func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}
