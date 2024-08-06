// 定义了包，指明这个文件属于哪个包，package main表示一个可独立执行的程序，每个go应用程序都包含一个名为main的包
package mypackage

/* 引入包 */
import (
	"fmt"
	"math"
)

/* 启动后第一个先执行init函数，如果没有init函数，先执行main函数*/
func New() {
	fmt.Println("mypackage.New")
}

/* 关键字 break default func interface select case defer go map struct chan else goto package switch const failthrough if range type continue for import return var
 constants : true false iota nil
types: int int8 int16 int32 int64 unit uint8 uint16 uint32 uint64 uintptr float32 float64 complex128 complex64 bool byte rune string error
functions: make len cap new append copy close delete complex real imag panic recover
*/
/*
变量声明：
var 变量名 变量类型
var name string
var age int
var isOk bool
// 变量初始化完成之后，int和float初始化成0，str初始化成'', bool初始化为false。切片 函数 指针变量默认为nil
批量声明
var (
	a string
	b int
	c bool
	d float32
)
var name string ="sss"
var age int = 18
var name,age = "shd", 20
如果省略类型的定义，会根据指定的值推到变量的类型
var name = "ads"
*/

// var m = 100 // 全局变量

// func main() {
// 	n := 10
// 	m := 20 // 局部变量
// 	fmt.Println(m, n)

// }

// func foo() (int, string) {
// 	return 10, "aaf"
// }

// func main() {
// 	x, _ := foo() // 匿名变量用_表示，:=不能使用在函数之外，函数外的每个语句必须以关键字开始（var，const，func）
// 	_, y := foo()
// 	fmt.Println("x=", x)
// 	fmt.Println("y=", y)
// }

// 常量
const pi = 3.1415
const (
	pi_1 = 3
	e    = 2.7
)

// iota 常量计数器，在const出现时会被重置为0，const中每增加一行常量声明，iota计数一次
// const (
// 	n1 = iota
// 	n2
// 	_
// 	n4
// )
// const (
// 	n1   = iota
// 	kb   = 1 << (10 * iota)
// 	n2   = 100
// 	n3   = iota
// 	n4   = 123_456
// 	a, b = iota + 1, iota + 2
// )

func main() {
	// 多行字符串
	// s1 := `a
	// b
	// vc
	// `
	// traversalString()
	// changeString()
	// sqrtDemo()
	// s()
	// forDemo()
	// switchDemo()
	// gotoDemo2()
	breakDemo1()
	// fmt.Println(s1)
	// fmt.Println("%f\n", math.Pi)
	// fmt.Println("%.2f\n", math.Pi)
	// fmt.Println("str := \"c:\\Code\\lesson1\"")
}

func traversalString() {
	s := "hello你好"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%v(%c)", r, r) // 如果有中文，不能按字节遍历
	}
	fmt.Println()
}

func changeString() {
	s1 := "big"
	//强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '白'
	fmt.Println(string(runeS2))
}

func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数为float类型
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// && || ! & | ^ <<

func s() {
	a := 1
	if a == 1 {
		fmt.Println("success")
	} else if a == 2 {
		fmt.Println("aa")
	} else {
		fmt.Println("ss")
	}
}

func ifDemo2() {
	if score := 65; score >= 90 { // score仅在if语句中可见
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

// for  初始语句；条件表达式；结束语句 {}
func forDemo() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	for i := range 5 {
		fmt.Println(i)
	}
	for range 2 {
		fmt.Println("aa")
	}
}

// for 循环可以通过break goto return panic强制退出循环
// for range 可以遍历数组、切片、字符串、map 通道

// switch case
func switchDemo() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")

	case 3:
		fmt.Println("3")
	}
}

// switch分支可以用一个值，多个值，也可以用表达式
func testSwitch() {
	// switch n := 7; n {
	// case 1 ,3, 5, 7, 9:
	// 	fmt.Println("奇数")
	// case 2, 4, 6, 8:
	// 	fmt.Println("偶数")
	// default:
	// 	fmt.Println(n)
	// }

	// age := 30
	// switch {
	// case age < 25:
	// 	fmt.Println("a")
	// case age > 25 && age < 35:
	// 	fmt.Println("b")
	// case age > 60 :
	// 	fmt.Println("c")
	// default:
	// 	fmt.Println("d")
	// }
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough // 可以执行满足条件的case的下一个case
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

// goto语句可以进行代码间的无条件跳转，可以在快速跳出循环、避免重复退出上有一定的帮助
func gotoDemo1() {
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				breakFlag = true
				break
			}
			fmt.Print("%v-%v\n", i, j)
		}
		if breakFlag {
			break
		}
	}
}

func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
breakTag:
	fmt.Println("结束for循环")
}

func breakDemo1() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}