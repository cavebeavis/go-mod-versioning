package main

import (
	"fmt"

	hell1 "github.com/cavebeavis/go-mod-versioning/pkgs/hello"
	hell2 "github.com/cavebeavis/go-mod-versioning/pkgs/hello/v2"
	
	yo1 "github.com/cavebeavis/go-mod-yo"
	yo2 "github.com/cavebeavis/go-mod-yo/v2"

)

func main() {
	fmt.Println(hell1.Hello())
	fmt.Println(hell2.Hello())

	fmt.Printf("\n%s\n\n", yo1.Yo())
	fmt.Printf("\n%s\n\n", yo2.Yo())
}

