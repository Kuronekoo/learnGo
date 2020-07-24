package main

import (
	"fmt"
)

/*
go 语言中声明的变量必须被使用，否则编译不通过
*/
// 全局变量m
// 函数外的每个语句都必须以关键字开始（var、const、func等）
var m = 100

//全局常量
//常量在定义的时候必须赋值。
//常量定义了之后不使用不会报错
const pi = 3.14

//一次性定义多个常量
//const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。 e2的值也是2.7182
const (
	e = 2.7182
	e2
)

func foo() (int, string) {
	return 10, "let's go"
}

func main() {
	testVar()
	testIota()
}

func testIota() {
	fmt.Println("-----------------iota----------------")

	//iota是go语言的常量计数器，只能在常量的表达式中使用。
	//iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。
	//如果在iota的关键前加了计算公示f(x) ，那么在下一个iota或者const出现之前，都会以这个公式的值输出,f(x)
	const (
		n1 = 3 * iota //0 0*3
		n2            //1 1*3
		_             //跳过2 2*3
		n4            //3 3*3
		n6 = iota     // 4
		n7 = 2 * iota // 5 5*2
	)

	const n10 = iota //0

	const (
		mutexLocked       = 1 << iota //0 1 << 0
		mutexWoken                    //1 1 << 1 = 10    2
		mutexStarving                 //2 1 << 2 = 100   4
		mutexStarving3                //3 1 << 3 = 100   8
		mutexWaiterShift  = iota      //4
		mutexWaiterShift0 = 10 * iota //5 10*5 = 50
		mutexWaiterShift1             //6 10*6 = 60
	)

	fmt.Printf(" n1 = %d \n", n1)
	fmt.Printf(" n2 = %d \n", n2)
	fmt.Printf(" n4 = %d \n", n4)
	fmt.Printf(" n5 = %d \n", n6)
	fmt.Printf(" n5 = %d \n", n7)
	fmt.Printf(" n10 = %d \n", n10)
	fmt.Printf(" mutexLocked = %d \n", mutexLocked)
	fmt.Printf(" mutexWoken = %d \n", mutexWoken)
	fmt.Printf(" mutexStarving = %d \n", mutexStarving)
	fmt.Printf(" mutexStarving3 = %d \n", mutexStarving3)
	fmt.Printf(" mutexWaiterShift = %d \n", mutexWaiterShift)
	fmt.Printf(" mutexWaiterShift0 = %d \n", mutexWaiterShift0)
	fmt.Printf(" mutexWaiterShift1 = %d \n", mutexWaiterShift1)
	fmt.Printf(" 1 << 4 = %d \n", 1<<4)
	fmt.Println("-----------------iota----------------")
}
func testVar() {
	fmt.Println("-----------------var----------------")

	//var 变量名 变量类型
	//go是静态语言，变量必须要声明类型
	var name string
	var age int
	var isOk bool
	//批量声明变量
	var (
		a string
		b int
		c bool
		d float32
	)
	//var 变量名 类型 = 表达式
	var username string = "vars"
	//类型推导
	//一次性初始化多个变量
	var password, count = "pass", 20

	//在函数内部，可以使用更简略的 := 方式声明并初始化变量。只能用在局部变量
	n := "it's n"
	m := 200

	//匿名变量
	//_用于占位，表示忽略值。
	//函数有两个返回值，但是只想接收其中一个
	//匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。
	var x, _ = foo()

	fmt.Println(x)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(name)
	fmt.Println(username)
	fmt.Println(age)
	fmt.Println(isOk)
	fmt.Println(password)
	fmt.Println(count)
	fmt.Println(n, m)

	fmt.Println("-----------------var----------------")

}
