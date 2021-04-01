package main

import (
	"fmt"
	"regexp"
)

func get_schema(input string) string {

	r, _ := regexp.Compile("piintu-([a-zA-Z]\\w+)")

	return r.FindStringSubmatch(input)[1]
}

func main() {
	fmt.Println(get_schema("<i piintu-root>Hello</i>"))
	fmt.Println(get_schema("<div><div piintu-rootbear title=”Oh My”>Hello <i sc-org>World</i></div></div>"))
}
