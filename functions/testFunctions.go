package main

import (
	"bytes"
	"fmt"
	"math"
	"sync"
)

/*
	函数是组织好的、可重复使用的、用来实现单一或相关联功能的代码段，其可以提高应用的模块性和代码的重复利用率。

	Go 语言支持普通函数、匿名函数和闭包，从设计上对函数进行了优化和改进，让函数使用起来更加方便。

	Go 语言的函数属于“一等公民”（first-class），也就是说：
	函数本身可以作为值进行传递。
	支持匿名函数和闭包（closure）。
	函数可以满足接口。

	函数构成了代码执行的逻辑结构，在Go语言中，函数的基本组成为：关键字 func、函数名、参数列表、返回值、函数体和返回语句，每一个程序都包含很多的函数，函数是基本的代码块。

	因为Go语言是编译型语言，所以函数编写的顺序是无关紧要的，鉴于可读性的需求，最好把 main() 函数写在文件的前面，其他函数按照一定逻辑顺序进行编写（例如函数被调用的顺序）。

	编写多个函数的主要目的是将一个需要很多行代码的复杂问题分解为一系列简单的任务来解决，而且，同一个任务（函数）可以被多次调用，有助于代码重用（事实上，好的程序是非常注意 DRY 原则的，即不要重复你自己（Don't Repeat Yourself），意思是执行特定任务的代码只能在程序里面出现一次）。

	当函数执行到代码块最后一行}之前或者 return 语句的时候会退出，其中 return 语句可以带有零个或多个参数，这些参数将作为返回值供调用者使用，简单的 return 语句也可以用来结束 for 的死循环，或者结束一个协程（goroutine）。

	Go语言里面拥三种类型的函数：
	* 普通的带有名字的函数
	* 匿名函数或者 lambda 函数
	* 方法

	函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体。
		func 函数名(形式参数列表)(返回值列表){
			函数体
		}
	形式参数列表 描述了函数的参数名以及参数类型，这些参数作为局部变量，其值由参数调用者提供。
	返回值列表描 述了函数返回值的变量名以及类型，如果函数返回一个无名变量或者没有返回值，返回值列表的括号是可以省略的。


	如果一个函数在声明时，包含返回值列表，那么该函数必须以 return 语句结尾，除非函数明显无法运行到结尾处，例如函数在结尾时调用了 panic 异常或函数中存在无限循环。

	相同类型的函数：函数的出入参数的数量、格式、顺序一样

	在函数中，实参通过值传递的方式进行传递，因此函数的形参是实参的拷贝，对形参进行修改不会影响实参。
	但是，如果实参包括引用类型，如指针、slice(切片)、map、function、channel 等类型，实参可能会由于函数的间接引用被修改。

	Go语言支持多返回值，多返回值能方便地获得函数执行后的多个返回参数，Go语言经常使用多返回值中的最后一个返回参数返回函数执行中可能发生的错误，示例代码如下：
	conn, err := connectToNetwork()

	在Go语言中，函数也是一种类型，可以和其他类型一样保存在变量中


	Go语言支持匿名函数，即在需要使用函数时再定义函数，匿名函数没有函数名只有函数体，函数可以作为一种类型被赋值给函数类型的变量，匿名函数也往往以变量方式传递
	匿名函数是指不需要定义函数名的一种函数实现方式，由一个不带函数名的函数声明和函数体组成，下面来具体介绍一下匿名函数的定义及使用。
	匿名函数的定义格式如下：
		func(参数列表)(返回参数列表){
			函数体
		}
	匿名函数的用途非常广泛，它本身就是一种值，可以方便地保存在各种容器中实现回调函数和操作封装。

*/
func main() {
	fmt.Println("--------------testFunc--------------")
	testFunc()
	fmt.Println("--------------testClosure--------------")
	testClosure()
	fmt.Println("--------------testDefer--------------")
	testDefer()
}

