/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-11-16 09:37:46
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-11-17 10:12:46
 * @FilePath: \Memo_frame\CSV\Csv.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */

package CSV

import (
	"encoding/csv"
	"fmt"
	// "fmt"
	"log"
	"os"
)
 
 func WriteCSV(date [][]string) {
	 //创建csv文件
	 f, err := os.OpenFile("Debt.csv", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	 if err != nil {
		 panic(err)
	 }
	 //异步管理
	 defer f.Close()
	 // 写入UTF-8 BOM
	 f.WriteString("\xEF\xBB\xBF")
	 //创建一个新的写入文件流
	 w := csv.NewWriter(f)
	 data := [][]string {
		{"姓名","借款","还款","总额","有效"},
		{"张三","1000.720000","1500.720000","500.000000","y"},
		{"李四","500.000000","300.000000","-200.000000","y"},
		{"pig","500.000000","500.000000","0.000000","n"},
		{"fzf404","50.000000","30.000000","-20.000000","y"},
		{"小红","6500.000000","6500.000000","0.000000","n"},
	}
	 //写入数据
	 w.WriteAll(date)
	 w.Flush()
 }
 
 func ReadCSV(date [][]string) {
	 fileName := "Debt.csv"
	 fs1, _ := os.Open(fileName)
	 r1 := csv.NewReader(fs1)
	 content, err := r1.ReadAll()
	 if err != nil {
		 log.Fatalf("can not readall, err is %+v", err)
	 }
	 for _, row := range content {
		fmt.Println(row)
	}
	 fmt.Println("\n---------------------------\n")
	 return content
 }
 
 func ShowCSV() {
	 fileName := "Debt.csv"
	 fs1, _ := os.Open(fileName)
	 r1 := csv.NewReader(fs1)
	 content, err := r1.ReadAll()
	 if err != nil {
		 log.Fatalf("can not readall, err is %+v", err)
	 }
	 fmt.Printfln("Debt.csv")
 }
 