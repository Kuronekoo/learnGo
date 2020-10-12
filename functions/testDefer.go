package main

import "fmt"

/*
	在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。
*/
func testDeferExercise() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	f5()
}

//5
func f1() int {
	x := 5
	defer func() {
		x++
		fmt.Printf("x ----> %d\n", x)
	}()
	return x // 先做x=5,然后执行到return ,意思是赋值完直接retrun了？
}

//6
func f2() (x int) {
	defer func() {
		x++
		fmt.Printf("x ----> %d \n", x)

	}()
	return 5 //// 直接声明返回x,最后return 5,是把5赋值给x,然后x++
}

// f2相当于
// func f2() (x int) {
// 	x = 5
// 	defer func() {
// 		x++
// 	}()
// 	return //// 直接声明返回x,最后return 5,是把5赋值给x,然后x++
// }

//5
func f3() (y int) {
	x := 5
	defer func() {
		x++ //// 修改的是x
	}()
	return x // 返回值 = y = x = 5
}

//f3相当于
// func f3() (y int) {
// 	x := 5
// 	y = x
// 	defer func() {
// 		x++ //// 修改的是x
// 	}()
// 	return // 返回值 = y = x = 5
// }

//5
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

// f4相当于
// func f4() (x int) {
// 	x = 5
// 	defer func(x int) {
// 		x++
// 	}(x)
// 	return
// }

func f5() {
	ok := true
	if ok {
		println("ok...")
		return
	}
	//defer代码不会执行
	defer println("defer exec...")

}
