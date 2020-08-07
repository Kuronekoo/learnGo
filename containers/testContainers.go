package main

import (
	"container/list"
	"fmt"
	"sync"
)

/*
	make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。

	make函数的函数签名如下：

	func make(t Type, size ...IntegerType) Type

*/
func main() {
	fmt.Println("--------------arr--------------")
	testArr()
	fmt.Println("--------------slice--------------")
	testSlice()

	fmt.Println("--------------range--------------")
	testRange()
	fmt.Println("--------------map--------------")
	testMap()
	fmt.Println("--------------syncMap--------------")
	testSyncMap()
	fmt.Println("--------------testList--------------")
	testList()

}

/*
	数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。因为数组的长度是固定的，所以在Go语言中很少直接使用数组。

	和数组对应的类型是 Slice（切片），Slice 是可以增长和收缩的动态序列，功能也更灵活

	数组的声明语法如下：

	var 数组变量名 [元素数量]Type

	数组变量名：数组声明及使用时的变量名。
	元素数量：数组的元素数量，可以是一个表达式，但最终通过编译期计算的结果必须是整型数值，元素数量不能含有到运行时才能确认大小的数值。
	Type：可以是任意基本类型，包括数组本身，类型为数组本身时，可以实现多维数组。

*/
func testArr() {
	var a [3]int             // 定义三个整数的数组
	fmt.Println(a[0])        // 打印第一个元素
	fmt.Println(a[len(a)-1]) // 打印最后一个元素
	// 打印索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// 仅打印元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
	//默认情况下，数组的每个元素都会被初始化为元素类型对应的零值，对于数字类型来说就是 0，同时也可以使用数组字面值语法，用一组值来初始化数组：
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q) // "0"
	fmt.Println(r) // "0"
	// 在数组的定义中，如果在数组长度的位置出现“...”省略号，则表示数组的长度是根据初始化值的个数来计算
	// 数组的长度是数组类型的一个组成部分，因此 [3]int 和 [4]int 是两种不同的数组类型，数组的长度必须是常量表达式，因为数组的长度需要在编译阶段确定。
	dynamic := [...]int{1, 2, 3}
	fmt.Printf("%T\n", dynamic) // "[3]int"

	/*
		如果两个数组类型相同（包括数组的长度，数组中元素的类型）的情况下，我们可以直接通过较运算符（==和!=）来判断两个数组是否相等。
		只有当两个数组的所有元素都是相等的时候数组才是相等的，不能比较两个类型不同的数组，否则程序将无法完成编译。
	*/
	a1 := [2]int{1, 2}
	a2 := [...]int{1, 2}
	a3 := [2]int{1, 3}
	fmt.Println(a1 == a2, a1 == a3, a2 == a3) // "true false false"
	// a4 := [3]int{1, 2}
	// fmt.Println(a1 == a4) // 编译错误：无法比较 [2]int == [3]int

	/*
		遍历数组——访问每一个数组元素
	*/
	var team = [3]string{"hammer", "soldier", "mum"}
	for k, v := range team {
		fmt.Println(k, v)
	}

	// 创建多维数组
	// var array_name [size1][size2]...[sizen] array_type
	// 声明一个 2×2 的二维整型数组
	var array [2][2]int
	// 设置每个元素的整型值
	array[0][0] = 10
	array[0][1] = 20
	array[1][0] = 30
	array[1][1] = 40

}

