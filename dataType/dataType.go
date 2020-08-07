package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	fmt.Println("-----------------testInt----------------")
	testInt()
	fmt.Println("-----------------testFloat----------------")
	testFloat()
	fmt.Println("-----------------testComplex----------------")
	testComplex()
	fmt.Println("-----------------testBool----------------")
	testBool()
	fmt.Println("-----------------testString----------------")
	testString()
	fmt.Println("-----------------testByteRune----------------")
	testByteRune()
	fmt.Println("-----------------typeChange----------------")
	typeChange()
	fmt.Println("-----------------testAlias----------------")
	testAlias()
	fmt.Println("-----------------testNil----------------")
	testNil()
}

/*

	在Go语言中，布尔类型的零值（初始值）为 false，数值类型的零值为 0，字符串类型的零值为空字符串""，而指针、切片、映射、通道、函数和接口的零值则是 nil。
	nil 标识符是不能比较的
	fmt.Println(nil==nil) //编译会报错
	nil 没有类型
	fmt.Printf("%T", nil) //运行报错
    print(nil) //运行报错
*/
func testNil() {
	//nil 没有默认类型
	fmt.Printf("%T %v \n", nil, nil) // 输出 <nil> <nil>

	//不同类型的nil指针的地址是一样的
	var arr []int
	var num *int
	fmt.Printf("%p %T \n", arr, arr) // 输出 0x0 []int
	fmt.Printf("%p %T \n", num, num) // 输出 0x0 *int

	var s1 []int
	var s2 []int
	//两个未初始化变量的值无法比较，即使他们都是nil
	// fmt.Printf(s1 == s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 不同类型的 nil 值占用的内存大小是不一样的
	var p *struct{}
	fmt.Println(unsafe.Sizeof(p)) // 8
	var s []int
	fmt.Println(unsafe.Sizeof(s)) // 24
	var m map[int]bool
	fmt.Println(unsafe.Sizeof(m)) // 8
	var c chan string
	fmt.Println(unsafe.Sizeof(c)) // 8
	var f func()
	fmt.Println(unsafe.Sizeof(f)) // 8
	var i interface{}
	fmt.Println(unsafe.Sizeof(i)) // 16
}

/**

	获取对象的长度的内建len()函数返回的长度可以根据不同平台的字节长度进行变化。
	实际使用中，切片或 map 的元素数量等都可以用int来表示。
	在涉及到二进制传输、读写文件的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用int和 uint。
	uint8	无符号 8位整型 (0 到 255)  uint8就是我们熟知的byte型
	uint16	无符号 16位整型 (0 到 65535)
	uint32	无符号 32位整型 (0 到 4294967295)
	uint64	无符号 64位整型 (0 到 18446744073709551615)
	int8	有符号 8位整型 (-128 到 127)
	int16	有符号 16位整型 (-32768 到 32767)
	int32	有符号 32位整型 (-2147483648 到 2147483647)
	int64	有符号 64位整型 (-9223372036854775808 到 9223372036854775807)

	uint	32位操作系统上就是uint32，64位操作系统上就是uint64
	int	32位操作系统上就是int32，64位操作系统上就是int64
	uintptr	无符号整型，用于存放一个指针

	Go1.13版本之后引入了数字字面量语法，这样便于开发者以二进制、八进制或十六进制浮点数的格式定义数字，例如：

	v := 0b00101101， 代表二进制的 101101，相当于十进制的 45。 v := 0o377，代表八进制的 377，相当于十进制的 255。 v := 0x1p-2，代表十六进制的 1 除以 2²，也就是 0.25。 而且还允许我们用 _ 来分隔数字，比如说：

	v := 123_456 等于 123456。
**/
func testInt() {

	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	var d int64 = 9223372036854
	fmt.Printf("%d \n", d) // 9223372036854
	fmt.Printf("%b \n", d) // 10000110001101111011110100000101101011110110

}

/**
	Go语言支持两种浮点型数：float32和float64。这两种浮点型数据格式遵循IEEE 754标准： float32 的浮点数的最大范围约为 3.4e38，可以使用常量定义：math.MaxFloat32。 float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64。

	打印浮点数时，可以使用fmt包配合动词%f
**/
func testFloat() {

	var pi float32 = math.Pi
	fmt.Printf(" pi = %f\n", pi)
	//打印小数点后两位
	fmt.Printf(" pi = %.2f\n", pi)

}

/**

复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。

*/
func testComplex() {

	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

}

/**
布尔类型变量的默认值为false。
Go 语言中不允许将整型强制转换为布尔型.
布尔型无法参与数值运算，也无法与其他类型进行转换。
*/
func testBool() {
	b := false
	var c bool = true
	fmt.Println(b)
	fmt.Println(c)

}

