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
