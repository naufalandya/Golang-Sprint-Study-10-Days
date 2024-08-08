package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// * struct adalah kerangka tipe data yang digabung menjadi kelompok sebagai attribut
// * yang diboleh digunakan pada suatu variabel sehingga lebih terstruktur datanya ketika ingin memberi value
// ya intinya protoype data yang ingin dibuat

// struct biasa
type Profile struct { //best practice pake pascal case
	Name, Username, Email string //bisa gabung
	Age                   int
	IsVerified            bool
	PhoneNumber           string //bisa satu2
	Country               string
	Password              string
	Hobby                 []string
	ScoreDaily            []int
	Detail                ProfileDetail
}

type ProfileDetail struct {
	TotalLike      int
	TotalPost      int
	TotalStudyTime int
	SolvedCase     int
	PreferedCase   [5]string
}

//struct method

func (profile Profile) sayHello() map[string]interface{} {

	message := map[string]interface{}{
		"nama":     "Naufal Andya",
		"umur":     profile.Age,
		"password": profile.Password,
		"detail":   profile.Detail,
	}
	return message
}

// interface

type HasName interface {
	GetName() string
}

func sayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

// interface kosong

func Ups() interface{} {
	return true
	// return "string"
	// return 1
}

// NIL

func NewMap(name string) map[string]string { //harus returnnya itu interface function map slice pointer dan channel
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

//  assertion

func random() any {
	return "OK"
}

//POINTER

// secara default mengoper value bukan reference
// artinya jika kita mengirim variabel ke function, metod, dan var lain , sebenernya yang dikirim duplikasi valuenya

type Address struct {
	Name     string
	Number   int
	City     string
	Province string
	Country  string
}

func playPointer() {
	address1 := Address{
		Name:     "Kumalajaya",
		Number:   8,
		City:     "New York City",
		Province: "West Java",
		Country:  "Thailand",
	}

	address2 := address1 //semua data di copy

	address2.City = "Washington DC" //yang satu tidak akan berubah

	fmt.Println(address1) //masih new york
	fmt.Println(address2) //beda value

	//kemampuan membuat mereferensi ke alamat memori pada suatu variabel, tanpa menduplikasi data yang sudah ada
	//pass by reference biar ga boros memori

	address3 := Address{
		Name:     "Kumalajaya",
		Number:   8,
		City:     "New York City",
		Province: "West Java",
		Country:  "Thailand",
	}

	address4 := &address3             //ambil address 3, //mengacu ke address 3
	var address5 *Address = &address3 //ga bisa poin ke address4

	address4.City = "Tokyo"

	fmt.Println(address3)
	fmt.Println(address4)
	fmt.Println(address5)

	//pake asterisk

	var data1 int = 1
	var data2 = &data1
	fmt.Println("data 1 has value", data1, "with memory address", &data1) // gabisa *data1, karena dia bukan pointer
	fmt.Println("data 2 has a value which an address (pointing to address)", data2, "with address data 2", &data2, "and the pointed value of address is", *data2)

	var p *int         //deklarasi yang benar
	fmt.Println(p, &p) //nil, alamat

	//kalo single value

	var x int = 10
	var a *int = &x //a adalah variabel pointer
	var b = &x      //a adalah variabel pointer, tapi yang lengkap *x

	// b = 2, ga boleh karena b itu pointer
	b = &x //bolehnya gini, x bisa diganti ama yang lain

	c := &x                  //a adalah variabel pointer
	fmt.Println(&x, a, b, c) //value dari variabel
	fmt.Println(&a, &b, &c)  //alamat punya masing2
	fmt.Println(*a, *b, *c)  //value dari variabel yangh di point
	// fmt.Println(*x) akan error karena dia bukan pointer, sehingga tidak ada value yang dipoint dari suatu alamat yang disimpan

	type MyInformation struct {
		Name     string
		Address  string
		Username string
		Email    string
		Age      int
	}

	andya := MyInformation{ // bisa gini
		Name:     "Naufal Andya Faiz",
		Address:  "Jalan Kemayoran 79",
		Username: "kemaraukering",
		Email:    "kalajengking5555@gmail.com",
		Age:      19,
	}

	jamal := &andya
	jamal.Name = "Jamaludin" // punya andya bakal berubah

	fmt.Println("Data Andya")
	fmt.Println(andya)
	fmt.Println(&andya) //ga akan munculin alamat
	fmt.Println("Data Jamal")
	fmt.Println(jamal)
	fmt.Println(&jamal) //alamat andya karena dia pointer, jadi bisa ngambil alamat dari variabel bertipe struct
	fmt.Println(*jamal)
	fmt.Println(jamal.Name)

	var kucing string = "kucing"          //langsung ke alamat
	var cats = []string{"cat", "persian"} //bakal &, karena ada banyak alamat
	fmt.Println(&kucing, &cats, &cats[0], &cats[1])

	//new pointer, ke data kosong

	orangKosong := new(MyInformation)
	fajar := orangKosong

	fajar.Name = "fajar"

	fmt.Println(fajar)

}

//khusus pointer ADVANCED

type MyInformation struct {
	Name     string
	Address  string
	Username string
	Email    string
	Age      int
}

func ChangeInformation(MyInformation *MyInformation) {
	MyInformation.Name = "Fajar"
}

type Man struct {
	Name string
}

// func (man Man) Married() {
// 	man.Name = "Mr ." + man.Name
// }

func (man *Man) Married() {
	man.Name = "Mr ." + man.Name
}

func main() {

	//struct biasa
	var andya Profile
	andya.Name = "Naufal Andya Faiz"
	andya.IsVerified = true
	andya.Country = "Indonesia"
	andya.Password = "ouhfuwfuowfjh94HBIUP"
	andya.Age = 21
	andya.Hobby = append(andya.Hobby, "main bola", "main sepeda", "gambar", "piano")
	andya.ScoreDaily = append(andya.ScoreDaily, 90, 80, 85, 70, 50)

	andya.Detail.PreferedCase = [5]string{"System Design", "Algorithm", "Design Patter", "SQL"}
	andya.Detail.SolvedCase = 90
	fmt.Println(andya)

	ajmal := Profile{
		Name: "Ajmal",
		Age:  21,
		Detail: ProfileDetail{
			SolvedCase: 20,
			TotalLike:  100,
		},
	}

	fmt.Println(ajmal)

	pookie := Profile{
		Age:      21,
		Password: "bg9h9gb9hg98h385",
		Detail: ProfileDetail{
			TotalLike:    500,
			PreferedCase: [5]string{"Golang", "CI/CD", "Jenkins", "Linux"},
		},
	}

	data, err := json.MarshalIndent(pookie.sayHello(), "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	//interface

	orang := Person{Name: "Andya"}
	sayHello(orang)

	animal := Animal{Name: "Kucing"}
	sayHello(animal)

	//interface kosong

	var kosong interface{} = Ups()
	fmt.Println(kosong)

	//NIL

	var mydata int //OTOMATIS 0
	fmt.Println(mydata)

	var mydatas string //otomatis kosong tapi bukan nil
	fmt.Println(mydatas)

	var mydataf float64 //otomatis nol
	fmt.Println(mydataf)

	// mynil := nil //gaboleh
	//nil cuma bisa di pointer, channel, func, interface, map, or slice type

	// var myslice = []int{} //bakal []
	var myslice []int
	fmt.Println(myslice)

	// var myarray = [12]string{} //bakal [       ]
	var myarray [12]string
	fmt.Println(myarray)

	// var mymap = map[string]string{} // bakal map[]
	var mymap map[string]interface{}
	fmt.Println(mymap)

	var myinterface interface{}
	fmt.Println(myinterface) //<nil> atau ga ada

	functionNil := NewMap("")

	if functionNil == nil {
		fmt.Println("Outputnya nil")
	} else {
		fmt.Println(functionNil)
	}
	fmt.Println(functionNil) //output map[] tapi kosong atau nil

	//assertion
	//kemampuan merubah tipe data menjadi tipe data yang diinginkan

	var result any = random()
	var resultString string = result.(string)
	fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("integer", value)
	default:
		fmt.Println("UNKNOWN")
	}

	//pointer

	playPointer()

	//pointer advanced

	// fajar := MyInformation{
	// 	Username: "naufalandya",
	// 	Email:    "andyakeren123",
	// 	Age:      21,
	// }

	// ChangeInformation(fajar)

	// fmt.Println(fajar) //tidak berubah namanya jadi fajar

	depal := MyInformation{
		Username: "fajar rusli",
		Email:    "fajar26@gmail.com",
		Age:      22,
	}

	ChangeInformation(&depal)

	fmt.Println(depal)

	nopal := Man{
		Name: "nopal",
	}

	nopal.Married()

	fmt.Println(nopal)
}