func testFunc() {
	a, b := hypot(3, 4)
	fmt.Println(a, b)
	//	在Go语言中，函数也是一种类型，可以和其他类型一样保存在变量中
	//变量 f 声明为 func() 类型，此时 f 就被俗称为“回调函数”，此时 f 的值为 nil。
	var f func()
	//将 fire() 函数作为值，赋给函数变量 f，此时 f 的值为 fire() 函数。
	f = fire
	//使用函数变量 f 进行函数调用，实际调用的是 fire() 函数。
	f()

	//匿名函数可以在声明后调用
	func(data int) {
		fmt.Println("hello", data)
	}(100)

	// 将匿名函数体保存到f()中
	f2 := func(data int) {
		fmt.Println("hello", data)
	}
	// 使用f()调用
	f2(200)

	// 使用匿名函数打印切片内容
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Printf("%d ", v)
	})
	fmt.Println()

	//可变参数
	fmt.Println(multiSum(1, 2, 3, 4, 5, 6))

	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	MyPrintf(v1, v2, v3, v4)

	// 输入3个字符串, 将它们连成一个字符串
	fmt.Println(joinStrings("pig ", "and", " rat"))
	fmt.Println(joinStrings("hammer", " mom", " and", " hawk"))

	//声明一个calculation类型的变量c
	var c calculation
	//将add函数赋值给c
	c = add
	//调用c
	fmt.Printf("%T %d \n", c, c(1, 2))
}

/*
	x 和 y 是形参名，3 和 4 是调用时的传入的实数，函数返回了一个 float64 类型的值.
	返回值也可以像形式参数一样被命名，在这种情况下，每个返回值被声明成一个局部变量，并根据该返回值的类型，将其初始化为 0。
	如果给返回参数命名了之后，可以在函数内部给对应名称的返回参数赋值之后，直接return就可以，当然写 retrun w,z 也是ok的

	当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。

	func someFunc(x string) []int {
		if x == "" {
			return nil // 没必要返回[]int{}
		}
		...
	}
*/
func hypot(x, y float64) (w, z float64) {
	w = math.Sqrt(x*x + y*y)
	z = math.Sqrt(x*x + y*y)
	return
}

func hypot2(x, y float64) (float64, float64) {
	return math.Sqrt(x*x + y*y), math.Sqrt(x*x + y*y)
}

func fire() {
	fmt.Println("fire")
}

// 遍历切片的每个元素, 通过给定函数进行元素访问
// 这里这个传入func类似于java中传入的接口，然后调用接口的方法
// 其他人调用的时候传入这个接口的匿名实现即可
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

/*
	可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识。

	注意：可变参数通常要作为函数的最后一个参数。

	它是一个语法糖（syntactic sugar），即这种语法对语言的功能并没有影响，但是更方便程序员使用，通常来说，使用语法糖能够增加程序的可读性，从而减少程序出错的可能。

	类型...type本质上是一个数组切片，也就是[]type

	任意类型的可变参数
	之前的例子中将可变参数类型约束为 int，如果你希望传任意类型，可以指定类型为 interface{}，下面是Go语言标准库中 fmt.Printf() 的函数原型：
	func Printf(format string, args ...interface{}) {
		// ...
	}
*/
func multiSum(y int, x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum + y
}

/*
	用 interface{} 传递任意类型数据是Go语言的惯例用法，使用 interface{} 仍然是类型安全的
*/
func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

// 定义一个函数, 参数数量为0~n, 类型约束为字符串
func joinStrings(slist ...string) string {
	// 定义一个字节缓冲, 快速地连接字符串,bytes.Buffer 在这个例子中的作用类似于 StringBuilder
	var b bytes.Buffer
	// 遍历可变参数列表slist, 类型为[]string
	for _, s := range slist {
		// 将遍历出的字符串连续写入字节数组,将每一个传入参数放到 bytes.Buffer 中
		b.WriteString(s)
	}
	// 将连接好的字节数组转换为字符串并输出
	return b.String()
}

