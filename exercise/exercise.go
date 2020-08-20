package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	countChineseWords()
	printSin()
	fmt.Println("------------student system -----------")
	testStudentSystem()
}

func countChineseWords() {
	words := "hello沙河小王子"
	count := 0
	//汉字的unicode码区间为0x2e80 - 0x9fff
	const startCode = 0x2E80
	const endCode = 0x9FFF
	for _, w := range words {
		if startCode < w && endCode > w {
			count++
		}
	}
	fmt.Println("中文字符数:", count)

	var (
		str  = "hello沙河小王子"
		char = []rune(str)
		d    = 0
	)
	for i := 0; i < len(char); i++ {
		//汉字在UTF8中的长度为3个字节。
		if len(string(char[i])) >= 3 {
			d++
		}
	}
	fmt.Printf("汉字总计数量：%d \n", d)

}

func printSin() {
	// 图片大小
	const size = 300
	// 根据给定大小创建灰度图
	// image.Rect根据矩形的两个点绘制一张矩形图片
	pic := image.NewGray(image.Rect(0, 0, size, size))
	// 遍历每个像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			// 填充为白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	// 从0到最大像素生成x坐标
	for x := 0; x < size; x++ {
		// 让sin的值的范围在0~2Pi之间
		s := float64(x) * 2 * math.Pi / size
		// sin的幅度为一半的像素。向下偏移一半像素并翻转
		y := size/2 - math.Sin(s)*size/2
		// 用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}
	// 创建文件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	// 使用png格式将数据写入文件
	png.Encode(file, pic) //将image信息写入文件中
	// 关闭文件
	file.Close()
}

func testStudentSystem() {
	class := newClassInstance()
	s1 := newStudentInstance(1, 2, "zhangsan", 12.33)
	s2 := newStudentInstance(2, 2, "lisi", 10.00)
	s3 := newStudentInstance(3, 2, "wangwu", 1.33)
	class.show()
	class.addStu(s1)
	class.addStu(s2)
	class.addStu(s3)

	class.show()

	class.delStu(1)
	class.show()

	s4 := class.getStu(2)
	fmt.Printf("add : %p , value :  %v \n", s4, s4)
}

type Student struct {
	Id    int
	Name  string
	Age   int
	Score float32
}

type Class struct {
	students map[int]*Student
}

func newClassInstance() *Class {
	students := make(map[int]*Student)
	return &Class{
		students: students,
	}
}

func newStudentInstance(id, age int, name string, score float32) *Student {
	return &Student{id, name, age, score}
}

func (c *Class) addStu(student *Student) {
	c.students[student.Id] = student
}
func (c *Class) getStu(id int) *Student {
	return c.students[id]
}

func (c *Class) delStu(id int) {
	delete(c.students, id)
}
func (c *Class) show() {
	fmt.Println("------------print class-----------")
	if len(c.students) == 0 {
		fmt.Println("empty class")
		return
	}
	for key, val := range c.students {
		fmt.Printf(" key : %v ,value : %v \n", key, val)
	}
}
