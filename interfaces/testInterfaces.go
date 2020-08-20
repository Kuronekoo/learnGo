package main

import "fmt"

/*
	在Go语言中接口（interface）是一种类型，一种抽象的类型。

	interface是一组method的集合，是duck-type programming的一种体现。接口做的事情就像是定义一个协议（规则），只要一台机器有洗衣服和甩干的功能，我就称它为洗衣机。不关心属性（数据），只关心行为（方法）。

	为了保护你的Go语言职业生涯，请牢记接口（interface）是一种类型。

	一、接口的定义

		Go语言提倡面向接口编程。

		每个接口由数个方法组成，接口的定义格式如下：

		type 接口类型名 interface{
			方法名1( 参数列表1 ) 返回值列表1
			方法名2( 参数列表2 ) 返回值列表2
			…
		}
		其中：

		1.接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
		2.方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
		3.参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。

	二、实现接口
		一个对象只要全部实现了接口中的方法，那么就实现了这个接口。换句话说，接口就是一个需要实现的方法列表。

		我们来定义一个Sayer接口：

		// Sayer 接口
		type Sayer interface {
			say()
		}
		定义dog和cat两个结构体：

		type dog struct {}

		type cat struct {}

		因为Sayer接口里只有一个say方法，所以我们只需要给dog和cat 分别实现say方法就可以实现Sayer接口了。

		// dog实现了Sayer接口
		func (d dog) say() {
			fmt.Println("汪汪汪")
		}

		// cat实现了Sayer接口
		func (c cat) say() {
			fmt.Println("喵喵喵")
		}
		接口的实现就是这么简单，只要实现了接口中的所有方法，就实现了这个接口。

	三、接口类型变量
		接口类型变量能够存储所有实现了该接口的实例。 例如上面的示例中，Sayer类型的变量能够存储dog和cat类型的变量。
		func main() {
		var x Sayer // 声明一个Sayer类型的变量x
		a := cat{}  // 实例化一个cat
		b := dog{}  // 实例化一个dog
		x = a       // 可以把cat实例直接赋值给x
		x.say()     // 喵喵喵
		x = b       // 可以把dog实例直接赋值给x
		x.say()     // 汪汪汪
	}
	Tips： 观察下面的代码，体味此处_的妙用

	// 摘自gin框架routergroup.go
	type IRouter interface{ ... }

	type RouterGroup struct { ... }

	var _ IRouter = &RouterGroup{}  // 确保RouterGroup实现了接口IRouter

	四、空接口
	空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。

	空接口类型的变量可以存储任意类型的变量。

		func main() {
			// 定义一个空接口x
			var x interface{}
			s := "Hello 沙河"
			x = s
			fmt.Printf("type:%T value:%v\n", x, x)
			i := 100
			x = i
			fmt.Printf("type:%T value:%v\n", x, x)
			b := true
			x = b
			fmt.Printf("type:%T value:%v\n", x, x)
		}

	空接口作为函数的参数
	使用空接口实现可以接收任意类型的函数参数。
		// 空接口作为函数参数
		func show(a interface{}) {
			fmt.Printf("type:%T value:%v\n", a, a)
		}
	空接口作为map的值
	使用空接口实现可以保存任意值的字典。

		// 空接口作为map值
		var studentInfo = make(map[string]interface{})
		studentInfo["name"] = "沙河娜扎"
		studentInfo["age"] = 18
		studentInfo["married"] = false
		fmt.Println(studentInfo)
	五、类型断言
		一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值。
		想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：
			x.(T)
		其中：
		x：表示类型为interface{}的变量
		T：表示断言x可能是的类型。
		该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。
			func main() {
			var x interface{}
			x = "Hello 沙河"
			v, ok := x.(string)
			if ok {
				fmt.Println(v)
			} else {
				fmt.Println("类型断言失败")
			}

			func justifyType(x interface{}) {
				switch v := x.(type) {
				case string:
					fmt.Printf("x is a string，value is %v\n", v)
				case int:
					fmt.Printf("x is a int is %v\n", v)
				case bool:
					fmt.Printf("x is a bool is %v\n", v)
				default:
					fmt.Println("unsupport type！")
				}
			}


	关于接口需要注意的是，只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗。

*/
func main() {
	testInterface()
}

func testInterface() {
	var p Person = &Man{}
	p.Walk()
	//无法通过编译，因为实现类的方法是指针接收者
	// var p1 Person = Man{}
	// p1.Walk()

}

type Person interface {
	Walk()
}
type Man struct {
}

func (p *Man) Walk() {
	fmt.Println("man walking")
}
