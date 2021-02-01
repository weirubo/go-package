package main

import (
	"bytes"
	"fmt"
)

// bytes

func main () {
	// func Compare(a, b []byte) int

	// func Contains(b, subslice []byte) bool
	// 判断切片 b 是否包含子切片 subslice
	// fmt.Println(bytes.Contains([]byte("hello"), []byte("he")))
	// fmt.Println(bytes.Contains([]byte("hello"), []byte("wo")))
	// fmt.Println(bytes.Contains([]byte(""), []byte("")))

	// func ContainsAny(b []byte, chars string) bool

	// func ContainsRune(b []byte, r rune) bool

	// func Count(s, sep []byte) int
	// 判断切片 s 中有非重叠的 sep 数量，如果 sep 是一个空切片，返回 s 数量加 1。
	// fmt.Println(bytes.Count([]byte("hello"), []byte("l")))
	// fmt.Println(bytes.Count([]byte("hello"), []byte("")))

	// func Equal(a, b []byte) bool
	// 判断 a 和 b 是否长度相同，并且包含相同的字节。
	// 如果参数是 nil，等效于一个空切片
	// fmt.Println(bytes.Equal([]byte("Go"), []byte("Go")))
	// fmt.Println(bytes.Equal([]byte("Go"), []byte("Golang")))
	// fmt.Println(bytes.Equal([]byte("golang"), nil))
	// fmt.Println(bytes.Equal([]byte("golang"), []byte("")))

	// func EqualFold(s, t []byte) bool
	// 判断 s 和 t 两个字节切片是否相等
	// fmt.Println(bytes.Equal([]byte("Go"), []byte("go")))
	// fmt.Println(bytes.EqualFold([]byte("Go"), []byte("go")))

	// func Fields(s []byte) [][]byte
	// 按照一个或多个空格拆分字符串为字节切片。
	// 如果 s 仅有空格，则返回一个空切片。
	fmt.Printf("%q\n", bytes.Fields([]byte("hello world")))
	fmt.Printf("%q\n", bytes.Fields([]byte("     ")))

	// func FieldsFunc(s []byte, f func(rune) bool) [][]byte

	// func HasPrefix(s, prefix []byte) bool
	// 判断字节切片 s 是否已字节切片 prefix 开头
	// fmt.Println(bytes.HasPrefix([]byte("golang"), []byte("go")))
	// fmt.Println(bytes.HasPrefix([]byte("golang"), []byte("")))

	// func HasSuffix(s, suffix []byte) bool
	// 判断字节切片 s 是否以切片 suffix 结尾
	// fmt.Println(bytes.HasSuffix([]byte("golang"), []byte("ng")))
	// fmt.Println(bytes.HasSuffix([]byte("golang"), []byte("")))

	// func Index(s, sep []byte) int
	// 判断字节切片 sep 在字节切片 s 中首次出现的索引位置。如果不存在则返回 -1。
	// fmt.Println(bytes.Index([]byte("hello"), []byte("e")))
	// fmt.Println(bytes.Index([]byte("hello"), []byte("g")))

	// func IndexAny(s []byte, chars string) int
	// 返回 chars 字符串中任何一个字符在字节切片 s 中第一次出现的索引位置。
	// 如果 chars 字符串中的字符都没有在字节切片 s 中出现，则返回 -1。
	// fmt.Println(bytes.IndexAny([]byte("hello"), "wei"))
	// fmt.Println(bytes.IndexAny([]byte("golang"), "php"))

	// func IndexByte(b []byte, c byte) int
	// 返回字节 c 在字节切片 b 中第一次出现的索引位置。
	// 如果字节 c 在字节切片 b 中不存在，则返回 -1。
	// fmt.Println(bytes.IndexByte([]byte("golang"), byte('a')))
	// fmt.Println(bytes.IndexByte([]byte("golang"), byte('b')))

	// func IndexFunc(s []byte, f func(r rune) bool) int

	// func IndexRune(s []byte, r rune) int
	// 返回字符 r 在字节切片 s 中第一次出现的索引位置。
	// 如果字符 r 在字节切片 s 中不存在，则返回 -1。
	// fmt.Println(bytes.IndexRune([]byte("golang"), 'a'))
	// fmt.Println(bytes.IndexRune([]byte("golang"), 'b'))

	// func Join(s [][]byte, sep []byte) []byte
	// 以 sep 作为分隔符，连接所有字节切片，返回一个新切片
	// fmt.Printf("%s\n", bytes.Join([][]byte{[]byte("hello"), []byte("world")}, []byte(",")))

	// func LastIndex(s, sep []byte) int
	// 返回字节切片 sep 在字节切片 s 中最后出现的字符的索引位置。
	// 如果字节切片 sep 的所有字符在字节切片 s 中都不存在，则返回 -1。
	// fmt.Println(bytes.LastIndex([]byte("gopher golang"), []byte("go")))
	// fmt.Println(bytes.LastIndex([]byte("gopher golang"), []byte("php")))

	// func LastIndexAny(s []byte, chars string) int

	// func LastIndexByte(s []byte, c byte) int

	// func LastIndexFunc(s []byte, f func(r rune) bool) int

	// func Map(mapping func(r rune) rune, s []byte) []byte

	// func Repeat(b []byte, count int) []byte
	// 返回一个重复字节切片 b count 次的新切片。
	// 如果 count 是负数，或 字节切片 b 的长度乘以 count 的结果溢出，程序将会 panic。
	// fmt.Printf("%s\n", bytes.Repeat([]byte("go"), 3))

	// func Replace(s, old, new []byte, n int) []byte
	// 返回一个将字符切片 s 中前 n 个 old 切片替换为 new 切片的切片。
	// 如果 n 小于 0，则替换所有 old 切片。
	// fmt.Printf("%s\n", bytes.Replace([]byte("hello world"), []byte("o"), []byte("x"), 2))
	// fmt.Printf("%s\n", bytes.Replace([]byte("hello world"), []byte("o"), []byte("x"), 1))
	// fmt.Printf("%s\n", bytes.Replace([]byte("hello world"), []byte("o"), []byte("x"), -1))

	// func ReplaceAll(s, old, new []byte) []byte
	// 返回一个将字符切片 s 中的所有 old 切片替换为 new 切片的切片。
	// fmt.Printf("%s\n", bytes.ReplaceAll([]byte("go go go"), []byte("go"), []byte("golang")))

	// func Runes(s []byte) []rune

	// func Split(s, sep []byte) [][]byte
	// 将字节切片 sep 作为分割符，拆分字节切片 s。
	fmt.Printf("%q\n", bytes.Split([]byte("I am a gopher."), []byte(" ")))
	fmt.Printf("%q\n", bytes.Split([]byte("I am a gopher."), []byte("")))

	// func SplitAfter(s, sep []byte) [][]byte
	// 将字节切片 sep 作为分割符，拆分字节切片 s。
	// fmt.Printf("%q\n", bytes.SplitAfter([]byte("I,am,a,gopher"), []byte(",")))

	// func SplitAfterN(s, sep []byte, n int) [][]byte
	// 将字节切片 sep 作为分割符，拆分字节切片 s 为 n 个子切片。
	// n > 0: 拆分为最多 n 个子切片，最后一个子切片将是未分割的剩余切片。
	// n == 0: 结果是 nil，零个子切片
	// n < 0: 按照分割符拆分切片
	// fmt.Printf("%q\n", bytes.SplitAfterN([]byte("I,am,a,gopher"), []byte(","), 2))
	// fmt.Printf("%q\n", bytes.SplitAfterN([]byte("I,am,a,gopher"), []byte(","), 0))
	// fmt.Printf("%q\n", bytes.SplitAfterN([]byte("I,am,a,gopher"), []byte(","), -1))

	// func SplitN(s, sep []byte, n int) [][]byte
	// 将字节切片 sep 作为分割符，拆分切片 s 为 n 个子切片。
	// n > 0: 拆分为最多 n 个子切片，最后一个子切片将是未分割的剩余切片。
	// n == 0: 结果是 nil，零个子切片
	// n < 0: 按照分割符拆分切片
	// fmt.Printf("%q\n", bytes.SplitN([]byte("I,am,a,goper"), []byte(","), 2))
	// fmt.Printf("%q\n", bytes.SplitN([]byte("I,am,a,goper"), []byte(","), 0))
	// fmt.Printf("%q\n", bytes.SplitN([]byte("I,am,a,goper"), []byte(","), -1))

	// func Title(s []byte) []byte
	// 返回字节切片 s 中每个字符首字母改为大写的切片副本
	// fmt.Printf("%s\n", bytes.Title([]byte("php is the best language")))

	// func ToLower(s []byte) []byte
	// 返回字节切片 s 中所有字符改为小写的切片副本
	// fmt.Printf("%s\n", bytes.ToLower([]byte("PHP")))

	// func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte

	// func ToTitle(s []byte) []byte
	// 返回字节切片 s 的一个副本，所有字符改为大写字符。
	// fmt.Printf("%s\n", bytes.ToTitle([]byte("i am a gopher")))

	// func ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte

	// func ToUpper(s []byte) []byte
	// 返回字节切片 s 的一个副本，所有字符改为大写字母。
	// fmt.Printf("%s\n", bytes.ToUpper([]byte("golang")))

	// func ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte

	// func ToValidUTF8(s, replacement []byte) []byte

	// func Trim(s []byte, cutset string) []byte
	// 将字节切片 s 中开头和结尾的 cutset 字符全部清除。
	// fmt.Printf("%s\n", bytes.Trim([]byte("####gopher####"), "#"))

	// func TrimFunc(s []byte, f func(r rune) bool) []byte

	// func TrimLeft(s []byte, cutset string) []byte
	// 将字节切片 s 中开头的 cutset 字符清除。
	// fmt.Printf("%s\n", bytes.TrimLeft([]byte("####gopher####"), "#"))

	// func TrimLeftFunc(s []byte, f func(r rune) bool) []byte

	// func TrimPrefix(s, prefix []byte) []byte
	// 清除字节切片 s 中包含的 prefix 切片中的字符。
	// 如果字节切片 s 不是以 prefix 切片开头，将不做任何改变。
	// fmt.Printf("%s\n", bytes.TrimPrefix([]byte("hello,world"), []byte("hello")))
	// fmt.Printf("%s\n", bytes.TrimPrefix([]byte("hello,world"), []byte("world")))

	// func TrimRight(s []byte, cutset string) []byte
	// 将字节切片 s 中结尾的 cutset 字符清除。
	// fmt.Printf("%s\n", bytes.TrimRight([]byte("####gopher####"), "#"))

	// func TrimRightFunc(s []byte, f func(r rune) bool) []byte

	// func TrimSpace(s []byte) []byte
	// 将字节切片 s 中的所有空格清除。
	// fmt.Printf("%s\n", bytes.TrimSpace([]byte(" \t\n I am a gopher \n\t\r\n")))

	// func TrimSuffix(s, suffix []byte) []byte
	// 清除字节切片 s 中包含的 suffix 字节切片的字符。
	// 如果字节切片 s 不是以字节切片 suffix 结尾，将不做任何改变。
	// fmt.Printf("%s\n", bytes.TrimSuffix([]byte("hello,world"), []byte("world")))
	// fmt.Printf("%s\n", bytes.TrimSuffix([]byte("hello,world"), []byte("hello")))
}
