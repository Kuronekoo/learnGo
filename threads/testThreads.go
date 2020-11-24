package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
	一、并发与并行

		顺序(sequential)：多个任务顺序执行，先开始的任务一定先完成
		并发(concurrent)：多个任务可以争抢cpu空闲资源，先开始的任务不一定先完成（处理多个任务的能力）。

		串行(serial)：同一时刻只有一个任务在执行
		并行(parallel)：同一时刻执行多个任务（同时处理多个任务的能力）。

		单核CPU：并发，串行
		多核CPU：并发，并行

		Go语言的并发通过goroutine实现。goroutine类似于线程，属于用户态的线程，我们可以根据需要创建成千上万个goroutine并发工作。goroutine是由Go语言的运行时（runtime）调度完成，而线程是由操作系统调度完成。

		Go语言还提供channel在多个goroutine间进行通信。goroutine和channel是 Go 语言秉承的 CSP（Communicating Sequential Process）并发模式的重要实现基础。

		在java/c++中我们要实现并发编程的时候，我们通常需要自己维护一个线程池，并且需要自己去包装一个又一个的任务，同时需要自己去调度线程执行任务并维护上下文切换，这一切通常会耗费程序员大量的心智。那么能不能有一种机制，程序员只需要定义很多个任务，让系统去帮助我们把这些任务分配到CPU上实现并发执行呢？

		Go语言中的goroutine就是这样一种机制，goroutine的概念类似于线程，但 goroutine是由Go的运行时（runtime）调度和管理的。Go程序会智能地将 goroutine 中的任务合理地分配给每个CPU。Go语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制。

		在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能–goroutine，当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个goroutine去执行这个函数就可以了，就是这么简单粗暴。

	二、使用goroutine

		Go语言中使用goroutine非常简单，只需要在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine。

		一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数。

		在程序启动时，Go程序就会为main()函数创建一个默认的goroutine。

		当main()函数返回的时候该goroutine就结束了，所有在main()函数中启动的goroutine会一同结束
	三、channel
		单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。

		虽然可以使用共享内存进行数据交换，但是共享内存在不同的goroutine中容易发生竞态问题。为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。

		Go语言的并发模型是CSP（Communicating Sequential Processes 顺序通信进程），提倡通过通信共享内存而不是通过共享内存而实现通信。

		如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。

		Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。


		channel是一种类型，一种引用类型。声明通道类型的格式如下：
			var 变量 chan 元素类型
		创建channel的格式如下：
			make(chan 元素类型, [缓冲大小])
		channel操作
			通道有发送（send）、接收(receive）和关闭（close）三种操作。
			发送和接收都使用<-符号。
		我们通过调用内置的close函数来关闭通道。
			close(ch)
		只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。
		关闭后的通道有以下特点：
			1.对一个关闭的通道再发送值就会导致panic。
			2.对一个关闭的通道进行接收会一直获取值直到通道为空。
			3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
			4.关闭一个已经关闭的通道会导致panic。
	四、worker pool（goroutine池）
		在工作中我们通常会使用可以指定启动的goroutine数量–worker pool模式，控制goroutine的数量，防止goroutine泄漏和暴涨。
	五、select多路复用
		Go内置了select关键字，可以同时响应多个通道的操作。

		select的使用类似于switch语句，它有一系列case分支和一个默认的分支。每个case会对应一个通道的通信（接收或发送）过程。select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句。

			select{
				case <-ch1:
					...
				case data := <-ch2:
					...
				case ch3<-data:
					...
				default:
					默认操作
			}
		使用select语句能提高代码的可读性。
			可处理一个或多个channel的发送/接收操作。
			如果多个case同时满足，select会随机选择一个。
			对于没有case的select{}会一直等待，可用于阻塞main函数。
	六、锁
	七、sync.Once
		Go语言中的sync包中提供了一个针对只执行一次场景的解决方案–sync.Once。

		sync.Once只有一个Do方法，其签名如下：
		func (o *Once) Do(f func()) {}
		备注：如果要执行的函数f需要传递参数就需要搭配闭包来使用。
	八、原子操作
		代码中的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高。
		针对基本数据类型我们还可以使用原子操作来保证并发安全，因为原子操作是Go语言提供的方法它在用户态就可以完成，因此性能比加锁操作更好。
		Go语言中原子操作由内置的标准库sync/atomic提供。
		atomic包提供了底层的原子级内存操作，对于同步算法的实现很有用。
		这些函数必须谨慎地保证正确使用。除了某些特殊的底层应用，使用通道或者sync包的函数/类型实现同步更好。
		atomic包
		读取操作
		func LoadInt32(addr *int32) (val int32)
		func LoadInt64(addr *int64) (val int64)
		func LoadUint32(addr *uint32) (val uint32)
		func LoadUint64(addr *uint64) (val uint64)
		func LoadUintptr(addr *uintptr) (val uintptr)
		func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)

		修改操作
		func StoreInt32(addr *int32, val int32)
		func StoreInt64(addr *int64, val int64)
		func StoreUint32(addr *uint32, val uint32)
		func StoreUint64(addr *uint64, val uint64)
		func StoreUintptr(addr *uintptr, val uintptr)
		func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)
		修改操作
		func AddInt32(addr *int32, delta int32) (new int32)
		func AddInt64(addr *int64, delta int64) (new int64)
		func AddUint32(addr *uint32, delta uint32) (new uint32)
		func AddUint64(addr *uint64, delta uint64) (new uint64)
		func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)
		交换操作
		func SwapInt32(addr *int32, new int32) (old int32)
		func SwapInt64(addr *int64, new int64) (old int64)
		func SwapUint32(addr *uint32, new uint32) (old uint32)
		func SwapUint64(addr *uint64, new uint64) (old uint64)
		func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)
		func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)

		比较并交换操作
		func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
		func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
		func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)
		func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)
		func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)
		func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)

