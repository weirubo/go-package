package main

import (
	"io"
	"log"
	"os"
)

func main() {
	//r := strings.NewReader("some io.Reader stream to be read\n")
	//if _, err := io.Copy(os.Stdout, r); err != nil {
	//	log.Fatal(err)
	//}

	//r := bytes.NewReader([]byte("some io.Reader stream to be read"))
	//if _, err := io.Copy(os.Stdout, r); err != nil {
	//	log.Fatal(err)
	//}

	//rd := strings.NewReader("some io.Reader stream to be read")
	//r := bufio.NewReader(rd)
	//if _, err := io.Copy(os.Stdout, r); err != nil {
	//	log.Fatal(err)
	//}

	f, _ := os.Open("service.log")
	if _, err := io.Copy(os.Stdout, f); err != nil {
		log.Fatal(err)
	}
}
