package main

import (
	"fmt"
	"sync"
)

type userInfo struct {
	Name string
	Age  int
}

var m sync.Map

func main() {

	vv, ok := m.LoadOrStore("1", "one")
	fmt.Println(vv, ok) //one false

	vv, ok = m.Load("1")
	fmt.Println(vv, ok) //one true

	vv, ok = m.LoadOrStore("1", "oneone")
	fmt.Println(vv, ok) //one true

	vv, ok = m.Load("1")
	fmt.Println(vv, ok) //one true

	m.Store("1", "oneone")
	vv, ok = m.Load("1")
	fmt.Println(vv, ok) // oneone true

	m.Store("2", "two")
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
	fmt.Println("delete")
	m.Delete("1")
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})

	map1 := make(map[string]userInfo)
	var user1 userInfo
	user1.Name = "ChamPly"
	user1.Age = 24
	map1["user1"] = user1

	var user2 userInfo
	user2.Name = "Tom"
	user2.Age = 18
	m.Store("map_test", map1)

	mapValue, _ := m.Load("map_test")

	for k, v := range mapValue.(interface{}).(map[string]userInfo) {
		fmt.Println(k, v)
		fmt.Println("name:", v.Name)
	}
}

var myMap sync.Map

func init() {
	// 增
	myMap.Store("1", []string{"hi"}) //业务逻辑，其实两个都是string类型
	myMap.Store(1, 11111)
	// 查
	if val, ok := myMap.Load("1"); ok {
		fmt.Println("查", val)
	}
	if val, ok := myMap.Load(1); ok {
		fmt.Println("查int:", val)
	}
	// 改
	myMap.Store("1", 2222222)
	if val, ok := myMap.Load("1"); ok {
		fmt.Println("改", val)
	}
	// key是否存在，存在ok为true，key存储的值更新，不存在，OK为false， key存储该值
	if v, ok := myMap.LoadOrStore("22", "1223333"); ok {
		fmt.Println("LoadOrStore", v)
	}

	v, ok := myMap.LoadOrStore("221", "33333333")
	fmt.Println("v:", v, "ok:", ok)

	//v, ok = myMap.LoadOrStore("221", "221")
	//fmt.Println("v:", v, "ok:", ok)

	v, ok = myMap.Load("22")
	fmt.Println("v:", v, "ok:", ok)
	// 删除
	myMap.Delete("1")
	// 遍历
	f := func(key, value interface{}) bool {
		fmt.Println("遍历", key, value)
		return true
	}
	myMap.Range(f)
	fmt.Println("---------------------")
}