/*
	切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
	切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
	1.声明切片类型的基本语法如下：
		var name []T
		使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。
	2.切片表达式从字符串、数组、指向数组或切片的指针构造子字符串或切片,按索引从中截取，截取的字符串包含begin，不包含end,相当于截取[begin,end-1]
		可以省略切片表达式中的任何索引。省略了low则默认为0；省略了high则默认为切片操作数的长度:
		切片的容量: cap = len(arr)-begin
		slice := arr[begin,end]
	3.使用make()函数构建新的切片
		make( []Type, size, cap )
		其中 Type 是指切片的元素类型，size 指的是为这个类型分配多少个元素，cap 为预分配的元素数量。
		cap设定后不影响 size，用来提前分配连续的空间，降低多次分配空间造成的性能问题。
		len()函数取的是切片的size
	4.对于数组，指向数组的指针，或切片a(注意不能是字符串)支持完整切片表达式：
		简单切片表达式a[low: high]相同类型、相同长度和元素的切片。另外，它会将得到的结果切片的容量设置为max-low
		满足条件 max>=end>=begin
		a[begin : end : max]

	使用 make() 函数生成的切片一定会发生内存分配操作
	但给定开始与结束位置（包括切片复位）的切片只是将新的切片结构指向已经分配好的内存区域，设定开始与结束位置，不会发生内存分配操作。

	切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。

	要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断

	切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。
	切片唯一合法的比较操作是和nil比较。 一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。
	但是我们不能说一个长度和容量都是0的切片一定是nil


	slice拷贝，这种拷贝实际上是新建另一个s2的slice指针指向s1的头部
	s1 := make([]int, 3) //[0 0 0]
	s2 := s1

	使用copy()函数进行拷贝，完全生成一个新的slice
	目标切片必须分配过空间且足够承载复制的元素个数，并且来源和目标的类型必须一致，copy() 函数的返回值表示实际发生复制的元素个数。
	copy(destSlice, srcSlice []T)
	srcSlice: 数据来源切片
	destSlice: 目标切片

	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 5)
	copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c


	slice遍历
	s := []int{1, 3, 5}

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	for index, value := range s {
		fmt.Println(index, value)
	}


	append()方法为切片添加元素

	var s []int
	s = append(s, 1)        // [1]
	s = append(s, 2, 3, 4)  // [1 2 3 4]
	s2 := []int{5, 6, 7}
	//在s的尾部追加s2
	s = append(s, s2...)
	s = append(s, []]int{5, 6, 7})

	在切片的开头添加元素
	var a = []int{1,2,3}
	a = append([]int{0}, a...) // 在开头添加1个元素
	a = append([]int{-3,-2,-1}, a...) // 在开头添加1个切片


	slice删除元素
	使用append进行删除
	a = []int{1, 2, 3, ...}
	a = append(a[:i], a[i+1:]...) // 删除下标为i的1个元素
	a = append(a[:i], a[i+N:]...) // 删除下标为i的N个元素
	本质就是 Go语言中删除切片元素的本质是，以被删除元素为分界点，将前后两个部分的内存重新连接起来。

	如果需要频繁的添加和删除元素，那么slice和数组的性能并不是非常好，建议使用链表


	多维切片
	//声明一个二维切片
	var slice [][]int
	//为二维切片赋值
	slice = [][]int{{10}, {100, 200}}
	// 声明一个二维整型切片并赋值
	slice := [][]int{{10}, {100, 200}}


*/
func testSlice() {
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	// var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)        //[]
	fmt.Println(b)        //[]
	fmt.Println(c)        //[false true]
	fmt.Println(a == nil) //true
	fmt.Println(b == nil) //false,已经分配了内存，所以不为nil
	fmt.Println(c == nil) //false,已经分配了内存，所以不为nil
	// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较

	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[1:3] // s := a[low:high]
	//cap = len(arr)-1=5-1=4
	fmt.Printf("s's Type %T s:%v len(s):%v cap(s):%v\n", s, s, len(s), cap(s))

	//可以省略切片表达式中的任何索引。省略了low则默认为0；省略了high则默认为切片操作数的长度:
	s1 := arr[2:]    // 等同于 a[2:len(a)::len(a)]
	s2 := arr[:3]    // 等同于 a[0:3::len(a)]
	s3 := arr[:]     // 等同于 a[0:len(a)::len(a)]
	s4 := arr[0:0]   // 空切片,等同于a[0:0::len(a)]
	s5 := arr[1:3:4] // 手工指定max,cap=max-begin = 3
	//cap = len(arr) - 2 = 5-2=3
	fmt.Printf("s1's Type %T s:%v len(s):%v cap(s):%v\n", s1, s1, len(s1), cap(s1))
	//cap = len(arr) - 0 = 5-0=5
	fmt.Printf("s2's Type %T s2:%v len(s2):%v cap(s2):%v\n", s2, s2, len(s2), cap(s2))
	//cap = len(arr) - 0 = 5-0=5
	fmt.Printf("s3's Type %T s3:%v len(s3):%v cap(s3):%v\n", s3, s3, len(s3), cap(s3))
	//cap = len(arr) - 0 = 5-0=5
	fmt.Printf("s4's Type %T s4:%v len(s4):%v cap(s4):%v\n", s4, s4, len(s4), cap(s4))
	// cap=3
	fmt.Printf("s5's Type %T s5:%v len(s5):%v cap(s5):%v\n", s5, s5, len(s5), cap(s5))

	//对切片操作修改值，实际是修改数组的元素值
	s1[0] = 10
	fmt.Printf("arr = %v\n", arr)
	//make创建clice，不传cap的话，默认cap=size
	var makeSlice = make([]int, 2)
	fmt.Printf("makeSlice's Type %T makeSlice:%v len(makeSlice):%v cap(makeSlice):%v\n", makeSlice, makeSlice, len(makeSlice), cap(makeSlice))

	//直接声明一个新的slice，size和cap都为0
	var numbers []int
	fmt.Printf("numbers's Type %T numbers:%v len(numbers):%v cap(numbers):%v\n", numbers, numbers, len(numbers), cap(numbers))
	//直接往slice里面一次性添加元素，cap会扩容为size + 1
	numbers = append(numbers, 1, 2, 3, 4, 5, 6, 7, 1, 1)
	fmt.Printf("numbers's Type %T numbers:%v len(numbers):%v cap(numbers):%v\n", numbers, numbers, len(numbers), cap(numbers))

	// fmt.Printf("sAppend's Type %T sAppend:%v len(sAppend):%v cap(sAppend):%v\n", sAppend, sAppend, len(sAppend), cap(sAppend))
	fmt.Printf("arr = %v\n", arr)
	//动态每次添加一个元素，slice会按照2^n 进行动态扩容
	//旧切片长度小于1024，扩容时直接就是oldcap*2
	//旧切片长度大于1024，扩容时循环扩容，每次增加 oldcap*0.25
	//扩容需要重新分配新的扩容后的内存，然后将旧数据拷贝到新的内存里面去，同时改变指针的指向
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	var headInsertSlice = make([]int, 2, 5)
	var headInsertSlice2 = []int{4, 2, 3}
	fmt.Printf("headInsertSlice's Type %T headInsertSlice:%v len(headInsertSlice):%v cap(headInsertSlice):%v ptr:%p \n", headInsertSlice, headInsertSlice, len(headInsertSlice), cap(headInsertSlice), headInsertSlice)
	// 把s2追加到s1的尾部，返回一个新的slice，slice指针指向s1头部，如果append之后的slice没有扩容，则指针地址不变，否则头指针指向扩容后的新slice
	headInsertSlice = append(headInsertSlice, headInsertSlice2...)
	fmt.Printf("headInsertSlice's Type %T headInsertSlice:%v len(headInsertSlice):%v cap(headInsertSlice):%v ptr:%p \n", headInsertSlice, headInsertSlice, len(headInsertSlice), cap(headInsertSlice), headInsertSlice)
	// s1追加到新切片的尾部，返回一个新的slice，slice指针指向新切片的头部
	headInsertSlice = append([]int{0}, headInsertSlice...)
	fmt.Printf("headInsertSlice's Type %T headInsertSlice:%v len(headInsertSlice):%v cap(headInsertSlice):%v ptr:%p \n", headInsertSlice, headInsertSlice, len(headInsertSlice), cap(headInsertSlice), headInsertSlice)
	// s1追加到新切片的尾部，返回一个新的slice，slice指针指向新切片的头部
	headInsertSlice = append([]int{-3, -2, -1}, headInsertSlice...)
	fmt.Printf("headInsertSlice's Type %T headInsertSlice:%v len(headInsertSlice):%v cap(headInsertSlice):%v ptr:%p \n", headInsertSlice, headInsertSlice, len(headInsertSlice), cap(headInsertSlice), headInsertSlice)
	//清空slice
	headInsertSlice = headInsertSlice[0:0]
	fmt.Printf("headInsertSlice's Type %T headInsertSlice:%v len(headInsertSlice):%v cap(headInsertSlice):%v ptr:%p \n", headInsertSlice, headInsertSlice, len(headInsertSlice), cap(headInsertSlice), headInsertSlice)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{7, 8, 9}
	fmt.Printf("slice1's Type %T slice1:%v len(slice1):%v cap(slice1):%v ptr:%p \n", slice1, slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2's Type %T slice2:%v len(slice2):%v cap(slice2):%v ptr:%p \n", slice2, slice2, len(slice2), cap(slice2), slice2)
	// 只会复制slice1的前3个元素到slice2中
	//返回值是拷贝的元素数量
	//copy(destSlice, srcSlice []T)
	var cCount = copy(slice2, slice1)
	fmt.Println(cCount)
	fmt.Printf("slice2's Type %T slice2:%v len(slice2):%v cap(slice2):%v ptr:%p \n", slice2, slice2, len(slice2), cap(slice2), slice2)
	slice2 = []int{7, 8, 9}
	//copy(destSlice, srcSlice []T)
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Printf("slice1's Type %T slice1:%v len(slice1):%v cap(slice1):%v ptr:%p \n", slice1, slice1, len(slice1), cap(slice1), slice1)

	slice1[0] = 100
	fmt.Printf("slice1's Type %T slice1:%v len(slice1):%v cap(slice1):%v ptr:%p \n", slice1, slice1, len(slice1), cap(slice1), slice1)

	fmt.Printf("slice2's Type %T slice2:%v len(slice2):%v cap(slice2):%v ptr:%p \n", slice2, slice2, len(slice2), cap(slice2), slice2)

}

/*

 */
func testRange() {
	// 创建一个整型切片，并赋值
	slice := []int{10, 20, 30, 40}
	// 迭代每一个元素，并显示其值,index和value的名称可以随便定义，比如 for i, v := range slice
	// 返回的value是数组元素的一个拷贝，两者的内存地址不同
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}
}

