package main

import "fmt"

var name,gen string="chentao","男"
var age int = 28

func main() {
   /* 这是我的第一个简单的程序 */
   fmt.Println("Hello, World!")
   //每行代表一个语句结束
   fmt.Println("第二行");
   //字符串通过+连接。   这一点与javascript一样
   fmt.Println("who is"+"chentao");

   // var 声明变量 
   var a string = "这是一个字符串"
   fmt.Println(a)

   //连续声明变量
   var b, c int = 2, 3
   fmt.Println(b,c)

   //根据值自行判断数据类型
   var e, f = 123, "hello"
   fmt.Println(e, f)

   //这种不带声明格式的只能在函数体中出现.
   //变量的初始化。变量不能重复初始化，否则会报错
   g, h := 123, "hello"
   fmt.Println(g,h)

   //声明的变量必须使用，否则会报错 declared and not used
   //j := 56
   //全局变量可以不使用
   fmt.Println(name)

}