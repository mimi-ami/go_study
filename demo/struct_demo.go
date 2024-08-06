package main

import (
	"fmt"
)

// 结构体的内嵌配合接口
// type MyInt int // 定义为int类型

// 类型别名
// type TypeAlias = Type

// 类型别名与类型定义的区别
// 类型定义
// type NewInt int

// 类型别名
// type MyInt = int

// type person struct {
// 	name string
// 	city string
// 	// 也可以写成 name, city string
// 	age int8
// }

// 结构体实例化时才会真正的分配内存，必须实例化之后才能使用结构体的字段
// func main() {
// var a NewInt
// var b MyInt
// fmt.Printf("type of a:%T\n", a)
// fmt.Printf("type of b:%T\n", b)

//  var p1 person
//  p1.name = "aa"
//  p1.city = "beijing"
//  p1.age = 18
//  fmt.Println("p1=%v\n", p1)
//  fmt.Println("p1=%#v\n", p1)
// 匿名结构体
// var user struct {
// 	Name string
// 	Age  int
// }
// user.Name = "king"
// user.Age = 34
// fmt.Println("%#v\n", user)

// 指针类型结构体
// 可以通过使用new关键字对结构体进行实例化，得到结构体的地址
// var p2 = new(person)
// fmt.Printf("%T\n", p2)
// fmt.Printf("p2=%#v\n", p2)

// var p2 = new(person)
// p2.name = "xiaowangzi"
// p2.age = 12
// p2.city = "beijing"
// fmt.Printf("p2=%#v\n", p2)

// 先取结构体的地址，然后实例化
// 使用&对结构体进行取地址操作相当于对该结构体进行了一次new实例化操作
// 	p3 := &person{}
// 	fmt.Printf("%T\n", p3)
// 	fmt.Printf("p3=%#v\n", p3)
// 	p3.name = "a"
// 	p3.age = 34
// 	p3.city = "sfd"
// 	fmt.Printf("p3=%#v\n", p3)
// }

// 结构体初始化
// type person struct {
// 	name string
// 	city string
// 	age  int8
// }

// func main() {
// 	var p4 person
// 	fmt.Printf("p4=%#v", p4)

// 	// 使用键值对初始化
// 	p5 := person{
// 		name: "ss",
// 		city: "beijign",
// 		age:  23,
// 	}
// 	fmt.Printf("p5=%#v\n", p5)

// 	// 对结构体指针进行键值对初始化,初始值可以不全写，此时没写的就默认为该类型的0值
// 	p6 := &person{
// 		name: "ss",
// 		city: "dsf",
// 		age:  12,
// 	}
// 	fmt.Printf("p6=%#v\n", p6)

// 	// 使用值的列表初始化,这种方式必须初始化结构体的所有字段，初始值的填充顺序必须与字段在结构体中的声明顺序一致，不能与键值初始化方式混用
// 	p8 := &person{
// 		"fsd",
// 		"fs",
// 		34,
// 	}
// 	fmt.Printf("p8=%#v\n", p8)

// }

// 结构体内存布局 结构体占用一块连续的内存
// type test struct {
// 	a int8
// 	b int8
// 	c int8
// 	d int8
// }

// func main() {
// 	n := test{
// 		1, 2, 3, 4,
// 	}

// 	fmt.Printf("n.a %p\n", &n.a)
// 	fmt.Printf("n.b %p\n", &n.b)
// 	fmt.Printf("n.c %p\n", &n.c)
// 	fmt.Printf("n.d %p\n", &n.d)
// 	var v struct{}
// 	fmt.Println(unsafe.Sizeof(v)) // 空结构体不占用空间
// }

// type student struct {
// 	name string
// 	age  int
// }

// // 结构体的构造函数 ,因为当结构体比较复杂的时候，值拷贝性能开销会比较大，所以返回结构体指针类型
// func newStudent(name string, age int) *student {
// 	return &student{
// 		name: name,
// 		age: age,
// 	}
// }

// func main() {
// m := make(map[string]*student)
// stus := []student{
// 	{name: "小王子", age: 18},
// 	{name: "娜扎", age: 23},
// 	{name: "大王八", age: 9000},
// }

// for _, stu := range stus {
// 	m[stu.name] = &stu
// }
// for k, v := range m {
// 	fmt.Println(k, "=>", v.name)
// }
// 调用构造函数
// 	p9 := newStudent("sss", 34)
// 	fmt.Println(p9)
// }
// method 方法是一种作用于特定类型变量的函数，这种特定类型叫做接收者，接收者类似于其他语言中的this /self

// func (接收者变量 接收者类型) 方法名（参数列表） （返回参数） {}

// type Person struct {
// 	name string
// 	age  int8
// }

// // 构造函数
// func NewPerson(name string, age int8) *Person {
// 	return &Person{
// 		name: name,
// 		age:  age,
// 	}
// }

// // 定义方法
// func (p Person) Dream() {
// 	fmt.Printf("%s的梦想是。。。\n", p.name)
// }

// func main() {
// 	p1 := NewPerson("ss", 45)
// 	p1.Dream()
// }

// 方法和函数的区别是  函数不属于任何类型 方法属于特定的类型

// 指针类型的接收者
// 指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量
// func (p *Person) SetAge(newAge int8) {
// 	p.age = newAge
// }

// 值类型的接收者
// 在代码运行时，会把接收者的值复制一份，在值类型的接收者的方法中可以获取接收者的成员值，但是修改操作只针对副本，无法修改接收者变量本身
// func (p Person) SetAge2(newAge int8) {
// 	p.age = newAge
// }

