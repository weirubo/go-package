package main

import (
	"bytes"
	"fmt"
	"os"
)

// bytes.buffer
// 缓冲区是具有读取和写入方法的字节的可变大小的缓冲区。缓冲区的零值是可供使用的空缓冲区。
// Bytes() []byte
// String() string
// empty() bool
// Len() int
// Cap() int
// Truncate(n int)
// Reset()
// tryGrowByReslice(n int) (int, bool)
// grow(n int) int
// Grow(n int)
// Write(p []byte) (n int, err error)
// WriteString(s string) (n int, err error)
// ReadFrom(r io.Reader) (n int64, err error)
// WriteTo(w io.Writer) (n int64, err error)
// WriteByte(c byte) error
// WriteRune(r rune) (n int, err error)
// Read(p []byte) (n int, err error)
// Next(n int) []byte
// ReadByte() (byte, error)
// ReadRune() (r rune, size int, err error)
// UnreadRune() error
// UnreadByte() error
// ReadBytes(delim byte) (line []byte, err error)
// readSlice(delim byte) (line []byte, err error)
// ReadString(delim byte) (line string, err error)
func main () {
	// type Buffer struct {
		// contains filtered or unexported fields
	// }
	// 定义一个 bytes.Buffer 类型的变量 b
	// var b bytes.Buffer
	// b.Write([]byte("Hello "))
	// fmt.Fprintf(&b, "world")
	// b.WriteTo(os.Stdout)

	// b1 := new(bytes.Buffer)
	// b1.Write([]byte("Hello "))
	// fmt.Fprintf(b1, "world")
	// b1.WriteTo(os.Stdout)

	// func NewBuffer(buf []byte) *Buffer
	// NewBuffer 使用 buf 作为初始内容创建和初始化新缓冲区。
	// 新的缓冲区接管 buf 的所有权，调用方不应在此调用后使用 buf。
	// NewBuffer 旨在准备缓冲区以读取现有数据。
	// 它还可用于设置用于写入的内部缓冲区的初始大小。
	// 为此，buf 应具有所需的容量，但长度为零。
	// 在大多数情况下，new(Buffer)（或只是声明缓冲区变量）足以初始化缓冲区。

	// b2 := bytes.NewBuffer([]byte("Hello "))
	// fmt.Fprintf(b2, "world")
	// b2.WriteTo(os.Stdout)


	// func NewBufferString(s string) *Buffer
	// NewBufferString 使用字符串 s 作为初始内容创建和初始化新的缓冲区。
	// 它旨在准备一个缓冲区来读取现有字符串。
	// 在大多数情况下，new(Buffer)（或只是声明缓冲区变量）足以初始化缓冲区。

	// b3 := bytes.NewBufferString("Hello")
	// fmt.Fprintf(b3, "world")
	// b3.WriteTo(os.Stdout)


	// func (b *Buffer) Bytes() []byte
	// Bytes() 方法返回包含缓冲区未读部分的长度 b.Len() 的切片。
	// 切片仅在下一次缓冲区修改之前有效（也就是说，直到下一次调用读取、写入、重置或截断的方法）
	// 切片别名缓冲区内容至少到下一次缓冲区修改，因此对切片的即时更改将影响将来读取的结果。
	buf := bytes.NewBufferString("golang")
	// fmt.Fprintf(buf, "gopher")
	// buf.Grow(10)
	fmt.Printf("%q\n", buf.Bytes())
	// fmt.Println(buf.Cap())
	// fmt.Println(buf.Len())
	// fmt.Printf("%q\n", buf.Next(4))
	// fmt.Printf("%q\n", buf.Next(4))

	// n, _ := buf.Read(buf.Bytes())
	// fmt.Println(n)

	// b, _ := buf.ReadByte()
	// fmt.Printf("%q\n", b)

	// b, _ := buf.ReadBytes('p')
	// fmt.Printf("%q\n", b)

	// r, size, _ := buf.ReadRune()
	// fmt.Printf("%q %d\n", r, size)

	// l, _ := buf.ReadString('p')
	// fmt.Println(l)

	// buf.Reset()
	// fmt.Println(buf.Len())

	// s := buf.String()
	// fmt.Println(s)

	// buf.Truncate(4)
	// fmt.Printf("%q\n", buf)

	// n, _ := buf.Write([]byte("php"))
	// fmt.Println(n)
	// fmt.Printf("%q\n", buf)

	// _ = buf.WriteByte('a')
	// fmt.Printf("%q\n", buf)

	// n, _ := buf.WriteRune('编')
	// fmt.Println(n)
	// fmt.Printf("%q\n", buf)

	// n, _ := buf.WriteString("编程")
	// fmt.Println(n)
	// fmt.Printf("%q\n", buf)

	n, _ := buf.WriteTo(os.Stdout)
	fmt.Println(n)
}