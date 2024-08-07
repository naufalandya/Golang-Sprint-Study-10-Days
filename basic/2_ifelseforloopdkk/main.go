package main

import (
	"fmt"
	_ "unsafe"
)

var _ = ifsection

func ifsection() {
	num := 10
	if num > 0 {
		fmt.Println("Positif")
	} else {
		fmt.Println("Negatif atau Nol")
	}

	num = 5
	if num%2 == 0 {
		fmt.Println("Genap")
	}

	score := 75
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

	age := 20
	if age >= 18 && age <= 25 {
		fmt.Println("Usia antara 18 dan 25")
	}

	isStudent := true
	if isStudent {
		fmt.Println("Adalah seorang siswa")
	}

	city := "Jakarta"
	if city == "Jakarta" {
		fmt.Println("Kota adalah Jakarta")
	} else {
		fmt.Println("Kota bukan Jakarta")
	}

	height := 180
	if height > 150 {
		fmt.Println("Tinggi di atas 150 cm")
	} else {
		fmt.Println("Tinggi di bawah atau sama dengan 150 cm")
	}

	number := 30
	if number > 0 {
		if number%2 == 0 {
			fmt.Println("Positif dan Genap")
		} else {
			fmt.Println("Positif dan Ganjil")
		}
	}

	if num := 9; num%3 == 0 {
		fmt.Println("Kelipatan 3")
	}

	if num := 9; num%3 == 0 {
		fmt.Println(num, "adalah kelipatan 3")
	} else {
		fmt.Println(num, "bukan kelipatan 3")
	}

	if str := "Golang"; len(str) > 5 {
		fmt.Println(str, "memiliki panjang lebih dari 5 karakter")
	} else {
		fmt.Println(str, "memiliki panjang 5 karakter atau kurang")
	}

	if arr := []int{1, 2, 3, 4, 5}; len(arr) > 4 {
		fmt.Println("Array memiliki lebih dari 4 elemen")
	} else {
		fmt.Println("Array memiliki 4 elemen atau kurang")
	}

	if m := map[string]int{"A": 1, "B": 2}; m["A"] == 1 {
		fmt.Println("Nilai A adalah 1")
	} else {
		fmt.Println("Nilai A bukan 1")
	}

	if isActive := true; isActive {
		fmt.Println("Status aktif")
	} else {
		fmt.Println("Status tidak aktif")
	}

}

var _ = switchsection

func switchsection() {
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Hari Senin")
	case "Tuesday":
		fmt.Println("Hari Selasa")
	default:
		fmt.Println("Hari lain")
	}

	day = "Saturday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Akhir pekan")
	default:
		fmt.Println("Hari kerja")
	}

	num := 3
	switch num {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	default:
		fmt.Println("Bukan 1, 2, atau 3")
	}

	isActive := true
	switch {
	case isActive:
		fmt.Println("Aktif")
	default:
		fmt.Println("Tidak aktif")
	}

	grade := "B"
	switch grade {
	case "A":
		fmt.Println("Luar Biasa")
	case "B":
		fmt.Println("Bagus")
		fallthrough
	case "C":
		fmt.Println("Cukup")
	default:
		fmt.Println("Perlu Peningkatan")
	}

	var x interface{} = 42
	switch x.(type) {
	case int:
		fmt.Println("Tipe data adalah int")
	case string:
		fmt.Println("Tipe data adalah string")
	default:
		fmt.Println("Tipe data tidak dikenal")
	}

	num = 15
	switch {
	case num%2 == 0:
		fmt.Println("Genap")
	case num%2 != 0:
		fmt.Println("Ganjil")
	}

	fruit := "Apple"
	switch fruit {
	case "Apple":
		fmt.Println("Ini adalah apel")
	case "Banana":
		fmt.Println("Ini adalah pisang")
	default:
		fmt.Println("Buah lain")
	}

	value := 3.14
	switch {
	case value < 3:
		fmt.Println("Kurang dari 3")
	case value == 3.14:
		fmt.Println("Pi")
	default:
		fmt.Println("Lebih dari 3.14")
	}

	identifier := "100"
	switch {
	case identifier == "100":
		fmt.Println("Identifikasi 100")
	case identifier == "200":
		fmt.Println("Identifikasi 200")
	default:
		fmt.Println("Identifikasi lain")
	}
}

var _ = forsection

func forsection() {
	counter1 := 1

	for counter1 <= 10 {
		fmt.Println("Perulangan ke", counter1)
		counter1 = counter1 + 1
	}

	var array_satu = [...]int{1, 2, 4, 5, 6, 7, 8}

	counter2 := 0

	fmt.Println(len(array_satu))

	for counter2 < len(array_satu) {
		fmt.Println("Perulangan ke", array_satu[counter2])
		counter2 = counter2 + 1
	}

	var array_dua = [...]string{"januari", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember"}
	var array_dua_pointer = &array_dua
	counter3 := 0

	for counter3 < len(*array_dua_pointer) {
		fmt.Println("Bulan ke -", (*array_dua_pointer)[counter3])

		if (*array_dua_pointer)[counter3] == "april" {
			fmt.Println("pergi ke pasar")
		}

		counter3++
	}

	//bisa gini juga

	for counter4 := 1; counter4 <= 10; counter4++ {
		fmt.Printf("my counter -%d \n", counter4)
	}

	//for untuk data collection

	// * array slice map

	names := []string{"naufal", "andya", "faiz"}

	for i := 0; i < len(names); i++ {
		fmt.Printf("nama saya adalah %s \n", names[i])
	}

	for _, data := range names {
		fmt.Println(data)
	}

	for i, data := range names {
		fmt.Println(i, data)
	}
}

func main() {
	forsection()
}
