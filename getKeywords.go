package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func main() {

// 	root := "."
// 	filepath.Walk(root, walkfunc)
// }

func walkfunc(path string, info os.FileInfo, err error) error {

	//过滤目录
	if info.IsDir() {
		return nil
	}
	// 打印文件名
	// fmt.Println(filepath.Base(path))

	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	// 每行读取
    line := 1
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// 抓取fmt所在的行
		if strings.Contains(scanner.Text(), "fmt") {
			// fmt.Println(scanner.Text()), // 这里就可以当成字符串处理该行
            fmt.Println(f.Name(), ":", line)
        }
        line++
	}

	err = scanner.Err()
	fmt.Println(err)

	return nil
}