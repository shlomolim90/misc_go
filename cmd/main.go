package main

import (
	"fmt"

	"github.com/shlomolim90/misc_go"
)


func main() {
	mg, _ := misc_go.New("Test Case");
	fmt.Printf("%s\n", mg.GetInfo());
	mg, _ = misc_go.New("Test Case");
	fmt.Printf("%s\n", mg.GetInfo());
}
