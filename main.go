package main

import (
	"fmt"
	"strings"
)



func main(){

	greeting :="Hello world!"

	// return boolean
	contain:=strings.Contains(greeting,"Hello")
	fmt.Println(contain)

	replace:=strings.ReplaceAll(greeting,"Hello","hi")

	toUpper:=strings.ToUpper(greeting)

	// return index word in string
	index:=strings.Index(greeting,"w")

	split:=strings.Split(greeting," ")

	fmt.Println(split)

}