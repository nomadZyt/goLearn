package main

import "fmt"

func dataType() {
	fmt.Println("Hello, World!")
	var name string = "张三"
	var age int = 18
	var age1 = 18 // 类型推断, 根据变量值自动推断变量类型, 自动推断为int类型的位数
	// 有符号整数, 范围 -128 ~ 127, 有符号位, 第一位为符号位, 0为正数, 1为负数
	var age2 int8 = -18
	var age3 int8 = 18

	// 无符号整数, 范围 0 ~ 255, 无符号位, 第一位为数据位, 0为正数, 1为负数
	var age4 uint8 = 18
	var year uint = 2026     // 类型推断, 根据变量值自动推断变量类型, 自动推断为uint类型的位数
	var year1 uintptr = 2026 // 指针, 指向内存地址

	// 浮点数
	var age5 float32 = 18.5 // 单精度浮点数, 范围 1.2e-38 ~ 3.4e+38, 精度为7位小数
	var age6 float64 = 18.5 // 双精度浮点数, 范围 2.3e-308 ~ 1.7e+308, 精度为15位小数

	var pi float32 = 3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679 // 圆周率, 精度为32位小数

	fmt.Println(pi) // 会被截断为3.141593，精度为7位小数

	var pi1 float64 = 3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679 // 圆周率, 精度为64位小数

	fmt.Println(pi1)
	// 布尔值
	var age7 bool = true
	var age8 bool = false

	fmt.Println(name, age, age1, age2, age3, age4, age5, age6, age7, age8)
	fmt.Println(year, year1)

	var result = age + age1
	// 类型不匹配, 报错, 需要将age转换为float32类型
	// var result1 = age - age5
	var result1 = float32(age) - age5
	fmt.Println(result)
	fmt.Println(result1)
}
