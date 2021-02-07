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

	// 文件信息
	// fi, err := f1.Stat()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	fi, err := os.Stat("file1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// fileName := fi.Name()
	// fmt.Println(fileName)

	// fileSize := fi.Size()
	// fmt.Println(fileSize)

	// fileTime := fi.ModTime()
	// time := fileTime.Format("2006-01-02 15:04:05")
	// fmt.Println(time)

	// fm := fi.Mode()
	// fileMode := fm.String()
	// fmt.Println(fileMode)

	isDir := fi.IsDir()
	fmt.Println(isDir)
}