// func main() {
// 	// p1 := NewPerson("xiaowangzi", 25)
// 	// fmt.Println(p1.age)
// 	// p1.SetAge(30)
// 	// fmt.Println(p1.age)

// 	p1 := NewPerson("xiaowangzi", 25)
// 	p1.Dream()
// 	fmt.Println(p1.age)
// 	p1.SetAge2(30)
// 	fmt.Println(p1.age)
// }

// 接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法，我们基于内置的int类型使用type关键字可以自定义新的类型
// type MyInt int

// func (m MyInt) Sayhello() {
// 	fmt.Println("hello")
// }

// func main() {
// 	var m1 MyInt
// 	m1.Sayhello()
// 	m1 = 100
// 	fmt.Printf("%#v %T\n", m1, m1)
// }

// 结构体的匿名字段  结构体允许成员字段在声明时没有字段名只有类型，这种没有名字的字段就称为匿名字段

// type Person struct { // 一个结构体中 同种类型的匿名字段只有一个
// 	string
// 	int
// }
// func main() {
// 	p1 := Person{
// 		"xiaowangzi",
// 		18,
// 	}
// 	fmt.Printf("%#v\n", p1)
// 	fmt.Println(p1.string, p1.int)
// }

// 嵌套结构体
// type Address struct {
// 	Province string
// 	City string
// }

// type User struct {
// 	Name string
// 	Gender string
// 	Address

// }

// func main() {
// 	user1 := User{
// 		Name: "aa",
// 		Gender: "nv",
// 		Address: Address{
// 			Province: "shandong",
// 			City: "weihai",
// 		},
// 	}
// 	var user2 User
// 	user2.Name = "xiaowangzi"
// 	user2.Gender = "nan"
// 	user2.Address.Province = "shandong"
// 	user2.Address.City = "weihai"
// 	fmt.Printf("user1=%v\n", user1)
// 	fmt.Printf("user1=%v\n", user2)

// }

// 嵌套结构体命名相同
// type Address struct {
// 	Province string
// 	City string
// 	createTime string
// }

// // email 邮箱
// type Email struct {
// 	Account string
// 	createTime string
// }

// type User struct {
// 	Name string
// 	Gender string
// 	Address
// 	Email
// }

// func main() {
// 	var user3 User
// 	user3.Name = "aa"
// 	user3.Gender = "nv"
// 	user3.Address.createTime = "2000"
// 	user3.Email.createTime = "2000"
// }

// 结构体的继承
// type Animal struct {
// 	name string
// }

// func (a *Animal) move() {
// 	fmt.Println("%s can move\n", a.name)
// }

// type Dog struct {
// 	Feet    int8
// 	*Animal // 通过嵌套匿名结构体实现继承
// }

// func (d *Dog) wang() {
// 	fmt.Printf("%s can wolf\n", d.name)
// }

// func main() {
// 	d1 := &Dog{
// 		Feet: 4,
// 		Animal: &Animal{
// 			name: "lele",
// 		},
// 	}
// 	d1.wang()
// 	d1.move()
// }

// 结构体中字段大写开头表示可以公开访问，小写表示私有，仅在定义当前结构体的包中可访问
// 结构体与json序列化
// type Student struct {
// 	ID     int
// 	Gender string
// 	Name   string
// }

// type Class struct {
// 	Title    string
// 	Students []Student
// }

// func main() {
// 	c := &Class{
// 		Title:    "101",
// 		Students: make([]Student, 0, 200),
// 	}
// 	for i := 0; i < 10; i++ {
// 		stu := Student{
// 			Name:   fmt.Sprintf("stu%02d", i),
// 			Gender: "nan",
// 			ID:     i,
// 		}
// 		c.Students = append(c.Students, stu)
// 	}
// 	data, err := json.Marshal(c)
// 	if err != nil {
// 		fmt.Println("json marshal failed")
// 		return
// 	}
// 	fmt.Printf("json:%s\n", data)

// 	// json反序列化  json格式的字符串->结构体
// 	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
// 	c1 := &Class{}
// 	err = json.Unmarshal([]byte(str), c1)
// 	if err != nil {
// 		fmt.Println("json unmarshal failed")
// 	}

// 	fmt.Printf("%#v\n", c1)

// }

// 结构体标签
// Tag 是结构体的元信息，可以在运行的时候通过反射的机制读取出来，tag在结构体字段的后方定义，由一对反引号包起来
// `key1: "value" key2: "value2"`
// type Student struct {
// 	ID int `json:"id"` // 通过tag实现json序列化该字段的key
// 	Gender string // json序列化默认使用字段名作为key
// 	name string // 私有不能被json包访问
// }

// func main() {
// 	s1 := Student{
// 		ID: 1,
// 		Gender: "nan",
// 		name: "nask",
// 	}
// 	data, err := json.Marshal(s1)
// 	if err != nil {
// 		fmt.Println("json marshal failed")
// 	}
// 	fmt.Printf("json str:%s\n", data)
// }

// 结构体和方法补充
// type Person struct {
// 	name   string
// 	age    int8
// 	dreams []string
// }

// func (p *Person) SetDreams(dreams []string) {

// 	p.dreams = make([]string, len(dreams))
// 	copy(p.dreams, dreams)
// }

// func main() {
// 	p1 := Person{name: "小王子", age: 18}
// 	data := []string{"吃饭", "shuijiao", "dadoudou"}
// 	p1.SetDreams(data)

// 	data[1] = "bushuijiao"
// 	fmt.Println(p1.dreams)
// }

