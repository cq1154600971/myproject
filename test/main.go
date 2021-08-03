package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//str := strings.ToLower(scanner.Text())
	//scanner.Scan()
	//ch := strings.ToLower(scanner.Text())
	//ans := 0
	//for i := range str {
	//	if str[i] == ch[0] {
	//		ans++
	//	}
	//}
	//fmt.Printf("%d\n", ans)
	scan := bufio.NewScanner(os.Stdin)
	for {
		scan.Scan()
		if len(scan.Text()) == 0 {
			break
		}
		count,_ := strconv.Atoi(scan.Text())
		n := [1001]bool{}
		for i := 0;i < count;i++ {
			scan.Scan()
			numstr := scan.Text()
			num,_ := strconv.Atoi(numstr)
			n[num] = true
		}

		for i,v := range n {
			if v {
				fmt.Println(i)
			}
		}
	}
}