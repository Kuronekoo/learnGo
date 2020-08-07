package main

import "fmt"

/*

 */
func main() {
	fmt.Println("--------------testIf--------------")
	testIf(3)
	fmt.Println("--------------testFor--------------")
	testFor()
	fmt.Println("--------------testSwitch--------------")
	testSwitch()
	fmt.Println("--------------testGoto--------------")
	testGoto()
	fmt.Println("--------------testBreack--------------")
	testBreack()
	fmt.Println("--------------testContinue--------------")
	testContinue()
}

func testIf(i int) {
	if 1 == i {
		fmt.Println("11111")
	} else if 2 == i {
		fmt.Println("22222")
	} else if i-3 == 0 {
		fmt.Println("oooooo")
	} else {

	}
	// Connect 是一个带有返回值的函数，err:=Connect() 是一个语句，执行 Connect 后，将错误保存到 err 变量中。
	// err != nil 才是 if 的判断表达式，当 err 不为空时，打印错误并返回。
	// 这种写法可以将返回值与判断放在一行进行处理，而且返回值的作用范围被限制在 if、else 语句组合中。
	// if err := Connect(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

/*
	for range 结构是Go语言特有的一种的迭代结构，在许多情况下都非常有用，for range 可以遍历数组、切片、字符串、map 及通道（channel），for range 语法上类似于其它语言中的 foreach 语句，一般形式为：
	for key, val := range coll {
		...
	}

	val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值。一个字符串是 Unicode 编码的字符（或称之为 rune ）集合，因此也可以用它来迭代字符串：
	range遍历中每次返回的变量都是每个元素的值的拷贝，但是用的总是同一个内存地址，因此如果需要取内存地址等操作就可能在这里发生意想不到的问题
	for index, value := range str {
		...
	}

	for index, value := range str {
		//这样写的话最后只会所有可以都指向最后一个地址，值也是最后一个地址
		//map[index]=&value
		//新生成一块内存地址就可以避免这个问题
		newVavalue := value;
		map[index]=&newVavalue
	}

	通过 for range 遍历的返回值有一定的规律：
	数组、切片、字符串返回索引和值。
	map 返回键和值。
	通道（channel）只返回通道内的值
*/
func testFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("sum : %v \n", sum)
	//无限循环
	for {
		if sum > 100 {
			break
		}
		sum++
	}
	fmt.Printf("sum : %v \n", sum)
	//for循环的第一个赋值语句可以不写
	step := 2
	for ; step > 0; step-- {
		fmt.Println(step)
	}
	//忽略赋值和判断语句
	var i int
	for ; ; i++ {
		if i > 10 {
			break
		}
	}
	//只有一个条件判断的for循环
	//看起来就是while循环
	var j int
	for j <= 10 {
		j++
	}
	//遍历数组、切片——获得索引和值
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d  value:%d\n", key, value)
	}
	//遍历字符串——获得字符
	var str = "hello 你好"
	for key, value := range str {
		fmt.Printf("key:%d value:0x%x\n", key, value)
	}

	//遍历 map——获得 map 的键和值
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
	//遍历通道（channel）——接收通道数据
	//第 1 行创建了一个整型类型的通道。
	c := make(chan int)
	//第 3 行启动了一个 goroutine，其逻辑的实现体现在第 5～8 行，实现功能是往通道中推送数据 1、2、3，然后结束并关闭通道。
	//这段 goroutine 在声明结束后，在第 9 行马上被执行。
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	//第 11 行开始，使用 for range 对通道 c 进行遍历，其实就是不断地从通道中取数据，直到通道被关闭。
	for v := range c {
		fmt.Println(v)
	}
}

/*
	Go语言的 switch 要比C语言的更加通用，表达式不需要为常量，甚至不需要为整数，case 按照从上到下的顺序进行求值，直到找到匹配的项，如果 switch 没有表达式，则对 true 进行匹配，因此，可以将 if else-if else 改写成一个 switch。
*/
func testSwitch() {
	//case 与 case 之间是独立的代码块，不需要通过 break 语句跳出当前 case 代码块以避免执行到下一行
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}
	//一分支多值
	var b = "mum"
	switch b {
	case "mum", "daddy":
		fmt.Println("family")
	}
	//case 后不仅仅只是常量，还可以和 if 一样添加表达式
	var r int = 11
	switch {
	case r > 10 && r < 20:
		fmt.Println(r)
	}

	//在Go语言中 case 是一个独立的代码块，执行完毕后不会像C语言那样紧接着执行下一个 case，但是为了兼容一些移植代码，依然加入了 fallthrough 关键字来实现这一功能
	var s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")
	}
}

/*
	Go语言中 goto 语句通过标签进行代码间的无条件跳转，同时 goto 语句在快速跳出循环、避免重复退出上也有一定的帮助，使用 goto 语句能简化一些代码的实现过程。
*/
func testGoto() {
	simpleGoto()
	gotoDealWithError()
}

func simpleGoto() {
	for x := 0; x < 2; x++ {
		for y := 0; y < 10; y++ {
			fmt.Printf("x = %d,y = %d \n", x, y)
			if y == 2 {
				// 跳转到标签
				//标签只能被 goto 使用，但不影响代码执行流程，此处如果不手动返回，在不满足条件时，也会执行第标签后的代码
				goto breakHere
			}
		}
	}
	// 手动返回, 避免执行进入标签
	// 如果这里没有return的话，会执行标签之后的代码
	return
	// 标签
	//使用 goto 语句跳转到指明的标签处，标签在第 23 行定义。
breakHere:
	fmt.Println("done")
}

func gotoDealWithError() {
	fmt.Println("gotoDealWithError")
	// 	err := firstCheckError()
	// 	if err != nil {
	//发生错误时，跳转错误标签 onExit。
	// 		goto onExit
	// 	}
	// 	err = secondCheckError()
	// 	if err != nil {
	// 		goto onExit
	// 	}
	// 	fmt.Println("done")
	// 	return
	// 汇总所有流程进行错误打印并退出进程。
	// onExit:
	// 	fmt.Println(err)
	// 	exitProcess()
}

/*
	Go语言中 break 语句可以结束 for、switch 和 select 的代码块，另外 break 语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的 for、switch 和 select 的代码块上。
*/
func testBreack() {
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("i = %d,j = %d \n", i, j)
			switch j {
			case 2:
				fmt.Println(i, j)
				break OuterLoop
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
		}
	}
	fmt.Println("loop exit!")
}

/*
	Go语言中 continue 语句可以结束当前循环，开始下一次的循环迭代过程，仅限在 for 循环内使用，在 continue 语句后添加标签时，表示开始标签对应的循环
*/
func testContinue() {
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				//结束当前循环，开启下一次的外层循环，而不是内循环
				continue OuterLoop
			}
		}
	}
}