/*
	闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使已经离开了自由变量的环境也不会被释放或者删除，在闭包中可以继续使用这个自由变量
	函数 + 引用环境 = 闭包

	一个函数类型就像结构体一样，可以被实例化，函数本身不存储任何信息，只有与引用环境结合后形成的闭包才具有“记忆性”，函数是编译期静态的概念，而闭包是运行期动态的概念。

	闭包（Closure）在某些编程语言中也被称为 Lambda 表达式。

	闭包对环境中变量的引用过程也可以被称为“捕获”

	被捕获到闭包中的变量让闭包本身拥有了记忆效应，闭包中的逻辑可以修改闭包捕获的变量，变量会跟随闭包生命期一直存在，闭包本身就如同变量一样拥有了记忆效应。

	对闭包的简单理解 ： 定义在一个函数内部的函数，闭包就是将函数内部和函数外部连接起来的一座桥梁。

	使用闭包的注意点

		1）由于闭包会使得函数中的变量都被保存在内存中，内存消耗很大，所以不能滥用闭包，否则会造成网页的性能问题，在IE中可能导致内存泄露。解决方法是，在退出函数之前，将不使用的局部变量全部删除。

		2）闭包会在父函数外部，改变父函数内部变量的值。所以，如果你把父函数当作对象（object）使用，把闭包当作它的公用方法（Public Method），把内部变量当作它的私有属性（private value），这时一定要小心，不要随便改变父函数内部变量的值。
*/
func testClosure() {
	adder := Accumulate(1)
	fmt.Printf("adder : %d\n", adder())
	fmt.Printf("adder : %d\n", adder())

}

// 提供一个值, 每次调用函数会指定对值进行累加
func Accumulate(value int) func() int {
	// 返回一个闭包
	return func() int {
		// 累加
		value++
		// 返回一个累加值
		return value
	}
}

/*
	定义函数类型
	我们可以使用type关键字来定义一个函数类型，具体格式如下：

	type calculation func(int, int) int
	上面语句定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。

	凡是满足这个条件的函数都是calculation类型的函数，例如下面的add和sub是calculation类型。
*/
type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

/*
	Go语言的 defer 语句会将其后面跟随的语句进行延迟处理，在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 的逆序进行执行，也就是说，先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行。
	关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源，典型的例子就是对一个互斥解锁，或者关闭一个文件。
	典型的栈式调用，后进先出
	延迟调用是在 defer 所在函数结束时进行，函数结束可以是正常返回时，也可以是发生宕机时。

	在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。
*/
var (
	// 一个演示用的映射
	valueByKey = make(map[string]int)
	// 保证使用映射时的并发安全的互斥锁
	valueByKeyGuard sync.Mutex
)

func testDefer() {
	fmt.Println("defer begin")
	// 将defer放入延迟调用栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	// 最后一个放入, 位于栈顶, 最先调用
	defer fmt.Println(3)

	valueByKeyGuard.Lock()

	// defer后面的语句不会马上调用, 而是延迟到函数结束时调用
	defer valueByKeyGuard.Unlock()

	fmt.Println("defer end")
	//一些defer相关的练习
	testDeferExercise()
}

/*
	Go语言的错误处理思想及设计包含以下特征：
	一个可能造成错误的函数，需要返回值中返回一个错误接口（error），如果调用是成功的，错误接口将返回 nil，否则返回错误。
	在函数调用后需要检查错误，如果发生错误，则进行必要的错误处理。

	Go语言希望开发者将错误处理视为正常开发必须实现的环节，正确地处理每一个可能发生错误的函数，同时，Go语言使用返回值返回错误的机制，也能大幅降低编译器、运行时处理错误的复杂度，让开发者真正地掌握错误的处理。

	错误接口的定义格式
	error 是 Go 系统声明的接口类型，代码如下：
		type error interface {
			Error() string
		}
	所有符合 Error()string 格式的方法，都能实现错误接口，Error() 方法返回错误的具体描述，使用者可以通过这个字符串知道发生了什么错误。
*/
func testError() {

}
