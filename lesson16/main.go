package main

import (
	"errors"
	"fmt"
)

// errors

func main () {
	// 定义错误的两种方法：
	// New
	// 返回一个错误，错误格式为给定文本。
	err := errors.New("this is a error example")
	// 不透明错误处理
	// 错误处理方不关心错误提供方的错误值的上下文
	if err != nil {
		fmt.Printf("%p\n", err)
		fmt.Println(err)
	}

	// 即使文本相同，每次对New的调用也会返回不同的错误值。
	// 通过查阅源码可以发现，原因是返回的是 errorString 的指针类型
	err2 := errors.New("this is a error example")
	if err2 != nil {
		fmt.Printf("%p %T\n", err2, err2)
		fmt.Println(err2)
	}

	err3 := fmt.Errorf("err3:%s\n", "this is another error example")
	fmt.Println(err3)

	err4 := errors.New("包装错误")
	wrapErr := fmt.Errorf("wrap error: %w\n", err4)
	fmt.Println(wrapErr)

	// golang 1.13 新增函数
	// Unwrap

	// As
	// 判断一个 error 类型的变量是否为特定的自定义错误类型，error 类型的变量的底层错误值是一个包装错误（wrap error），与错误链（error chain）上的所有被包装的错误（wrapped error）的类型做比较，直到找到一个匹配的错误类型。
	// wrap error
	err7 := &DefineError{"this is a define error type"}
	err8 := fmt.Errorf("wrap err8: %w\n", err7)
	err9 := fmt.Errorf("wrap err9: %w\n", err8)
	var err10 *DefineError
	if errors.As(err9, &err10) {
		// errors.As() 顺着错误链，从 err9 一直找到被包装最底层的错误值 err7，并且将 err9 与其自定义类型 `var err10 *DefineError` 匹配成功。
		fmt.Println("err7 is a variable of the DefineError type")
		fmt.Println(err10 == err7)
		return
	}
	fmt.Println("err7 is not a variable of the DefineError type")

	// Is
	b := errors.Is(wrapErr, err4)
	fmt.Println(b)

	// 如果 error 类型变量的底层错误值是一个包装错误 （wrap error），在 golang 1.13 版本开始，errors.Is() 方法会顺着错误链 （error chain），与 error chain 上所有被包装的错误（wrapped error）的类型做比较，直到找到一个匹配的错误。
	err5 := fmt.Errorf("wrap err5: %w\n", ErrInvalidUser)
	err6 := fmt.Errorf("wrap err6: %w\n", err5)
	if errors.Is(err6, ErrInvalidUser) {
		fmt.Println(ErrInvalidUser)
		return
	}
	fmt.Println("success")
}

// 哨兵错误处理
var (
	ErrInvalidUser     = errors.New("invalid user")
	ErrNotFoundUser    = errors.New("not found user")
)

// 自定义的错误类型
type DefineError struct {
	msg string
}

func (d *DefineError) Error() string {
	return d.msg
}