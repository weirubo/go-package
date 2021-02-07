package main

import (
	"fmt"
	"os"
)

// type File

func main () {
	// Create 创建文件，如果文件已存在，将会被截断
	// f, err := os.Create("file1.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// Open 打开文件，权限是只读
	// f, err := os.Open("file1.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// OpenFile 打开文件，如果文件不存在，可通过传递 flag 参数 os.O_CREATE 创建文件
	// f, err := os.OpenFile("file1.txt", os.O_RDWR|os.O_CREATE, 0755)
	f, err := os.OpenFile("file1.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Stat 返回文件的 FileInfo
	// fi, err := f.Stat()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fileMode := fi.Mode().String()
	// fmt.Println(fileMode)

	// Chmod 修改权限
	// err = f.Chmod(0644)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// fi2, err := f.Stat()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fileMode2 := fi2.Mode().String()
	// fmt.Println(fileMode2)

	// type File Method
	// Close 关闭文件
	defer f.Close()

	// Name 打开文件的文件名
	// name := f.Name()
	// fmt.Println(name)

	// Write 写文件
	// n, err := f.Write([]byte("hello"))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(n)

	// WriteAt 指定文件的字节偏移位置，如果文件是 O_APPEND，WriteAt 返回错误。
	// n, err := f.WriteAt([]byte("world"), 15)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(n)

	// WriteString // 写入字符串内容，而不是字节切片内容
	// n, err := f.WriteString("golang")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(n)

	// Truncate 截断文件
	// err = f.Truncate(15)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// Read 读文件 读取文件的最大长度的字节
	// b := make([]byte, 5)
	// n, err := f.Read(b)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(b)
	// fmt.Println(string(b))
	// fmt.Println(n)

	// ReadAt 从文件的字节偏移量位置开始读取最大长度的字节
	b := make([]byte, 5)
	n, err := f.ReadAt(b, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b)
	fmt.Println(n)
	fmt.Println(string(b))
}
