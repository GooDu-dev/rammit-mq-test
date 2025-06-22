package scripts

import "fmt"

func PrintHelloWorldNTimes(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Hello world ", i)
	}
}
