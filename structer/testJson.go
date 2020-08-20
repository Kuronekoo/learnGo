package main

import (
	"encoding/json"
	"fmt"
)

func testJson() {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name: fmt.Sprintf("stu%02d", i),
			ID:   i,
		}
		c.Students = append(c.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:\n%s\n", data)
	fmt.Println()
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"name":"张三"},{"ID":1,"name":"李四"},{"ID":2,"name":"王五"},{"ID":3,"name":"赵六"}]}`
	//穿件一个空结构
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}

//Student 学生
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	gender string //json序列化是默认使用字段名作为key
	Name   string //私有不能被json包访问
}

//Class 班级
type Class struct {
	Title    string
	Students []*Student
}
