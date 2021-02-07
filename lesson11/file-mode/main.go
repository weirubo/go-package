package main

import (
	"fmt"
	"os"
)

func main () {
	// 创建文件
	// f1, err := os.Create("file1.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// fmt.Println(f1)

	// 文件权限
	fi, err := os.Stat("file1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fileName := fi.Name()
	// fmt.Println(fileName)

	// fileSize := fi.Size()
	// fmt.Println(fileSize)

	// time := fi.ModTime()
	// fmt.Println(time)

	// isDir := fi.IsDir()
	// fmt.Println(isDir)

	fm := fi.Mode()
	// isDir := fm.IsDir()
	// fmt.Println(isDir)

	// 权限
	// fm2 := fm.Perm()
	// fmt.Println(fm2)

	// 权限
	str := fm.String()
	fmt.Println(str)

}