*/

func main() {
	fmt.Println("--------------testGoRutine--------------")
	testGoRutine()
	fmt.Println("--------------testChannel--------------")
	testChannel()
	fmt.Println("--------------testWorkPool--------------")
	testWorkPool()
	fmt.Println("--------------testSelectAndChannel--------------")
	testSelectAndChannel()
	fmt.Println("--------------testMutexLock--------------")
	testMutexLock()
	fmt.Println("--------------testSingleton--------------")
	testSingleton()
	fmt.Println("--------------testAtomic--------------")
	testAtomic()

}

func testAtomic() {
	c3 := AtomicCounter{} // 并发安全且比互斥锁效率更高
	test(&c3)
}

type Counter interface {
	Inc()
	Load() int64
}

func test(c Counter) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(c.Load(), end.Sub(start))
}

// 原子操作版
type AtomicCounter struct {
	counter int64
}

func (a *AtomicCounter) Inc() {
	atomic.AddInt64(&a.counter, 1)
}

func (a *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.counter)
}

func testSingleton() {
	//起3个线程去获取单例
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetInstance()
		}()
	}
	wg.Wait()
}

type singleton struct{}

var instance *singleton
var once sync.Once

//借助sync.Once实现的并发安全的单例模式：
//sync.Once其实内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，而布尔值用来记录初始化是否完成。
//这样设计就能保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次。
func GetInstance() *singleton {
	once.Do(func() {
		fmt.Println("init singleton...")
		instance = &singleton{}
	})
	return instance
}

var x int64
var y int64

//互斥锁
var lock sync.Mutex

//读写互斥锁
//读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
//当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。
var rwlock sync.RWMutex

func write() {
	// lock.Lock()   // 加互斥锁
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg.Done()
}

//互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。
//Go语言中使用sync包的Mutex类型来实现互斥锁。 使用互斥锁来修复上面代码的问题：
//使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；
//当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的。
func add() {
	defer wg.Done()
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
	}

}

func addUnsafe() {
	for i := 0; i < 5000; i++ {
		y += 1
	}
	wg.Done()
}

func testMutexLock() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)

	wg.Add(3)
	go addUnsafe()
	go addUnsafe()
	go addUnsafe()
	wg.Wait()
	fmt.Println(y)

}

