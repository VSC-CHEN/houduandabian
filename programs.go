/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-11-16 09:16:04
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-11-17 11:20:37
 * @FilePath: \Memo_frame\Program\Programs.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */

 package Programs

 import (
	 "debt/CSV"
	 "fmt"
	 // "strconv"
 )
 
 type Debt struct {
	 Name   string
	 Borrow float64
	 Refund float64
	 Total  float64
	 Valid  string
	 key    string
	 loop   bool
 }
 
 func NewDebt() *Debt {
	 return &Debt{
		 key:    "",
		 loop:   false,
		 Name:   "",
		 Borrow: 0,
		 Refund: 0,
		 Total:  0,
		 Valid:  "y",
	 }

 }
 
 func (this *Debt) exit() {
	fmt.Println("确定退出吗?输入y/n...")
                chice := ""
                for{
                    fmt.Scanln(&chice)
                    if chice == "y" || chice == "n"{
                        break
                    }
                    fmt.Println("输入正确的选项")
                }
                if chice == "y"{
                    return
				}
 }
 
 func (this *Debt) ShowCSV() {
	 CSV.ShowCSV()
 }
 
 func (this *Debt) Borrows() {
	date :=[][] string{
		{"姓名","借款","还款","总额","有效"},
		{"张三","1000.720000","1500.720000","500.000000","y"},
		{"李四","500.000000","300.000000","-200.000000","y"},
		{"pig","500.000000","500.000000","0.000000","n"},
		{"fzf404","50.000000","30.000000","-20.000000","y"},
		{"小红","6500.000000","6500.000000","0.000000","n"},
	}
	fmt.Printf("请输入贷款人姓名")
	var n string
	var i int
	var j int
    for i := 0; i < 10; i++ {

     var tmp []string
     for j:= 0; j < 10; j++ {
		
        tmp = append(tmp, j)
		fmt.Scanf("%s",&n)
     }

     date = append(date, tmp)
	 date=n
	}
	if n==date[i][j] {
		fmt.Println("输入贷款金额")
		var sum=date[i][j+1]
		fmt.Scanf("%f",&date[i][j+1])
		sum+=date[i][j+1]
		date[i][j+1]=sum
	}
 }
 
 func (this *Debt) Refunds() {
	 date:=[][]string
	var n string
	var j int
	var i int
	fmt.Printf("请输入贷款人姓名")
	fmt.Scanf("%s",&n)
	if n==date[i][j]{
		fmt.Println("输入还债金额")
		var sum1=date[i][j+1]
		fmt.Scanf("%f",&date[i][j+1])
		sum1+=date[i][j+1]
		date[i][j+1]=sum1
	}
 }
 
 func (this *Debt) Totals() {
	func ReadCSV(date [][]string)
	date:=[][]string
	var n string
	var j int
	var i int
	fmt.Printf("请输入贷款人姓名")
	fmt.Scanf("%s",&n)
	if n==date[i][j]{
		fmt.Println("以下为您的总额")
		var date[i][j+3]string
		date[i][j+3] = date[i][j+1]+date[i][j+2]
		fmt.Printf("%f",date[i][j+3])
	}
 }
 
 func (this *Debt) IsDebt() {
	date:=[][]string
	var n string
	var j int
	var i int
	fmt.Printf("请输入贷款人姓名")
	fmt.Scanf("%s",&n)
	
	if n==date[i][j]{
		if date[i][j+3]<0{
			fmt.Printf("您负债%f元",date[i][j+3] )
		}else {
			fmt.Printf("您的账户里还剩%f",date[i][j+3])
		}
	}
 }
 
 func (this *Debt) MainMenu() {
	 for {
		 fmt.Println("------------------ 债务备忘录软件-----------------")
		 fmt.Println("                   1 借债")
		 fmt.Println("                   2 还债/存入")
		 fmt.Println("                   3 总额和(模拟后台查看记录)")
		 fmt.Println("                   4 欠债查询")
		 fmt.Println("                   5 总览(模拟后台查看记录)")
		 fmt.Println("                   6.退出系统 ")
		 fmt.Println("请选择(1-6):")
		 fmt.Println("首次借款后，下次可还款存入")
		 fmt.Scanln(&this.key)
		 switch this.key {
		 case "1":
			 this.Borrows()
		 case "2":
			 this.Refunds()
		 case "3":
			 this.Totals()
		 case "4":
			 this.IsDebt()
		 case "5":
			 this.ShowCSV()
		 case "6":
			 this.exit()

		 default:
			 fmt.Println("您的输入不正确，请按提示输入正确的选项..")
		 }
		 if this.loop {
			 break
		 }
	 }
 }
 