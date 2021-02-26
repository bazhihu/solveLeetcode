package golang

import (
	"fmt"
	"reflect"
)

/**
反射：是程序在运行时可以访问、检测、修改自身状态或者行为的一种能力。

反射是把双刃剑，虽然代码更加灵活了，但也有问题：
1、代码阅读起来困难了
2、一定程序上破坏了静态类型语言的编译期检查，运行时会有panic风险
3、降低了系统性能

需要反射的时机
1、无法预定义参数类型
2、函数需要根据入参来动态执行

*** Go中只有接口类型 才可以反射

//-----
interface{}
接口结构：
	类型 (Type)
	数据（Value）
//-----

反射 Reflect
两个重要类型

Type  => reflect.Type    ==> reflect.TypeOf()
Value => reflect.Value   ==> reflect.ValueOf()
类型   =>   类型         ==> 获取方法

反射三大定律
1、接口数据 --> 反射

var x float64 = 3.4
fmt.Println("type:", reflect.TypeOf(x))

2、反射对象 --> 接口数据

y := v.Interface().(float64)
fmt.Println(y)

3、若数据可修改，可通过反射对象来修改它

var a float64
fmt.Println(a)
va := reflect.ValueOf(&a)
va.Elem().SetFloat(11)
fmt.Println(a)

Type 重要的一些方法

// 返回类型的特定种类
Kind() kind

// 判断该类型是否可赋值给u类型
AssignableTo(u Type) bool

// 返回元素类型
// 非Array, Chan, Map, Ptr, or Slice 会panic
Elem() Type

// 返回结构体类型的 第i个字段
Field(i int) StructField

// 返回结构体类型字段数量
NumField() int

**** reflect.Value 通过Value.Type() 也可以直接获取 reflect.Type


反射的应用
1、对象序列化
2、OR（object Relational Mappingsss）

*/

type Animal interface {
	Say()
}

type Dog struct {
}

func (d *Dog) Say() {
	fmt.Println("wang wang")
}
func main() {
	// 1
	//var animal Animal
	//
	//dog := &Dog{}
	//// 2
	//animal = dog
	//
	//// 3
	//var e interface{}
	//e = dog

	var a float64
	fmt.Println(a)
	va := reflect.ValueOf(&a)
	va.Elem().SetFloat(11)
	fmt.Println(a)
}
