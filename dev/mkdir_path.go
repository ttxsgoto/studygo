package main

import (
	"fmt"
	"os"
)

func Pathexists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	_dir := "./testabc"
	exist, err := Pathexists(_dir)
	if err != nil {
		fmt.Printf("get dir error %v\n", err)
		return
	}

	if exist {
		fmt.Println("get dir is exist")
	} else {
		err := os.Mkdir(_dir, os.ModePerm)
		if err != nil {
			fmt.Println("mkdir faild")
		} else {
			fmt.Println("mkdir ok.")
		}
	}

}
