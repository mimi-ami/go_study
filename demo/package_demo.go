//包的依赖和管理
//init初始化函数的使用
// 管理工具go module的使用

// 包支持代码模块化和代码复用 内置包比如：fmt os io。。
package main

import "fmt"

func a() {
	fmt.Println("Hello World!")

}

// 自定义包 同一个文件夹下每个.go文件的第一个非注释行要声明文件归属的包

// package packagename
// 一个文件夹下面直接包含的文件只能归属一个包，同一个包的文件不能在多个文件夹下 包名为main的包是应用程序的入口包

// 标识符可见性
// 同一个包内部声明的标识符都位于同一个命名空间下 如果需要在包的外部使用包内部的标识符就需要添加包名前缀
// go中通过标识符的首字母大小写来控制标识符的对外可见或者不可见  在一个包内部只有首字母大写的标识符才是对外可见的

// 批量引入包
/* import (
	"fmt"
	"os"
) */
// 指定新的包名
// import f "fmt"

// 如果设置一个特殊的字符作为包名，这种引入方式就称为匿名引入  匿名引入的目的是为了加载这个包，从而使得这个包的资源得以初始化，被引入的包的init函数将被执行并且只执行一次
//import _ "github.com/go-sql-driver/mysql"

// 每个.go文件中都可以有无数个init的函数，没有输入参数也没有输出参数，所有在这个包中声明的函数都会被串行调用，并且仅调用一次，每个包初始化的时候都是先执行依赖的包中声明的init，在执行当前包中声明的init函数

// goproxy  主要用于设置go模块代理 使go在后续拉取模块版本时能够通过镜像站点来快速拉取
// 默认值是 https://proxy.golang.org,direct


// 使用go module引入包
