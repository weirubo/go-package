package main

import (
	"embed"
	_ "embed"
	"fmt"
)

// 方式 1
//go:embed hello.txt
// 将文件嵌入到字符串中
var s string

// 方式 2
//go:embed hello.txt
// 将文件嵌入到字节切片中
var b []byte

// 方式 3
//go:embed hello.txt
// 将一个或多个文件嵌入到文件系统中
var f embed.FS

/*
1. 指令需要导入 embed 包，字符串类型模式和字节类型的切片模式，可以使用空白导入。
2. 对于嵌入单个文件，字符串类型或字节类型的切片的变量通常是最好的，FS 类型允许嵌入文件树，例如静态 Web 服务器内容的目录。
3. FS 实现 io/fs 包的 FS 接口，因此，它可以与任何了解文件系统的包一起使用，包括 net/http, text/template, 和 html/template。
 */

/*
在变量声明上方的 //go:embed 指令，指定要嵌入的文件。
该指令必须紧接在包含单个变量声明的行之前，指令和变量声明之间仅允许空行和 // 注释行。
变量的类型必须是字符串类型，或者是字节类型的切片，或者是FS（或FS的别名）。
*/

/*
指令格式：
1. 文件系统模式允许多个路径以空格分隔，字符串类型模式和字节类型的切片模式，仅允许匹配单个文件路径。
2. 路径分隔符是正斜杠（即使在 Windows 系统上）。
3. 模式不能包含"."或".."或空路径元素，也不能以斜杠开始或结束。
4. 若要匹配当前目录中的所有内容，请使用"*"而不是"."。
 */

/*
1. 指令可用于导出变量和未导出变量，具体取决于包是否希望将数据提供给其他包。
2. 它只能在包作用域中与全局变量一起使用，而不能与局部变量一起使用。
 */
func main () {
	fmt.Println(s)
	fmt.Println(b)
	bs, err := f.ReadFile("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bs)
}
