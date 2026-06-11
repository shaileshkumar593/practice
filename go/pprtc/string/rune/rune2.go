package main

/*

bytes.Join()
Join concatenates the elements of input slice to create a new byte slice. In other words, Join concatenates the elements of s to create a new byte slice. The separator sep is placed between elements in the resulting slice:

Join(s [][]byte, sep []byte) []byte
*/

import (
	"bytes"
	"fmt"
)

func main() {
	b1 := byte('a')
	b2 := []byte("A")
	b3 := []byte{'a', 'b', 'c'}
	fmt.Printf("b1 = %c\n", b1)
	fmt.Printf("b2 = %c\n", b2)
	fmt.Printf("b3 = %s\n", b3)
	s1 := []byte("Hello")
	s2 := []byte("World")
	s3 := [][]byte{s1, s2}
	fmt.Println("len of [][]", len(s3))
	s4 := bytes.Join(s3, []byte(","))
	s5 := []byte{}
	s5 = bytes.Join(s3, []byte("--"))
	s6 := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	fmt.Printf("s1 = %s\n", s1)
	fmt.Printf("s2 = %s\n", s2)
	fmt.Printf("s3 = %s\n", s3)
	fmt.Printf("s4 = %s\n", s4)
	fmt.Printf("s5 = %s\n", s5)
	fmt.Printf("%s\n", bytes.Join(s6, []byte(", ")))
}
