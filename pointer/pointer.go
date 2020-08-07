package main

import (
	"flag"
	"fmt"
)

/*
	指针（pointer）在Go语言中可以被拆分为两个核心概念：
	1.类型指针，允许对这个指针类型的数据进行修改，传递数据可以直接使用指针，而无须拷贝数据，类型指针不能进行偏移和运算。
	2.切片，由指向起始元素的原始指针、元素数量和容量组成。

	受益于这样的约束和拆分，Go语言的指针类型变量即拥有指针高效访问的特点，又不会发生指针偏移，从而避免了非法修改关键性数据的问题。同时，垃圾回收也比较容易对不会发生偏移的指针进行检索和回收。
	切片比原始指针具备更强大的特性，而且更为安全。切片在发生越界时，运行时会报出宕机，并打出堆栈，而原始指针只会崩溃。

	指针是 C/C++ 语言拥有极高性能的根本所在，在操作大块数据和做偏移时即方便又便捷。

	new与make的区别
		二者都是用来做内存分配的。
		make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
		而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。


*/
func main() {
	testPoint()
	newPoint()
	fucFlag()
}

/*
	一个指针变量可以指向任何一个值的内存地址，它所指向的值的内存地址在 32 和 64 位机器上分别占用 4 或 8 个字节，占用字节的大小与所指向的值的大小无关。
	当一个指针被定义后没有分配到任何变量时，它的默认值为 nil。指针变量通常缩写为 ptr。

	每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go语言中使用在变量名前面添加&操作符（前缀）来获取变量的内存地址（取地址操作），格式如下：
	ptr := &v    // v 的类型为 T

	变量、指针和地址三者的关系是，每个变量都拥有地址，指针的值就是地址。

	当使用&操作符对普通变量进行取地址操作并得到变量的指针后，可以对指针使用*操作符，也就是指针取值

	取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。

*/
func testPoint() {
	var cat int = 1
	var str string = "banana"
	fmt.Printf("%p %p \n", &cat, &str)

	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"
	// 对字符串取地址, 字符串ptr类型为*string 整型ptr类型为*int
	ptr := &house
	prtInt := &cat
	// 打印ptr的类型
	fmt.Printf("prtInt type: %T\n", prtInt)
	// 打印ptr的类型，其类型为 *string。
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr的指针地址，地址每次重新运行都会发生变化。
	fmt.Printf("ptr address: %p\n", ptr)
	// 对指针进行取值操作,取出指针指向的值，变量 value 的类型为 string。
	value := *ptr
	// 取值后的类型，打印取值后 value 的类型。
	fmt.Printf("prt value type: %T\n", value)
	// 指针取值后就是指向变量的值，打印 value 的值。
	fmt.Printf("ptr value: %s\n", value)
}

func testPointModify() {
	// 准备两个变量, 赋值1和2
	x, y := 1, 2
	// 交换变量值
	swap(&x, &y)
	// 输出变量值
	fmt.Println(x, y)
}

// 交换函数 参数为 a、b，类型都为 *int 指针类型。
// *操作符作为右值时，意义是取指针的值，作为左值时，也就是放在赋值操作符的左边时，表示 a 指针指向的变量。
// 其实归纳起来，*操作符的根本意义就是操作指针指向的变量。当操作在右值时，就是取指向变量的值，当操作在左值时，就是将值设置给指向的变量。
func swap(a, b *int) {
	// 取a指针的值, 赋给临时变量t ， t 此时是 int 类型。
	t := *a
	// 取b指针的值, 赋给a指针指向的变量  此时*a的意思不是取 a 指针的值，而是“a 指向的变量”。
	*a = *b
	// 将 t 的值赋给指针 b 指向的变量。
	*b = t
}

// 交换的是 a 和 b 的地址，在交换完毕后，a 和 b 的变量值确实被交换。
// 但是a b变量的值没有变化 ， a指针还是指向a对象，内存地址变了 b指针还是指向b对象，内存地址变了
func swapAddress(a, b *int) {
	b, a = a, b
}

/*
	new() 函数可以创建一个对应类型的指针，创建过程会分配内存，被创建的指针指向默认值。

	new是一个内置的函数，它的函数签名如下：
	func new(Type) *Type

	Type表示类型，new函数只接受一个参数，这个参数是一个类型
	*Type表示类型指针，new函数返回一个指向该类型内存地址的指针

*/
func newPoint() {
	str := new(string)
	fmt.Printf("str default type = %T\n", str)
	fmt.Printf("*str default value = %v\n", *str)
	*str = "Golanguage"
	fmt.Printf("str type = %T\n", str)
	fmt.Printf("str value = %v\n", str)
	fmt.Printf("*str type = %T\n", *str)
	fmt.Printf("*str value = %v\n", *str)

}

/*
	flag.String，定义一个 mode 变量，这个变量的类型是 *string
	1.参数名称：在命令行输入参数时，使用这个名称。 即在 go run xxx.go --mode=flag!
	2.参数值的默认值：与 flag 所使用的函数创建变量类型对应，String 对应字符串、Int 对应整型、Bool 对应布尔型等。
	3.参数说明：使用 -help 时，会出现在说明中。

	由于之前已经使用 flag.String 注册了一个名为 mode 的命令行参数，flag 底层知道怎么解析命令行，并且将值赋给 mode*string 指针，在 Parse 调用完毕后，无须从 flag 获取值，而是通过自己注册的这个 mode 指针获取到最终的值。

*/
var mode = flag.String("mode", "", "process mode")

func fucFlag() {
	// 解析命令行参数,并将结果写入到变量 mode 中。
	flag.Parse()
	// 输出命令行参数,打印 mode 指针所指向的变量。
	fmt.Println(*mode)
}
