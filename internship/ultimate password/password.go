package main

import(
  "fmt"
  "time"
  "math/rand"
  )

func main() {
  var x,o int
  fmt.Println("終極密碼,請輸入1~1000內的數字")
  fmt.Scanln(&x)
  r:= rand.New(rand.NewSource(time.Now().UnixNano()))
  a:=r.Intn(1000)
  max:=1000
  min:=1
  
  for x!=a{
    o++
    if x>a{
	  max=x
      fmt.Println("請輸入",min,"~",max,"的數")
    }else{
	  min=x
      fmt.Println("請輸入",min,"~",max,"的數")
    }
    fmt.Scanln(&x)
	for x>max || x<min{
	  fmt.Println("請輸入",min,"~",max,"範圍內的數值")
	  fmt.Scanln(&x)
    }
  }
  fmt.Println("恭喜答對,正確數字為",x)
  fmt.Println("一定答錯了",o,"次")
}