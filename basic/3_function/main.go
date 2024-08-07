package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	_ "unsafe"
)

var _ = sayHelloTo

func sayHelloTo(firstname string, lastname string) {
	fmt.Printf("my name is %s %s \n", firstname, lastname)
}

func getSum(a int, b int) int {
	return a + b
}

func getGradeAverage(a int, b int, name string) (int, string) {
	return (a + b) / 2, name
}

func totalGrade(scores []int) (int, int) {
	var total int = 0

	for _, score := range scores {
		total = total + score
	}

	return total, total / len(scores)
}

// variadic
func sumAll(numbers ...int) []int {

	var mydata []int = []int{10, 9, 8, 7, 6, 5, 4}
	mydata = append(mydata, numbers...)

	return mydata
}

func cobaSum(numbers ...int) int {
	var total int
	for _, data := range numbers {
		total = total + data
	}
	return total
}

func laporanRapot(biodata []string, totalnilai int) []byte {
	laporan := map[string]interface{}{
		"nama":  biodata[0],
		"umur":  biodata[1],
		"kelas": biodata[2],
		"detail": map[string]interface{}{
			"nilai": totalnilai,
		},
	}

	//return laporan //map[string]interface{}

	data, err := json.MarshalIndent(laporan, "", "  ")

	if err != nil {
		log.Fatal(err)
	}
	return data
}

func namaMurid(nama string, umur int, kelas string) []string {
	myslice := []string{}
	umurString := strconv.Itoa(umur)
	resultBiodata := append(myslice, nama, umurString, kelas)

	return resultBiodata
}

func itungNilai(nilai ...int) int {
	var total int = 0
	for _, data := range nilai {
		total = total + data
	}

	return total
}

//type declaration function

type Function func(string) string

// func sayHelloWithFilter(name string, filter func(string) string) {

// function as parameter

func sayHelloWithFilter(name string, filter Function) {
	filteredname := filter(name)
	fmt.Println(filteredname)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

//ANONYMOUS FUNCTION

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	fmt.Println(blacklist(name))
}

//recursive

func factorialLoop(number int) int {
	if number == 1 {
		return number
	} else {
		return number * factorialLoop(number-1)
	}
}

//CLOSURE

func closure() {
	counterParent := -1

	increment := func() {
		counterChild := -1
		counterParent = counterParent + 1
		counterChild = counterChild + 1
		fmt.Printf("counter Parent -%d counter Child -%d \n", counterParent, counterChild)
	}

	increment()
	increment()
	increment()
}

func main() {
	sayHelloTo("naufal", "andya")
	fmt.Println("LOL")

	result := getSum(12, 20)
	fmt.Println(result)

	result, name := getGradeAverage(90, 20, "naufal andya faiz")
	fmt.Println(name, result)

	//wajib buat array dulu
	var score = []int{90, 80, 50, 70, 75, 80}

	result2, _ := totalGrade(score)

	fmt.Println(result2)

	result3 := sumAll(9, 2, 5, 10)

	fmt.Println("My Data", result3)

	mydata := []int{1, 2, 4, 5, 66, 9}
	result4 := cobaSum(mydata...)
	fmt.Println(result4)

	nilaiAndya := []int{90, 88, 70, 60, 40, 90}
	result5 := laporanRapot
	result6 := result5(namaMurid("andya", 21, "3KA25"), itungNilai(nilaiAndya...))

	fmt.Println(string(result6))

	filter := spamFilter
	sayHelloWithFilter("andya", filter)
	sayHelloWithFilter("anjing", filter)

	//anonymous

	//cara 1

	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("andya", blacklist)

	//cara 2
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})

	//recursive

	fmt.Println(factorialLoop(10))

	closure()
}
