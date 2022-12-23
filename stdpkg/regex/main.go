package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	simple()
	verbose()
}
func simple() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
}

func verbose() {
	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.Match([]byte("peach")))

	fmt.Println(r.FindString("peach punch"))
	fmt.Println("idx:", r.FindStringIndex("peach punch"))
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	//----------------------------------------------------------------
	r = regexp.MustCompile("p([a-z]+)ch")

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	fmt.Println(r.ReplaceAllFunc([]byte("a peach"), bytes.ToUpper))

}
