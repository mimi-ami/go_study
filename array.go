package main

import (
	"fmt"
	"strings"
	// "go_study/mypackage"
)

// func main() {
// 	// mypackage.New() // 通过包来调用方法
// 	// fmt.Println("main")
// 	// var a [3]int
// 	// var b [4]int // 此时的a和b是不同的类型
// 	// 数组初始化方式
// 	// var testArray [3]int
// 	// var testArray = [...]int{1, 2} //根据初始值个数自行判断数组长度
// 	// var numArray = [3]int{1, 2}
// 	// var cityArray = [3]string{"beijing", "shanghai", "shenzhen"}
// 	// fmt.Println(testArray)
// 	// fmt.Println(numArray)
// 	// fmt.Println(cityArray)
// 	a := [...]int{1: 1, 3: 5}
// 	fmt.Println(a)
// }

// func main() {
// 	var a = [...]string{"beijing", "shanghai"}
// 	for i := 0; i < len(a); i++ {
// 		fmt.Println(a[i]) // 输出换行
// 		fmt.Print(a[i])   // 普通输出
// 		fmt.Printf(a[i])  // 格式化输出 传值的那种
// 	}
// }

// 二维数组
// func main() {
// a := [3][2]string{
// 	{"a", "b"},
// 	{"c", "d"},
// 	{"e", "f"},
// }
// fmt.Println(a)
// // range 会返回索引和对应的值
// for _, v1 := range a {
// 	for _, v2 := range v1 {
// 		fmt.Printf("%s\t", v2)
// 	}
// 	fmt.Println()
// }
// sum := 0
// a := [5]int{1, 3, 5, 7, 8}
// for _, v := range a {
// 	sum += v
// }
// fmt.Print(sum)

// 	for i := 0; i < len(a); i++ {
// 		for j := i + 1; j < len(a); j++ {
// 			if a[i]+a[j] == 8 {
// 				fmt.Println(i, j)
// 			}

// 		}
// 	}

// }

// 数组在作为参数传递时，会复制一个新的数组，因为子函数只改变副本的值，不会改变本身的值

// 切片 slice 因为数组长度也属于类型的一部分，所以如果a [3]int作为一个函数的参数，那他只能接收这样长度的参数，数组长度定义之后也无法扩展

// 切片是一个拥有相同数据类型的可变长度的序列，是基于数组类型做的一层封装，支持自动扩容
// func main() {
// var a []string
// var b = []int{}
// var c = []bool{false, true}
// fmt.Println(a)
// fmt.Println(b)
// fmt.Println(c)
// fmt.Println(a == nil)
// fmt.Println(b == nil)
// fmt.Println(c == nil)
// a := [5]int{1, 2, 3, 4, 5}
// s := a[1:3]
// fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
// s2 := s[3:4]// 虽然看起来是对s操作，但是实际是对切片s的底层数组a操作
// fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
// // a[low : high : max]
// t := a[1:3:5]
// fmt.Printf("t%v len(t):%v cap(t):%v\n", t, len(t), cap(t))
// 通过使用 make函数动态创建切片  make([]T, size, cap)

// a := make([]int, 2, 10)
// fmt.Println(a)
// fmt.Println(len(a))
// fmt.Println(cap(a))

// s := []int{1, 3, 5}
// for i := 0; i < len(s); i++ {
// 	fmt.Println(i, s[i])
// }

// for index, value := range s {
// 	fmt.Println(index, value)
// }

