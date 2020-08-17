package main

import "fmt"

// fmt

type person struct {
	name  string
	sex   string
	age   int
	score float64
	level string
}

func main() {
	// verb 格式化动作
	// 通用
	lucy := person{
		name:  "lucy",
		sex:   "女",
		age:   17,
		score: 658.5,
		level: "b+",
	}
	// %v 值得默认格式
	fmt.Printf("lucy:%v\n", lucy)
	// %+v 类似%v，但输出结构体时会添加字段名
	fmt.Printf("lucy:%+v\n", lucy)
	// %#v 值的 Go 语法表示
	fmt.Printf("lucy:%#v\n", lucy)
	// %T 值的类型 Go 语法表示
	fmt.Printf("lucy:%T\n", lucy)
	// %% 输出百分号
	fmt.Printf("lucy:%v 100%%\n", lucy)

	// 布尔值，Go 语言中 bool 类型的变量默认值是 false
	var bool bool
	fmt.Printf("%t\n", bool)

	// 数值输出不区分有无符号，不区分长度（int8, int16），但是区分宽度和精度
	// 整数
	// %b 二进制
	fmt.Printf("lucy's age is:%b\n", lucy.age)
	// %c 该值对应的 unicode 码值
	fmt.Printf("lucy's age is:%c\n", 1)
	// %d 十进制
	fmt.Printf("lucy's age is:%d\n", lucy.age)
	// %o 八进制
	fmt.Printf("lucy's age is:%o\n", lucy.age)
	// %q 该值对应的单引号括起来的 go 语法字符字面值，必要时会采用安全的转义表示
	fmt.Printf("lucy's age is:%q\n", lucy.age)
	// %x 十六进制，使用 a-f
	fmt.Printf("lucy's age is:%x\n", lucy.age)
	// %X 十六进制，使用 A-F
	fmt.Printf("lucy's age is:%X\n", lucy.age)
	// %U Unicode 格式
	fmt.Printf("lucy's age is:%U\n", lucy.age)

	// 浮点数
	// %b 无小数部分、二进制指数的科学计数法
	fmt.Printf("lucy's score is:%b\n", lucy.score)
	// %e 科学计数法
	fmt.Printf("lucy's score is:%e\n", lucy.score)
	// %E 科学计数法
	fmt.Printf("lucy's score is:%E\n", lucy.score)
	// %f 有小数部分，但无指数部分
	fmt.Printf("lucy's score is:%f\n", lucy.score)
	// %F 等价于%f
	fmt.Printf("lucy's score is:%F\n", lucy.score)
	// %g 根据实际情况采用%e 或%f 格式
	fmt.Printf("lucy's score is:%g\n", lucy.score)
	// %G 根据实际情况采用%E 或%F格式
	fmt.Printf("lucy's score is:%G\n", lucy.score)

	// 字符串和 byte
	// %s 直接输出字符串或[]byte
	fmt.Printf("lucy's sex is:%s\n", lucy.sex)
	// %q 该值对应的双引号括起来的 go 语法字符串字面量，必要时会采用安全的转义表示
	fmt.Printf("lucy's sex is:%q\n", lucy.sex)
	// %x 每个字节用两字符十六进制数表示，a-f
	fmt.Printf("lucy's level is:%x\n", lucy.level)
	// %X 每个字节用两字符十六进制数表示，A-F
	fmt.Printf("lucy's level is:%X\n", lucy.level)

	// 指针
	// %p 内存地址，表示为十六进制，并加上前导的 0x
	fmt.Printf("lucy's sex is:%q, %p\n", lucy.sex, &lucy.sex)

	// 其它 flag
	// '+' '' '-' '#' '0'

	// fmt 包常用方法

}
