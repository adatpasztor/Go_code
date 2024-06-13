package main

import "fmt"
import "example/user/hello/morestrings"
import "github.com/google/go-cmp/cmp"


func main() {
    fmt.Println(morestrings.ReverseRunes("Hello, world + PSX"))
	fmt.Println(cmp.Diff("Hello, world + PSX", "Hello Go"))
}