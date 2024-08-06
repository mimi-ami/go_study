package main

import "fmt"

// 关于指针

// 任何程序数据载入内存之后，在内存中都有他们的地址，这就是指针，为了保存一个数据在内存中的地址，需要指针变量
// 指针不能进行偏移和运算 所以只需要两个符号  & 取址 * 根据地址取值

// 指针地址和指针类型
//ptr := &v  v的类型为T ptr的类型为 *T    * 代表指针

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

// 关于new 和 make

func cc() {
	// a := 10
	// b := &a
	// fmt.Printf("type of b:%T\n", b)
	// c := *b //指针取值 根据指针去内存中取值
	// fmt.Printf("type of c:%T\n", c)
	// fmt.Printf("value of c:%v\n", c)

	// a := 10
	// modify1(a)
	// fmt.Println(a)
	// modify2(&a)
	// fmt.Println(a)

	// new  make 是用来分配内存的两个函数

	//func new(Type) *Type
	// a := new(int)
	// b := new(bool)
	// fmt.Println("%T\n", a) // *int
	// fmt.Println("%T\n", b) // *bool
	// fmt.Println(*a) // 0
	// fmt.Println(*b) // false
	// var a *int
	// a = new(int)
	// *a = 10
	// fmt.Println(*a, a)
	// make也是用来做内存分配的，和new的区别是，make只用于slice map channel的内存创建
	// func make(t Type, size ...IntegerType) Type
	// 声明为map类型的变量也需要用make进行初始化操作之后才能进行赋值
	var b map[string]int
	b = make(map[string]int, 10)
	b["ss"] = 100
	fmt.Println(b)
}
