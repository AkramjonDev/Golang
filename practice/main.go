package main

import (
	"fmt"
)

// func main(){
// 	fmt.Println("Hello World practice" + " " + "Hello second")
// 	fmt.Println("Hello World second")
// 	fmt.Print("checking")
// }

// func main2(){
// 	fmt.Println("Hello first" + "Hello second")
// }

// func main(){
// 	var f1 int = 100;
// 	f1 = 200

// 	fmt.Println(f1)
// }

// func main(){
// 	// any value
// 	a := fmt.Sprintf("Meni yoshim %v da", 20)
// 	fmt.Println(a)

// 	// string version
// 	b := fmt.Sprintf("Meni %s bor", "tajribam")
// 	fmt.Println(b)

// 	// Integer version
// 	c := fmt.Sprintf("I have %d years of experience", 2)
// 	fmt.Println(c)

// 	// Float version
// 	d := fmt.Sprintf("I have %f experience", 10.1334)
// 	d2 := fmt.Sprintf("I have %.2f experience", 10.1234)

// 	fmt.Println(d, d2)
// }

// func main() {
// 	senderName := "Syl"
// 	recipient := "Kaladin"
// 	message := "The Words, Kaladin. You have to speak the Words!"

// 	fmt.Printf("%s to %s, %s\n", senderName, recipient, message)
// }

// func main() {
// 	word1 := "имя"
// 	word2 := "твое"
// 	word3 := "мне"
// 	word4 := "знакомо"

// 	fmt.Println(word4, word3, word2, word1)
// 	fmt.Println(word3, word1, word4, word2)
// 	fmt.Println(word2, word4, word1, word3)
// }

// func main(){
// 	var (
// 		name string = "akmal"
// 		age int = 13
// 	)

// 	fmt.Printf("My name %s, and my age is %d", name, age)

// }

// func main(){
// var (
// name string = "Аркадий"
// surname string = "Петров"
// age int = 19
// )
// var (
// 	name2 string = "Елена"
// 	surname2 string = "Сидорова"
// 	age2 int = 19
// )

// fmt.Printf("Имя: %s, Фамилия: %s, Возраст: %d, Студент BPS", name, surname, age)
// fmt.Printf("Имя: %s, Фамилия: %s, Возраст: %d, Студент BPS", name2, surname2, age2)

// }

// func main() {
// 	aBoolean, bBoolean := true, true
// 	fmt.Println("AND", aBoolean && bBoolean)
// 	fmt.Println("OR", aBoolean || bBoolean)
// 	fmt.Println("NOT:", !aBoolean)
// }

func main(){
	var integer1 int
	var integer2 int
	var integer3 int

	fmt.Scan(&integer1)
	fmt.Scan(&integer2)
	fmt.Scan(&integer3)

	var  largest int

	 if integer1 >= integer2 && integer1 >= integer3 {
        largest = integer1	
    } else if integer2 >= integer1 && integer2 >= integer3 {
        largest = integer2
    } else {
        largest = integer3
    }

    fmt.Printf("The largest number among %d, %d, and %d is %d\n", integer1, integer2, integer3, largest)
}
