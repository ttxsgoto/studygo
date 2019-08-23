package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	Chcount    int
	Numcount   int
	Spacecount int
	OtherCount int
}

func main() {
	// 从文件中缓存读取数据,统计英文，数字和空格等字符
	file, err := os.Open("./bufferio.go")
	if err != nil {
		fmt.Println("reader error")
		return
	}
	defer file.Close()

	var count CharCount

	reader := bufio.NewReader(file)
	for { // 读取所有文件
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("reader error", err)
			return
		}
		// fmt.Printf("Read success. ret: %s", str)
		runeArr := []rune(str)
		for _, v := range runeArr {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.Chcount++
			case v == ' ' || v == '\t':
				count.Spacecount++
			case v >= '0' && v <= '9':
				count.Numcount++
			default:
				count.OtherCount++
			}
		}
		fmt.Printf("For Char count:%d\n", count.Chcount)
	}
	fmt.Printf("Char count:%d\n", count.Chcount)

}