// var s []int
// s = append(s, 1)
// s = append(s, 2, 3, 4)
// s2 := []int{5, 6, 7}
// s = append(s, s2...)
// fmt.Println(s)
// var numSlice []int
// for i:=0; i < 10; i++ {
// 	numSlice = append(numSlice, i)
// 	fmt.Printf("%v len:%d cap:%d ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
// }
// a := []int{1, 2, 3, 4, 5}
// c := make([]int, 5, 5)
// copy(c, a)
// fmt.Println(a)
// fmt.Println(c)
// c[0] = 1000
// fmt.Println(a)
// fmt.Println(c)
// 切片中删除元素
// a := []int{30, 31, 32, 33, 34, 35}
// // 删除索引为2的元素
// a = append(a[:2], a[3:]...)
// fmt.Println(a)
// 	var a = make([]string, 5, 10)
// 	for i := 0; i < 10; i++ {
// 		a = append(a, fmt.Sprintf("%v", i))
// 	}
// 	fmt.Println(a)

// }

// 如果两个切片共享同一个数组，那么当一个切片修改过值之后，另一个切片也会受影响

// map 无序的基于key-value的数据结构，map是引用类型，必须要初始化
// map[keytype]valuetype
// func main() {
// make 函数分配内存 make(map[keyType]valueType, [cap]) cap指定容量
// scoreMap := make(map[string]int, 8)
// scoreMap["san"] = 90
// scoreMap["si"] = 100
// fmt.Println(scoreMap)
// fmt.Println(scoreMap["san"])
// fmt.Println("type of a:%T\n", scoreMap)

// userInfo := map[string]string{
// 	"username": "xiaowangzi",
// 	"password": "123456",
// }
// fmt.Println(userInfo)

// value, ok := scoreMap["san"]
// if ok {
// 	fmt.Println(value)
// }else{
// 	fmt.Println("none")
// }

// for k := range scoreMap { // 遍历scoreMap
// 	fmt.Println(k)
// }
// delete 删除键值对
// delete(scoreMap, "san")
// for k, v := range scoreMap { // 遍历scoreMap
// 	fmt.Println(k, v)
// }
// rand.Seed(time.Now().UnixNano()) //初始化随机数种子
// var scoreMap = make(map[string]int, 200)

// for i := 0; i < 100; i++ {
// 	key := fmt.Sprintf("stu%02d", i) // 生成stu开头的字符串
// 	value := rand.Intn(100) // 生成0-99之间的随机整数
// 	scoreMap[key] = value
// }
// // 取出map中的所有key存入切片keys
// var keys = make([]string, 0, 200)
// for key := range scoreMap {
// 	keys = append(keys, key)
// }
// // 对切片排序
// sort.Strings(keys)
// for _, key := range keys {
// 	fmt.Println(key, scoreMap[key])
// }

// 	var mapSlice = make([]map[string]string, 3)
// 	for index, value := range mapSlice {
// 		fmt.Printf("index:%d value:%v\n", index, value)
// 	}

// 	fmt.Println("after init")

// 	// 对切片中的map元素进行初始化
// 	mapSlice[0] = make(map[string]string, 10)
// 	mapSlice[0]["name"] = "king"
// 	mapSlice[0]["passeword"] = "123456"
// 	mapSlice[0]["address"] = "sha"
// 	for index, value := range mapSlice {
// 		fmt.Printf("index:%d value:%v\n", index, value)
// 	}

// }
// func main() {
// 	type Map map[string][]int
// 	m := make(Map)
// 	s := []int{1, 2}
// 	s = append(s, 3)
// 	fmt.Printf("%+v\n", s)
// 	m["q1mi"] = s
// 	fmt.Printf("%+v\n", m)
// 	s = append(s[:1], s[2:]...)
// 	fmt.Printf("%+v\n", s)
// 	fmt.Printf("%+v\n", m)
// }
// 写一个程序，统计一个字符串中每个单词出现的次数。比如：“how do you do"中how=1 do=2 you=1

func main() {
	para := "how do you do"
	a := strings.Split(para, " ")
	count := make(map[string]int) // 定义切片
	for _, word := range a {
		count[word]++
	}

	for word, c := range count {
		fmt.Printf("%s=%d\n", word, c)
	}
}

// %v 表示按默认格式输出 %c 表示按字符输出
