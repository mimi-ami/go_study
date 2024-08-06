package main

import "fmt"

// 接口 一个接口类型就是一组方法的集合，规定了需要实现的所有方法

/* type 接口类型名 interface {
	方法名1 （参数） 返回值
	方法名2 （参数） 返回值
}
*/ // 当方法名首字母和接口类型名首字母都是大写时，这个方法才可以被接口所在的包之外的代码访问
// type Writer  interface {
// 	Write([]byte) error
// }
// // 一个类型实现了接口中规定的所有方法，就称实现了这个接口
// type Singer interface {
// 	Sing()
// }
// type Bird struct {}

// func (b Bird) Sing() {
// 	fmt.Println("mimimi")
// }

// type Cat struct{}

// func (c Cat) Say() {
// 	fmt.Println("喵喵")
// }
// type Sheep struct{}

// func (s Sheep) Say() {
// 	fmt.Println("咩咩")
// }

// func CatHungry(c Cat) {
// 	c.Say()
// }
// func SheepHungry(s Sheep) {
// 	s.Say()
// }

// func main() {
// 	c := Cat{}
// 	c.Say()

// }

// type Sayer interface {
// 	Say()
// }

// func MakeHungry(s Sayer) {
// 	s.Say()
// }

// type ZhiFuBao struct {
// }

// func (z *ZhiFuBao) Pay(amount int64) {
// 	fmt.Println("使用支付宝付款：%.2f \n", float64(amount/100))
// }
// func Checkout(obj *ZhiFuBao) {
// 	obj.Pay(100)
// }

// func main() {
// 	Checkout(&ZhiFuBao{})
// }

// type Mover interface {
// 	Move()
// }

// type Dog struct{}

// func (d Dog) Move() {
// 	fmt.Println("dogs can move")
// }

// 值接收者实现接口与 指针接收者实现接口
// 一个类型可以同时实现多个接口，而接口间彼此独立
// type Sayer interface {
// 	Say()
// }

// type Mover interface {
// 	Move()
// }

// type Dog struct {
// 	Name string
// }

// func (d Dog) Say() {
// 	fmt.Printf("%s can cry\n", d.Name)
// }

// func (d Dog) Move() {
// 	fmt.Printf("%s can move", d.Name)
// }

// func main() {
// 	var d = Dog{Name: "nini"}
// 	var s Sayer = d
// 	var m Mover = d

// 	s.Say()
// 	m.Move()
// }

// 多种类型实现同一接口
// func (d Dog) Move() {
// 	fmt.Printf("%s can move", d.Name)
// }

// type Car struct {
// 	Brand string
// }

// func (c Car) Move() {
// 	fmt.Printf("%s speed", c.Brand)
// }
// func main() {
// 	var obj Mover
// 	obj = Dog{Name: "aa"}
// 	obj.Move()

// 	obj = Car{Brand: "bb"}
// 	obj.Move()
// }

// 一个接口的所有方法 不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现
// type Washer interface {
// 	wash()
// 	dry()
// }

// type dryer struct{}

// func (d dryer) dry() {
// 	fmt.Println("dry")
// }

// type haier struct {
// 	dryer
// }

// func (h haier) wash() {
// 	fmt.Println("wash")
// }

// 接口也可以作为结构体的一个字段
// type Interface interface {
// 	Len() int
// 	Less(i, j int) bool
// }

// type reverse struct {
// 	Interface // 在结构体中嵌入了一个接口类型，从而让结构体类型实现了该接口类型
// }

// // 改写接口中的方法
// func (r reverse) Less(i, j int) bool {
// 	return r.Interface.Less(j, i) // 重写了Less方法
// }

// // 空接口可以作为函数的参数，也可以作为map的值
// var x interface{}
// var student = make(map[string]interface{})

// 接口值
// 接口类型的值可以是任意一个实现了该接口的类型值，接口除了需要记录具体值，还需要记录这个值属于的类型

type Mover interface {
	Move()
}

type Dog struct {
	Name string
}

func (d *Dog) Move() {
	fmt.Println("dog can run")
}

type Car struct {
	Brand string
}

func (c *Car) Move() {
	fmt.Println("car can run")
}

func main() {
	var m Mover
	fmt.Println(m == nil)
	m = &Dog{Name: "abc"}

	
// 类型断言
// 接口值可以赋值为任意类型的值
var m Mover
m = &Dog{Name: "aa"}
fmt.Println("%T\n", m)

m = new(Car)
