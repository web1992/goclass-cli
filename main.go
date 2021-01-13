package main

import (
	"fmt"
	"github.com/web1992/goclass/core"
	"github.com/web1992/goclass/parse"
)

var file = "/Users/zl/Documents/DEV/github/goclass-cli/testfile/GoClass.class"

func main() {

	fmt.Printf("goclass-cli \n")
	a := []byte{0, 32, 0, 0}
	fmt.Println(core.Byte2U4(a))

	var cp parse.ClassParse
	err := cp.Parse(file)

	if err != nil {
		fmt.Print(err)
	}
	cf := cp.ClassFile()
	fmt.Println(cf)

}
