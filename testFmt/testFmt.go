package main

import (
	"errors"
	"fmt"
	"os"
)

/*

%a  浮点数、十六进制数字和p-记数法（c99
%A  浮点数、十六进制数字和p-记法（c99）
%c  一个字符(char)
%C  一个ISO宽字符
%d  有符号十进制整数(int)（%ld、%Ld：长整型数据(long),%hd：输出短整形。） 字符串和字符会输出表示该字符的整数
%e  浮点数、e-记数法
%E  浮点数、E-记数法
%f  单精度浮点数(默认float)、十进制记数法（%.nf  这里n表示精确到小数位后n位.十进制计数）
%g  根据数值不同自动选择%f或%e．
%G  根据数值不同自动选择%f或%e.
%i  有符号十进制数（与%d相同）
%o  无符号八进制整数
%p  指针,打印指针的地址，十六进制方式显示
%s  对应字符串char*（%s = %hs = %hS 输出 窄字符）
%S  对应宽字符串WCAHR*（%ws = %S 输出宽字符串）
%u  无符号十进制整数(unsigned int)
%x  使用十六进制数字0xf的无符号十六进制整数
%X  使用十六进制数字0xf的无符号十六进制整数
%%  打印一个百分号
%v	按值的本来值输出，字符串和字符会输出表示该字符的整数
%+v	在 %v 基础上，对结构体字段名和值进行展开
%#v	输出 Go 语言语法格式的值
%T 输出 Go 语言语法格式的类型和值 string int 之类的


%I64d 用于INT64 或者 long long
%I64u 用于UINT64 或者 unsigned long long
%I64x 用于64位16进制数据
*/
func main() {
	fmt.Printf("%v %v %v \n", "月球基地", 3.14, true)
	fmt.Printf("%T %T %T \n", "月球基地", 3.14, true)
	testFprint()
	testErrorf()
}

/*
	Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容
	func Fprint(w io.Writer, a ...interface{}) (n int, err error)
	func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	只要满足io.Writer接口的类型都支持写入。
*/
func testFprint() {
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("/Users/kuroneko/logs/xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "沙河小王子"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
}

/*
	Sprint系列函数会把传入的数据生成并返回一个字符串。

	func Sprint(a ...interface{}) string
	func Sprintf(format string, a ...interface{}) string
	func Sprintln(a ...interface{}) string
*/
func testSprint() {
	s1 := fmt.Sprint("沙河小王子")
	name := "沙河小王子"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("沙河小王子")
	fmt.Println(s1, s2, s3)
}

/*
Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。
func Errorf(format string, a ...interface{}) error

*/
func testErrorf() {
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误: %w", e)
	fmt.Println(w)
	fmt.Println(e)
}
