package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

// 全部文件的列表
var (
    files = []*excelize.File{}
)

func main() {
RESTART:
    fmt.Print("Please enter the keyword: ")
    var keyword string
    fmt.Scan(&keyword)
    root, _ := os.Getwd()
    root = root[:len(root)-11]
    // log.Println(root)
    traverseFiles(root, keyword)
    log.Println("Finished.")
    log.Println("Traverse in recursive strategy.")

    fmt.Println("Press ctrl+C to close...")
    var b byte
    fmt.Scan(&b)
    if b == '~' {
        goto RESTART
    } else {
        return
    }
}

// 遍历文件
func traverseFiles(dir string, keyword string) {
    // log.Println("dir:", dir)
    dirFiles, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Println("err: ", err)
        return
    }
    for _, file := range dirFiles {
        // log.Println(file.Name(), file.IsDir())
        if file.Name() == ".git" {
            continue
        }
        if !file.IsDir() {
            postfixSlice := strings.Split(file.Name(), ".")
            postfix := postfixSlice[len(postfixSlice)-1]
            if postfix != "xlsx" {
                continue
            }
            excelFile, err := excelize.OpenFile(file.Name())
            if err != nil {
                log.Printf("err(%d): %v", 36, err)
            }
            // files = append(files, excelFile)
            readExcel(excelFile, keyword)
        } else {
            traverseFiles(dir+file.Name()+"\\", keyword)
        }
    }
}

// 读取excel
func readExcel(f *excelize.File, keyword string) {
    defer func() {
        // Close the spreadsheet.
        if err := f.Close(); err != nil {
            log.Println(err)
        }
    }()
    sheet := f.GetSheetName(0)
    // log.Println("sheetName", sheet)
    rows, err := f.GetRows(sheet)
    if err != nil {
        log.Println(err)
        return
    }
    for rowIdx, row := range rows {
        for colIdx, colCell := range row {
            if strings.Contains(colCell, keyword) {
                var colName string
                if colIdx < 26 {
                    colName = string(colIdx + 'A')
                } else { // rowName目前只能支持到ZZ列
                    colName = string(colIdx / 26 - 1 + 'A') + string(colIdx % 26 + 'A')
                }
                fmt.Println(sheet, ":", colName, rowIdx, f.Path)
            }
        }
    }
}
