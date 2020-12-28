package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func tmd52(str string) string {
	//md := []byte("123")
	//md := []byte("These pretzels are making me thirsty.")
	//md := []byte("MD5testing")
	md := []byte(str)
	md5v := md5.Sum(md)
	//fmt.Println(md5v)
	//fmt.Printf("\n%x", md5v)
	md5str := fmt.Sprintf("%x", md5v)

	return md5str
}

func md5V3(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func main() {
	str := "MD5testing"
	fmt.Println(tmd52(str))
	fmt.Println(md5V3(str))
	//tmd5()

}