/*
Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。
Go 语言里的字符串的内部实现使用UTF-8编码。
字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符
*/
func testString() {

	s1 := "hello"
	var s2 string = "world"
	fmt.Println(s1 + s2)
	//多行字符串,使用反引号字符
	//反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效
	//在`间的所有代码均不会被编译器识别，而只是作为字符串的一部分。
	/*
			输出：
			第一行 \n
		        第二行 \s \t
		        第三行
	*/
	s3 := `第一行 \n
	第二行 \s \t
	第三行`
	fmt.Println(s3)

	/*

		+或fmt.Sprintf	拼接字符串
		strings.Split	分割
		strings.contains	判断是否包含
		strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
		strings.Index(),strings.LastIndex()	子串出现的位置
		strings.Join(a[]string, sep string)	join操作

		len(str)	字符串的字节(byte)长度
		utf8.RuneCountInString() 函数，统计 Uncode 字符数量。
		ASCII 字符串长度使用 len() 函数。
		Unicode 字符串长度使用 utf8.RuneCountInString() 函数。
	*/

	var listStr = "a,b,c"
	var a = strings.Split(listStr, ",")
	fmt.Println(a)

	/*
		字符串是不能修改的和 java中是final
		要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。
		无论哪种转换，都会重新分配内存，并复制字节数组。
	*/
	str1 := "big"
	// 强制类型转换
	byteStr1 := []byte(str1)
	byteStr1[0] = 'p'
	fmt.Println(string(byteStr1))

	str2 := "白萝卜"
	// 强制类型转换
	runeStr2 := []rune(str2)
	runeStr2[0] = '红'
	fmt.Println(string(runeStr2))

	fmt.Printf("白萝卜's size = %d \n", utf8.RuneCountInString(str2))

}

/*
	组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来
	当需要处理中文、日文或者其他复合字符时，则需要用到rune类型
	1.uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
	2.rune类型，代表一个 Unicode（UTF-8 8-bit Unicode Transformation Format）字符 rune类型实际是一个int32  Unicode 至少占用 2 个字节
	在书写 Unicode 字符时，需要在 16 进制数之前加上前缀\u或者\U。因为 Unicode 至少占用 2 个字节，所以我们使用 int16 或者 int 类型来表示。如果需要使用到 4 字节，则使用\u前缀，如果需要使用到 8 个字节，则使用\U前缀。

	go 默认使用int32来保存字符


	Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾。

	字符串底层是一个byte数组，所以可以和[]byte类型相互转换。
	字符串是不能修改的
	字符串是由byte字节组成，所以字符串的长度是byte字节的长度。
	rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。

	Unicode 包中内置了一些用于测试字符的函数，这些函数的返回值都是一个布尔值，如下所示（其中 ch 代表字符）：
	判断是否为字母：unicode.IsLetter(ch)
	判断是否为数字：unicode.IsDigit(ch)
	判断是否为空白符号：unicode.IsSpace(ch)


*/
func testByteRune() {
	var a = '中'
	var b = 'x'
	fmt.Println(a + b)
	s := "hello世界"
	fmt.Printf("s.len = %d \n", len(s))
	fmt.Println("print byte")
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println("\n print rune")
	//因为UTF8编码下一个中文汉字由3~4个字节(byte)组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串
	//todo
	for index, r := range s { //rune
		fmt.Printf("[%T][%d]%v(%c)\t", r, index, r, r)
	}
	fmt.Println()

	var w = 'w'
	fmt.Printf("%T ", w)
	fmt.Printf("%v ", w)
	fmt.Printf("%c \n", w)

	var r rune = '界'
	fmt.Printf("%T ", r)
	fmt.Printf("%v ", r)
	fmt.Printf("%c \n", r)

	var r2 = '界'
	fmt.Printf("%T ", r2)
	fmt.Printf("%v ", r2)
	fmt.Printf("%c \n", r2)
	fmt.Println()
}

/*
	Go语言中只有强制类型转换，没有隐式类型转换。

	该语法只能在两个类型之间支持相互转换的时候使用。

	强制类型转换的基本语法如下：

	T(表达式)

	其中，T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等.

	比如计算直角三角形的斜边长时使用math包的Sqrt()函数，该函数接收的是float64类型的参数，而变量a和b都是int类型的，这个时候就需要将a和b强制类型转换为float64类型。
*/
func typeChange() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	// Sqrt() : 求一个数的平方根
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

/*
	1.定义类型别名的写法为：
	type TypeAlias = Type
	TypeAlias 只是 Type 的别名，本质上 TypeAlias 与 Type 是同一个类型
	// 将int取一个别名叫IntAlias
	type IntAlias = int

	2.创建一个新类型
	// 将NewInt定义为int类型
	type NewInt int


*/
func testAlias() {
	// 将NewInt定义为int类型
	type NewInt int
	// 将int取一个别名叫IntAlias
	type IntAlias = int
	// 将a声明为NewInt类型
	var a NewInt
	// 查看a的类型名  输出 main.NewInt
	fmt.Printf("a type: %T\n", a)
	// 将a2声明为IntAlias类型
	var a2 IntAlias
	// 查看a2的类型名 处处int
	fmt.Printf("a2 type: %T\n", a2)
}
