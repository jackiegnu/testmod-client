package main

import "fmt"

import (
	"github.com/jackiegnu/testmod"
	v2 "github.com/jackiegnu/testmod/v2"
	v3 "github.com/jackiegnu/testmod/v3"
)

func main() {
	fmt.Println("vim-go")
	testmod.Hi("hi testmod")
	v2.Hi("hi testmod", "ch")
	v3.Hi("hi testmod", "eng")

}
