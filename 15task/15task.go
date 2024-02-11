/*
var justString string
func someFunc() {
v := createHugeString(1 << 10)
justString = v[:100]
}

func main() {
someFunc()
}

*/

package main

import (
	"fmt"
	"strings"
)

func someFunc() string {
	return createHugeString(1 << 10)[:100]

}

func main() {

	fmt.Println(someFunc())

	fmt.Println(someFuncFix())
}

func createHugeString(n int) string {
	str := strings.Builder{}

	for i := 0; i < n; i++ {
		str.WriteString("а")
	}

	return str.String()

}

func someFuncFix() string {
	v := []rune(createHugeString(1 << 10))
	return string(v[:100])
}
