package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 带缓存区的读
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("reader error")
		return
	}
	fmt.Printf("Read success. ret: %s", str)
}
