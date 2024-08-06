package calc

func Add(a, b int) int { // 首字母大写表示对外开放 不对外的都是小写的
	result := a + b
	return result
}

func Sec(a, b int) int {
	result := a - b
	return result
}

func Qq(a, b int) int {
	result := a * b
	return result
}

func Pp(a, b int) float32 {
	result := float32(a / b)
	return result
}
