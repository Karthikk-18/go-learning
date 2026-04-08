package main

import (
	"fmt"
	"strings"
)

type Account struct {
	Balance float64
}

func main() {
	acc := Account{}
    fmt.Println(acc.deposit(1000.0))
	fmt.Println(acc.checkBalance())
	acc.withdraw(500.0)
}

func (a *Account) deposit(amount float64) string {
	a.Balance += amount
   return "Deposit successful"
}

func (a Account) checkBalance() float64{
    return a.Balance;
}

func (a *Account) withdraw(amount float64) {
    if a.Balance >= amount {
        a.Balance -= amount
		fmt.Printf("withdraw successful of %.2f, remaining amount : %.2f",amount,a.Balance)
	} else {
		fmt.Print("Insufficient balance")
	}
}

func reverse() {
	var slice = make([]int,6,10)
    fmt.Print(slice)
	fmt.Println()
	for i := 0; i < len(slice)/2; i++ {
        temp := slice[i]
		slice[i] = slice[len(slice)-1-i]
		slice[len(slice)-1-i] = temp
	}
	fmt.Print(slice, len(slice), cap(slice))
}

func even_odd() {
	var num int
    fmt.Print("Enter the number : ")
	fmt.Scanln(&num)
	if num % 2 == 0 {
		fmt.Printf("%d is even",num)
	}else {
		fmt.Printf("%d is odd",num)
	}
}

func greet() {
	var name string
    fmt.Print("Enter your name : ")
	fmt.Scanln(&name)
	var age int
    fmt.Print("Enter your age : ")
	fmt.Scanln(&age)
	fmt.Printf("Hey %s! kudos to turning %d in hits year",name,age+1)
}

func sum() {
	var num int
    fmt.Print("Enter the number : ")
	fmt.Scanln(&num)

	var sum int
	for i := 1; i <= num; i++ {
		sum += i
	}

	fmt.Printf("Sum of 1st %d is %d",num,sum)
}

func countWords() {
	str := "go is awsome go is fast"
	str1 := strings.Split(str, " ")
	freqMap := make(map[string]int)
    
	for i := 0; i < len(str1); i++ {
       if _,ok := freqMap[str1[i]]; ok{
           freqMap[str1[i]]++
	   }else {
		 freqMap[str1[i]] = 1
	   }
	}

	for key, value := range freqMap {
        fmt.Println(key,":", value)
	}
}