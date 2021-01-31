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
	fmt.Println(bytes.Contains([]byte("hello"), []byte("he")))
	fmt.Println(bytes.Contains([]byte("hello"), []byte("wo")))
	fmt.Println(bytes.Contains([]byte(""), []byte("")))

	// func ContainsAny(b []byte, chars string) bool

	// func ContainsRune(b []byte, r rune) bool

	// func Count(s, sep []byte) int
	// 判断切片 s 中有非重叠的 sep 数量，如果 sep 是一个空切片，返回 s 数量加 1。
	fmt.Println(bytes.Count([]byte("hello"), []byte("l")))
	fmt.Println(bytes.Count([]byte("hello"), []byte("")))

	// func Equal(a, b []byte) bool
	// 判断 a 和 b 是否长度相同，并且包含相同的字节。
	// 如果参数是 nil，等效于一个空切片
	fmt.Println(bytes.Equal([]byte("Go"), []byte("Go")))
	fmt.Println(bytes.Equal([]byte("Go"), []byte("Golang")))
	fmt.Println(bytes.Equal([]byte("golang"), nil))
	fmt.Println(bytes.Equal([]byte("golang"), []byte("")))

	// func EqualFold(s, t []byte) bool
	// 判断 s 和 t 两个字节切片是否相等
	fmt.Println(bytes.Equal([]byte("Go"), []byte("go")))
	fmt.Println(bytes.EqualFold([]byte("Go"), []byte("go")))

	// func Fields(s []byte) [][]byte
	// 按照一个或多个空格拆分字符串为字节切片。
	// 如果 s 仅有空格，则返回一个空切片。
	fmt.Printf("%q\n", bytes.Fields([]byte("hello world")))
	fmt.Printf("%q\n", bytes.Fields([]byte("     ")))

	// func FieldsFunc(s []byte, f func(rune) bool) [][]byte

	// func HasPrefix(s, prefix []byte) bool
	// 判断字节切片 s 是否已字节切片 prefix 开头
	fmt.Println(bytes.HasPrefix([]byte("golang"), []byte("go")))
	fmt.Println(bytes.HasPrefix([]byte("golang"), []byte("")))

	// func HasSuffix(s, suffix []byte) bool
	// 判断字节切片 s 是否以切片 suffix 结尾
	fmt.Println(bytes.HasSuffix([]byte("golang"), []byte("ng")))
	fmt.Println(bytes.HasSuffix([]byte("golang"), []byte("")))

	// func Index(s, sep []byte) int
	// 判断字节切片 sep 在字节切片 s 中首次出现的索引位置。如果不存在则返回 -1。
	fmt.Println(bytes.Index([]byte("hello"), []byte("e")))
	fmt.Println(bytes.Index([]byte("hello"), []byte("g")))

	// func IndexAny(s []byte, chars string) int

	// func IndexByte(b []byte, c byte) int

	// func IndexFunc(s []byte, f func(r rune) bool) int

	// func IndexRune(s []byte, r rune) int

	// func Join(s [][]byte, sep []byte) []byte

	// func LastIndex(s, sep []byte) int

	// func LastIndexAny(s []byte, chars string) int

	// func LastIndexByte(s []byte, c byte) int

	// func LastIndexFunc(s []byte, f func(r rune) bool) int

	// func Map(mapping func(r rune) rune, s []byte) []byte

	// func Repeat(b []byte, count int) []byte

	// func Replace(s, old, new []byte, n int) []byte

	// func ReplaceAll(s, old, new []byte) []byte

	// func Runes(s []byte) []rune

	// func Split(s, sep []byte) [][]byte

	// func SplitAfter(s, sep []byte) [][]byte

	// func SplitAfterN(s, sep []byte, n int) [][]byte

	// func SplitN(s, sep []byte, n int) [][]byte

	// func Title(s []byte) []byte

	// func ToLower(s []byte) []byte

	// func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte

	// func ToTitle(s []byte) []byte

	// func ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte

	// func ToUpper(s []byte) []byte

	// func ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte

	// func ToValidUTF8(s, replacement []byte) []byte

	// func Trim(s []byte, cutset string) []byte

	// func TrimFunc(s []byte, f func(r rune) bool) []byte

	// func TrimLeft(s []byte, cutset string) []byte

	// func TrimLeftFunc(s []byte, f func(r rune) bool) []byte

	// func TrimPrefix(s, prefix []byte) []byte

	// func TrimRight(s []byte, cutset string) []byte

	// func TrimRightFunc(s []byte, f func(r rune) bool) []byte

	// func TrimSpace(s []byte) []byte

	// func TrimSuffix(s, suffix []byte) []byte
}