/*
	1.map 是引用类型，可以使用如下方式声明：
	var mapname map[keytype]valuetype
	举例： var mapLit map[string]int = map[string]int{"one": 1, "two": 2}

	mapname 为 map 的变量名。
	keytype 为键类型。
	valuetype 是键对应的值类型。

	在声明的时候不需要知道 map 的长度，因为 map 是可以动态增长的，未初始化的 map 的值是 nil，使用函数 len() 可以获取 map 中 pair 的数目。
	map在声明之后需要指定内存空间，只声明不分配内存空间是无法往里面放数据的

	2.和数组不同，map 可以根据新增的 key-value 动态的伸缩，因此它不存在固定长度或者最大限制，但是也可以选择标明 map 的初始容量 capacity，格式如下：
	make(map[keytype]valuetype, cap)
	举例：map2 := make(map[string]float, 100)

	make方法创建的map系统会默认分配内存空间，因此不用在手动指定了

	当 map 增长到容量上限的时候，如果再增加新的 key-value，map 的大小会自动加 1，所以出于性能的考虑，对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明。

	3.使用切片作为map的值
	//key int value intSlice
	mp1 := make(map[int][]int)

	//key int value intSlice类型的指针
	mp2 := make(map[int]*[]int)

	4.遍历map
	使用range即可

	5.从map中删除元素
	使用 delete() 内建函数从 map 中删除一组键值对，delete() 函数的格式如下：
	delete(map, 键)


	6.清空 map 中的所有元素
	创建一个新的map，垃圾回收机制会把旧map清除
	old = make(...)

	7.map存在线程安全问题
	多线程操作同一个map的情况下，可使用sync.Map

*/
func testMap() {
	//初始化map
	var mapLit map[string]int

	//var mapCreated map[string]float32
	var mapAssigned map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}
	//等价于 mapCreated := map[string]float{}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3
	var three = mapLit["three"]
	var four = mapLit["four"]
	fmt.Printf("%v %v\n", three, four)
	fmt.Printf("%p \n", mapLit)
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
	//key int value intSlice
	mp1 := make(map[int][]int)
	//key int value intSlice类型的指针
	mp2 := make(map[int]*[]int)

	slice1 := make([]int, 10, 10)

	mp1[1] = slice1
	mp2[1] = &slice1

	fmt.Printf("mp1 Type =  %T , mp1 value =  %v \n", mp1, mp1)
	fmt.Printf("mp2 Type =  %T , mp2 value =  %v &value %v \n", mp2, mp2, *mp2[1])

	//map的遍历
	for k, v := range mapLit {
		fmt.Printf("key =  %v , value = %v \n", k, v)
	}

	//从map中删除元素
	delete(mapLit, "one")
	fmt.Printf("mapLit =  %v\n", mapLit)

	//清除map中的所有元素
	mapLit = make(map[string]int, 10)
	fmt.Printf("mapLit =  %v\n", mapLit)

}

