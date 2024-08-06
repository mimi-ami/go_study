package main

import (
	"errors"
	"fmt"
)

// error接口只包含一个方法Error
type error interface {
	Error() string // 需要返回一个描述错误信息的字符串
}

// 当一个函数或方法需要返回错误时，通常是把错误作为最后一个返回值
// func Open(name string) (*File, error) {
// 	return OpenFile(name, O_RDONLY, 0)
// }

// file, err := os.Open("./xx.go")
// if err != nil {
// 	fmt.Println("打开文件失败,err:", err)
// 	return
// }

// 自定义error
func New(text string) error

func queryById(id int64) (*Info, error) {
	if id <= 0 {
		return nil, errors.New("id 无效")
	}
}

var EOF = errors.New("EOF")

func main() {
	fmt.Errorf("查询数据库，err:%v", err)
	fmt.Errorf("查询数据库失败， err:%w", err)

}

func Unwrap(err error) error
func Is(err, target error) bool             // 判断err是否包含target
func As(err error, target interface{}) bool // 判断err是否为target类型

// 自定义结构体类型
type OpError struct {
	Op string
}

func (e *OpError) Error() string {
	return fmt.Sprintln("无权执行%s操作", e.Op)
}
