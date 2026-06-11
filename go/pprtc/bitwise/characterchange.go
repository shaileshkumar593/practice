package main

import "fmt"

func findTheDifference(s string, t string) byte {
    var xor byte = 0

    for i := 0; i < len(s); i++ {
        xor ^= s[i]
    }

    for i := 0; i < len(t); i++ {
        xor ^= t[i]
    }

    return xor
}

func main() {
    fmt.Println(string(findTheDifference("abcd", "ceadb")))   // e
    fmt.Println(string(findTheDifference("abbdd", "dabadb"))) // a
    fmt.Println(string(findTheDifference("", "y")))           // y
    fmt.Println(string(findTheDifference("aaaa", "aabaa")))   // b
}
