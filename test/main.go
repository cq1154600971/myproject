package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := strings.ToLower(scanner.Text())
	scanner.Scan()
	ch := strings.ToLower(scanner.Text())
	ans := 0
	for i := range str {
		if str[i] == ch[0] {
			ans++
		}
	}
	fmt.Printf("%d\n", ans)
}