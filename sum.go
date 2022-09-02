package calculator

// 在文件所在目录执行
// go mod init github.com/zhg1992/calculator

var logMessage = "[LOG]"

// Version of the calculator
var Version = "2.0"

func internalSum(number int) int {
	return number - 1
}

// Sum two integer numbers
func Sum(number1, number2 int) int {
	return number1 + number2
}
