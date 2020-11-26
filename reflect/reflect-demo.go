package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (u User) Print(prfix string) {
	fmt.Printf("%s:Name is %s,Age is %d", prfix, u.Name, u.Age)
}

func main() {
	u := User{"张三", 20}
	//获取reflect.type类型
	t := reflect.TypeOf(u)
	fmt.Println(t)
	fmt.Println(t.Kind())
	//遍历字段
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}
	//遍历方法
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
	}
	//转换成Reflect.value类型
	v := reflect.ValueOf(u)
	fmt.Println(v)
	//反射调用方法
	mPrint := v.MethodByName("Print")
	args := []reflect.Value{reflect.ValueOf("前缀")}
	fmt.Println(mPrint.Call(args))
}
