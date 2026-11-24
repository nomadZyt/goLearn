// main 包, 是go语言的入口包, 每个程序都必须有一个main包
package main

//package 包名,  不是main 的时候，go run 不能运行, 报错,package command-line-arguments is nota main package
// package aaa

// format 格式化
import "fmt"

// 函数, func ,入口函数的名字必须是main

// 全局变量， 可以在函数外部使用
var name7 string = "孙七"
var age7 int = 50

// fmt.Println(...) 是函数调用，属于运行时才会执行的语句，和 if、for、赋值一样，必须放进某个函数里。编译器看到包这一层出现语句，就会报类似这样的错：
// fmt.Println(name7, age7) // 全局变量， 可以在函数外部使用

func main() {
	// Println()：fmt 包中的一个函数，用来打印内容并换行。
	// fmt.Println("Hello, World! 我是第一课")
	fmt.Println("hello go !!")
	// 变量声明, 变量名, 变量类型, 变量值
	var name string = "张三"
	var age int = 18 // 类型推断, 根据变量值自动推断变量类型
	var name2 = "李四" // 类型推断, 根据变量值自动推断变量类型
	var age2 = 20    // 类型推断, 根据变量值自动推断变量类型

	var name3, age3 = "王五", 30 // 类型推断, 根据变量值自动推断变量类型

	fmt.Println(name, age, name2, age2)
	fmt.Println(name3, age3)

	// 自动零值
	var name4 string
	var age4 int
	fmt.Println(name4, age4)

	// 自动零值
	var name5 string
	var age5 int
	fmt.Println(name5, age5)

	// hello go !!
	// 张三 18 李四 20
	// 王五 30
	//  0

	// 简洁写法， 省略类型声明, 根据变量值自动推断变量类型
	// 只能用在函数内部， 局部变量
	name6 := "赵六"
	age6 := 40
	fmt.Println(name6, age6)
	fmt.Println(name7, age7)
}

// go run index.go 运行文件
// go build index.go 编译文件
// go build -o index index.go 编译文件并指定输出文件名
// go build -o index.exe index.go 编译文件并指定输出文件名
// go build -o index.exe index.go 编译文件并指定输出文件名

// go mod init 初始化模块
// go mod tidy 清理模块
// go mod verify 验证模块
// go mod vendor 生成vendor目录
// go mod download 下载模块
// go mod edit 编辑模块
// go mod graph 打印模块依赖图
// go mod list 打印模块列表
// go mod why 打印为什么需要模块
// go mod tidy 清理模块
// go mod vendor 生成vendor目录
// go mod download 下载模块
