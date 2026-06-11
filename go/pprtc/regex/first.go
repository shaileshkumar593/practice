package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("This is the regular expression Code ")

	regexpdigit := regexp.MustCompile(`\d+`)

	matchdigit := regexpdigit.MatchString("1")
	fmt.Println("match is digit", matchdigit)

	matchdigit = regexpdigit.MatchString("123")
	fmt.Println("match is digit", matchdigit)

	matchdigit = regexpdigit.MatchString("aaa")
	fmt.Println("match is digit", matchdigit)

	regexpstring := regexp.MustCompile(`([a-zA-Z].)+`)
	matchstring := regexpstring.MatchString("123")
	matchstring1 := regexpstring.MatchString("shailesh kumar")
	fmt.Println("match is string", matchstring)
	fmt.Println("matched spaced string", matchstring1)

	matchstring = regexpstring.MatchString("apw")
	fmt.Println("match is string", matchstring)

	regexpstringdigit := regexp.MustCompile(`[a-zA-Z0-9]+`)

	matchstring = regexpstringdigit.MatchString("apw")
	fmt.Println("match is string", matchstring)

	matchdigit = regexpstringdigit.MatchString("123")
	fmt.Println("match is digit", matchdigit)

	regexpbool := regexp.MustCompile(`[truefalse]`)
	matchtrue := regexpbool.MatchString("true")
	fmt.Println("match is true", matchtrue)

	matchfalse := regexpbool.MatchString("false")
	fmt.Println("match is false", matchfalse)

	

}
