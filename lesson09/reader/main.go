package main

import (
	"bytes"
	"os"
)

// bytes.Reader
// type Reader struct {
// 	s        []byte
// 	i        int64
// 	prevRune int
// }

func main () {
	// r := bytes.Reader{}

	r := bytes.NewReader([]byte("hello"))
	// fmt.Println(r.Len())
	// n, _ := r.Read([]byte("world"))
	// fmt.Println(n)

	// n, _ := r.ReadAt([]byte("golang"), 2)
	// fmt.Println(n)

	// b, _ := r.ReadByte()
	// fmt.Printf("%q\n", b)

	// ch, size, _ := r.ReadRune()
	// fmt.Printf("%q %d\n", ch, size)

	r.WriteTo(os.Stdout)
}