/*
	Go语言在 1.9 版本中提供了一种效率较高的并发安全的 sync.Map，sync.Map 和 map 不同，不是以语言原生形态提供，而是在 sync 包下的特殊结构。

	sync.Map 有以下特性：
	无须初始化，直接声明即可。
	sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除。
	使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false。

	sync.Map 没有提供获取 map 数量的方法，替代方法是在获取 sync.Map 时遍历自行计算数量，sync.Map 为了保证并发安全有一些性能损失，因此在非并发情况下，使用 map 相比使用 sync.Map 会有更好的性能。
*/
func testSyncMap() {
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	//遍历需要提供一个匿名函数，参数为 k、v，类型为 interface{}， 返回值为布尔类型一般一定要返回true，返回false的话range会break
	//每次 Range() 在遍历一个元素时，Range方法会调用这个匿名函数，把key value传进来
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		fmt.Printf("type : %T\n", k)
		return true
	})

}

/*
	在Go语言中，列表使用 container/list 包来实现，内部的实现原理是 双链表 ，列表能够高效地进行任意位置的元素插入和删除操作。
		1) 通过 container/list 包的 New() 函数初始化 list
		变量名 := list.New()

		2) 通过 var 关键字声明初始化 list
		var 变量名 list.List

	列表与切片和 map 不同的是，列表并没有具体元素类型的限制。
	因此，列表的元素可以是任意类型，这既带来了便利，也引来一些问题。
	例如给列表中放入了一个 interface{} 类型的值，取出值后，如果要将 interface{} 转换为其他类型将会发生宕机。

	1.在列表中插入元素
	双链表支持从队列前方或后方插入元素，分别对应的方法是 PushFront 和 PushBack。
	这两个方法都会返回一个 *list.Element 结构，如果在以后的使用中需要删除插入的元素，则只能通过 *list.Element 配合 Remove() 方法进行删除，这种方法可以让删除更加效率化，同时也是双链表特性之一。

	l := list.New()
	l.PushBack("fist")
	l.PushFront(67)

	方  法														功  能
	InsertAfter(v interface {}, mark * Element) * Element	在 mark 点之后插入元素，mark 点由其他插入函数提供
	InsertBefore(v interface {}, mark * Element) *Element	在 mark 点之前插入元素，mark 点由其他插入函数提供
	PushBackList(other *List)								添加将两一个list列表元素到现有list尾部
	PushFrontList(other *List)								添加将两一个list列表元素到现有list头部

	type Element struct {
		next, prev *Element

		list *List

		Value interface{}
	}




*/
func testList() {
	// l := list.New()
	// var ptr = l.PushBack("fist")
	// l.PushFront(67)

	// fmt.Printf("%T %v \n", ptr, ptr)

	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushFront(67)
	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")
	// 在fist之后添加high
	l.InsertAfter("high", element)
	// 在fist之前添加noon
	l.InsertBefore("noon", element)
	// 使用
	l.Remove(element)
	// 遍历双链表需要配合 Front() 函数获取头元素，遍历时只要元素不为空就可以继续进行，每一次遍历都会调用元素的 Next() 函数
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
