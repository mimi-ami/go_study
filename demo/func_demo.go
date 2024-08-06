// func 函数名（参数）（返回值）
package main

import (
	"errors"
	"fmt"
	"strings"
)

// func intsum(x int, y int) int {
func intsum(x, y int) int { //相邻参数类型相同
	return x + y
}

func intsum2(x ...int) int { // 可变参数  函数的参数数量不固定，可变参数通常要作为函数的最后一个参数
	fmt.Println(x) // x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

func intsum3(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}

// 多返回值情况
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 返回值命名
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 返回值补充
//
//	func someFunc(x string) []int {
//		if x == "" {
//			return nil
//		}
//		return
//	}
//
// 变量作用域
// 全局变量
var num int64 = 10

func testGlobalVar() {
	fmt.Printf("num=%d\n", num)
}

// 局部变量 分为函数内定义有效的变量和语句块内定义的变量两种情况，通常在if for switch 语句中
func testlocalVar(x, y int) {
	fmt.Println(x, y)
	if x > 0 {
		z := 100 // z只在if语句中有效，在外部无效
		fmt.Println(z)
	}
	for i := 0; i < 10; i++ { // i 只在for语句中有效
		fmt.Println(i)
	}

}

// 函数类型与变量
// 定义函数类型
type calculation func(int, int) int // 定义了一个canculation类型
// 一个calculation类型 接收两个int类型的参数，返回一个int类型的参数，也就是说 所有符合这个条件的函数都是这个类型

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 高阶函数  包括两种：函数作为参数  函数作为返回值
// 函数作为参数
func calc3(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

// 匿名函数 没有参数名的函数
// func(参数)(返回值){函数体}
// 匿名函数因为没有函数名，所以不能像普通函数一样被调用，所以匿名函数需要保存到某个变量或者作为立即执行函数
// func main() {
// 	add := func(x, y int) {
// 		fmt.Println(x + y)
// 	}
// 	add(10, 20) //通过变量调用匿名函数

// 	// 自执行函数 匿名函数定义完 加一个()直接执行
// 	func(x, y int) {
// 		fmt.Println(x + y)
// 	}(10, 20)
// }

// 闭包 一个函数和与其相关的引用环境组合而成的实体 闭包=函数+引用环境
func adder() func(int) int {
	var x int
	fmt.Printf("x=%v\n", x)
	return func(y int) int {
		fmt.Printf("y=%v\n", y)
		x += y
		return x
	}
}

// 进阶
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 进阶2
func calc4(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int { // base是会一直传递的
		base -= i
		return base
	}
	return add, sub
}

// func main() {
// ret := intsum(10, 20)
// fmt.Println(ret)

// sum := intsum2(1, 23, 4)
// fmt.Println(sum)

// sum := intsum3(1, 2, 3, 4)
// fmt.Println(sum)

// sum, sub := calc2(10, 5)
// fmt.Println(sum, sub)
// testGlobalVar()

// var c calculation 函数类型变量
// c = add
// sum := c(10, 20)
// fmt.Println(sum)

// 	ret2 := calc3(10, 20, add)
// 	fmt.Println(ret2)
// var f = adder()
// fmt.Println(f(10))
// fmt.Println(f(20))
// fmt.Println(f(30))

// f1 := adder()
// fmt.Println(f1(40))

// jpgFunc := makeSuffixFunc(".jpg")
// txtFunc := makeSuffixFunc(".txt")
// fmt.Println(jpgFunc("test"))
// fmt.Println(txtFunc("tst"))

// f1, f2 := calc4(10)
// fmt.Println(f1(1), f2(2))
// fmt.Println(f1(3), f2(4))
// defer 语句会把他后面的语句进行延迟处理，在defer归属的函数即将返回时，将延迟处理的语句按defer的逆序执行
// 先被defer的语句最后执行，最后defer的语句先被执行
// 	fmt.Println("start")
// 	defer fmt.Println(1)
// 	defer fmt.Println(2)
// 	defer fmt.Println(3) // 方便用来处理资源释放的问题，如 资源清理 文件关闭 解锁 记录时间等。。
// 	fmt.Println("end")

// }

// defer经典案例
// func f1() int { //无输入参数 输出参数为int
// 	x := 5
// 	defer func() {
// 		x++  // x=6
// 	}()
// 	return x
// }

// func f2() (x int) {
// 	defer func() { // 在返回5之后，自己执行x++
// 		x++ //x = 7
// 	}()
// 	return 5
// }
// func f3() (y int) {
// 	x := 5
// 	defer func() {
// 		x++
// 	}()
// 	return x

// }

// func f4() (x int) {
// 	defer func(x int) { // 匿名函数  在后边加了(x) 所以直接执行不在等待
// 		x++
// 	}(x)
// 	return 5
// }

// func main() {
// 	fmt.Println(f1())
// 	fmt.Println(f2())
// 	fmt.Println(f3())
// 	fmt.Println(f4())

// }

func calc5(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// func main() {
// 	x := 1 // 声明并初始化一个新变量
// 	y := 2
// 	defer calc5("AA", x, calc5("A", x, y)) // 参数需要在defer注册时就定义完成，所以要先执行"A"
// 	x = 10                                 // 赋值操作 赋值给一个已经定义过的变量
// 	defer calc5("BB", x, calc5("B", x, y)) //
// 	y = 20
// }

// x=10 y=20 执行第二个defer 40 第一个defer

// 内置函数介绍
/*
close 关闭channel
len 用来求长度 string array slice map channel
new 用来分配内存 int struct 返回指针
make 分配内存，主要分配引用类型 chan map slice
append 追加元素到数组、slice中
panic recover 错误处理 panic可以在任何地方引发 recover只能在defer调用的函数中有效
*/

// func funcA() {
// 	fmt.Println("func A")
// }
// func funcB() {
// 	// defer必须要在引发panic之前定义
// 	defer func() {
// 		err := recover() // recover必须和defer一起使用
// 		//recover之后就可以继续程序啦
// 		if err != nil {
// 			fmt.Println("recover in B")
// 		}
// 	}()
// 	panic("panic in B")
// }
// func funcC() {
// 	fmt.Println("func C")
// }

// func main() {
// 	funcA()
// 	funcB()
// 	funcC()
// }

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	for _, name := range users {
		switch {
		case strings.Contains(name, "E") || strings.Contains(name, "e"):
			distribution[name]++
			coins--
		case strings.Contains(name, "I") || strings.Contains(name, "i"):
			distribution[name] += 2
			coins -= 2
		case strings.Contains(name, "O") || strings.Contains(name, "o"):
			distribution[name] += 3
			coins -= 3
		case strings.Contains(name, "U") || strings.Contains(name, "u"):
			distribution[name] += 4
			coins -= 4
		}
	}
	return coins
}

func cc() {
	left := dispatchCoin()
	for name, coins := range distribution {
		fmt.Printf("%v %v\n", name, coins)
	}
	fmt.Println("剩下：", left)
}