func testSelectAndChannel() {
	//通道是引用类型，通道类型的空值是nil。
	var ch1 chan int  // 声明一个传递整型的通道
	var ch2 chan bool // 声明一个传递布尔型的通道

	//初始化无缓冲的通道
	ch1 = make(chan int)
	//初始化容量为1有缓存的通道
	//只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。
	//我们可以使用内置的len函数获取通道内元素的数量，使用cap函数获取通道的容量，虽然我们很少会这么做。
	ch2 = make(chan bool, 1)
	//往通道中放入数据
	go func() {
		ch1 <- 10
		ch1 <- 10
	}()

	go func() {
		time.Sleep(time.Second * 1)
		ch2 <- false
	}()
	//select的使用类似于switch语句，它有一系列case分支和一个默认的分支。
	//每个case会对应一个通道的通信（接收或发送）过程。
	//select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句。
	//循环的次数大于通道中消息的数量会报死锁
	//select在通道加入消息之前执行也会报死锁
	for i := 0; i < 2; i++ {
		select {
		case x := <-ch1:
			fmt.Println("channel ", x)
		case x := <-ch2:
			fmt.Println("channel ", x)
		}
	}
}

func testWorkPool() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	// 输出结果
	for a := 1; a <= 5; a++ {
		v := <-results
		fmt.Println(v)
	}
	close(results)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	//只要jobs通道没有关闭，那么range就会一直循环
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

//单向只写通道
//入参 可传入双向通道和单向只写通道
func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

//单向只读通道
//入参 可传入双向通道和单向只读通道
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func testChannel() {
	//通道是引用类型，通道类型的空值是nil。
	var ch1 chan int   // 声明一个传递整型的通道
	var ch2 chan bool  // 声明一个传递布尔型的通道
	var ch3 chan []int // 声明一个传递int切片的通道

	//初始化无缓冲的通道
	ch1 = make(chan int)
	//初始化容量为1有缓存的通道
	//只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。
	//我们可以使用内置的len函数获取通道内元素的数量，使用cap函数获取通道的容量，虽然我们很少会这么做。
	ch2 = make(chan bool, 1)
	ch3 = make(chan []int)
	//无缓冲的通道只有在有人接收值的时候才能发送值。
	//所以这里要另外起一个线程，
	//否则在无人接收情况下，【发送】这个动作会一直阻塞着，形成死锁,并且报错  all goroutines are asleep - deadlock!
	//无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，两个goroutine将继续执行。
	//相反，如果接收操作先执行，接收方的goroutine将阻塞，直到另一个goroutine在该通道上发送一个值。
	//使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道。
	go func() {
		//睡2秒
		// time.Sleep(time.Second * 2)
		//发送值到通道
		ch1 <- 10
	}()
	fmt.Println("main thread")
	//从ch中接收值并赋值给变量iV
	//在接收到channel中的消息前，会一直阻塞
	//通道关闭后再取值ok=false
	iV, ok := <-ch1
	if ok {
		fmt.Println(iV)
	}

	//直接在main线程里面发送一个消息，不会出现死锁
	ch2 <- true
	//发第二个还是会死锁的
	// ch2 <- true
	go func() {
		for i := range ch2 { // 通道关闭后会退出for range循环
			fmt.Println(i)
		}
	}()
	//睡1s
	time.Sleep(time.Second * 1)
	close(ch1)
	close(ch2)
	close(ch3)

}

//WaitGroup用来等待一堆线程结束。
//主线程调用Add(delta int)方法来设置要等待的线程数量
//每个线程在结束前需要调用Done()方法
//mian线程可以调用Wait()方法来等待所有的线程的完成
//WaitGroup不可以拷贝
//sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。
//例如当我们启动了N 个并发任务时，就将计数器值增加N。
//每个任务完成时通过调用Done()方法将计数器减1。
//通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成。
var wg sync.WaitGroup

func testGoRutine() {
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	//main线程睡一段时间
	// time.Sleep(time.Second * 10)

	//启动多个线程并发执行
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello2(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}

func hello2(i int) {
	defer wg.Done() // goroutine结束就登记-1
	//hello线程睡500ms
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Hello Goroutine with WaitGroup!", i)
}

func hello() {
	fmt.Println("Hello Goroutine!")
}
