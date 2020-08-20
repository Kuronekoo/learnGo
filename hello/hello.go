//每一个go文件都需要归属于一个包
package main

//引入fmt包，注意双引号
import "fmt"

//func关键字定义方法，main是主函数，是程序启动的入口，没有返回值
func main() {

	//在屏幕上打印hello,world!
	fmt.Println("hello,world!" + " yes!")
	fmt.Println("")

}

/**
// go build hello.go 命令对该 go 文件进行编译，windows下生成 .exe 文件. uninx环境直接生成一个可执行文件
// 新版本go需要在项目的根目录下创建一个go.mod之后 可以直接在项目目录下go build ，生成的可执行文件默认是为项目名称
// ./hello 执行文件
// 或者通过 go run hello.go 直接运行源码
同一个package下有多个go文件，直接运行main方法所在的go文件，并调用其他的go文件，会报undefined错误
原因是该包下的其他go文件没有一起编译，以下方法可破
1.go run *.go
2.go build .



// 项目根目录下 go install ，编译build之后将项目拷贝到$GOPATH/bin/下，然后就可以在任意地方直接执行可执行文件了。
// 跨平台编译


Mac 下编译 Linux 和 Windows平台 64位 可执行程序：

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build


Linux 下编译 Mac 和 Windows 平台64位可执行程序：

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build


Windows下编译Mac平台64位可执行程序：

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build


**/

/*
	包（package）是多个Go源码的集合，是一种高级的代码复用方案，Go语言为我们提供了很多内置包，如fmt、os、io等
	我们还可以根据自己的需要创建自己的包。一个包可以简单理解为一个存放.go文件的文件夹。 该文件夹下面的所有go文件都要在代码的第一行添加如下代码，声明该文件归属的包。
		package 包名
	注意事项：
	一个文件夹下面直接包含的文件只能归属一个package，同样一个package的文件不能在多个文件夹下。
	包名可以不和文件夹的名字一样，包名不能包含 - 符号。
	包名为main的包为应用程序的入口包，这种包编译后会得到一个可执行文件，而编译不包含main包的源代码则不会得到可执行文件。

	如果想在一个包中引用另外一个包里的标识符（如变量、常量、类型、函数等）时，该标识符必须是对外可见的（public）。在Go语言中只需要将标识符的首字母大写就可以让标识符对外可见了。
	结构体中的字段名和接口中的方法名如果首字母都是大写，外部包可以访问这些字段和方法

	要在代码中引用其他包的内容，需要使用import关键字导入使用的包。具体语法如下:

	import "包的路径"
	注意事项：

	import导入语句通常放在文件开头包声明语句的下面。
	导入的包名需要使用双引号包裹起来。
	包名是从$GOPATH/src/后开始计算的，使用/进行路径分隔。
	Go语言中禁止循环导入包。

	在导入包名的时候，我们还可以为导入的包设置别名。通常用于导入的包名太长或者导入的包名冲突的情况。具体语法格式如下：

	import 别名 "包的路径"
	import (
    "fmt"
    m "github.com/Q1mi/studygo/pkg_test"
	)

	func main() {
		fmt.Println(m.Add(100, 200))
		fmt.Println(m.Mode)
	}

	init()函数介绍
	在Go语言程序执行时导入包语句会自动触发包内部init()函数的调用。需要注意的是： init()函数没有参数也没有返回值。 init()函数在程序运行时自动被调用执行，不能在代码中主动调用它。
	加载顺序 ：  全局声明-->init()-->main()






*/
