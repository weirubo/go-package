package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("/Users/frank/GolandProjects/go-package/lesson14/file.txt")
	defer f.Close()
	r := bufio.NewReader(f)

	//p := make([]byte, 1024)
	//index, _ := r.Read(p)
	//fmt.Println(index)
	//fmt.Println(string(p[:index]))

	bs, _ := r.ReadBytes('\n')
	fmt.Println(string(bs))
}
