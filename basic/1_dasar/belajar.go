package main

import (
	"encoding/json"
	"fmt"
	"log"

	// maka package unsafe buat ngilangin warn kalo gamau make func, tapi gamau diapus funcnya
	_ "unsafe"
)

var _ = typeNumber

func typeNumber() {

	// *overview

	// Menampilkan bilangan bulat
	fmt.Println("satu =", 1) // number
	// Menampilkan bilangan pecahan
	fmt.Println("satu koma 5 =", 1.5) // float

	// *Tipe Data Number
	// *int dengan nilai minimum dan maksimum untuk setiap ukuran bit
	// int8: -128 s/d 127
	var int8Example int8 = -128
	fmt.Println("int8 minimum =", int8Example)

	// int16: -32768 s/d 32767
	var int16Example int16 = 32767
	fmt.Println("int16 maksimum =", int16Example)

	// int32: -2147483648 s/d 2147483647
	var int32Example int32 = 2147483647
	fmt.Println("int32 maksimum =", int32Example)

	// int64: -9223372036854775808 s/d 9223372036854775807
	var int64Example int64 = 9223372036854775807
	fmt.Println("int64 maksimum =", int64Example)

	// *uint tanpa nilai negatif
	// uint8: 0 s/d 255
	var uint8Example uint8 = 255
	fmt.Println("uint8 maksimum =", uint8Example)

	// uint16: 0 s/d 65535
	var uint16Example uint16 = 65535
	fmt.Println("uint16 maksimum =", uint16Example)

	// uint32: 0 s/d 4294967295
	var uint32Example uint32 = 4294967295
	fmt.Println("uint32 maksimum =", uint32Example)

	// uint64: 0 s/d 18446744073709551615
	var uint64Example uint64 = 18446744073709551615
	fmt.Println("uint64 maksimum =", uint64Example)

	// *Tipe Data Float
	// float32: sekitar 1.18 x 10^-38 s/d 3.4 x 10^38
	var float32Example float32 = 3.4e+38
	fmt.Println("float32 maksimum =", float32Example)

	// float64: sekitar 2.23 x 10^-308 s/d 1.80 x 10^308
	var float64Example float64 = 1.8e+307
	fmt.Println("float64 maksimum =", float64Example)

	// *complex64 adalah float32 real dengan bagian imajiner
	var complex64Example complex64 = 1 + 2i
	fmt.Println("complex64 =", complex64Example)

	// *complex128 adalah float64 real dengan bagian imajiner
	var complex128Example complex128 = 1 + 2i
	fmt.Println("complex128 =", complex128Example)

	// *Alias
	// *byte adalah uint8
	var byteExample byte = 255
	fmt.Println("byte (uint8) =", byteExample)

	// *rune adalah int32
	var runeExample rune = 2147483647
	fmt.Println("rune (int32) =", runeExample)

	// *int minimal adalah int32
	var intExample int = 2147483647
	fmt.Println("int (minimal int32) =", intExample)

	// *uint minimal adalah uint32
	var uintExample uint = 4294967295
	fmt.Println("uint (minimal uint32) =", uintExample)
}

var _ = typeboolean

func typeboolean() {
	fmt.Println("benar", true)
	fmt.Println("false", false)
}

var _ = typeString

func typeString() {
	fmt.Println("Naufal Andya Faiz")
	fmt.Println("Naufal"+"Andya", "faiz")
	fmt.Println(len("NOPAL ANDYA"))
	fmt.Println("NOPALRUUU ANDYA"[8]) //Ke byte
}

var _ = typeVariable

func typeVariable() {
	// variabel wajib dipake, kalo function bakal di warn,

	var name string = "noparu"
	var kelas = "3KA25"

	fmt.Println(name, kelas)

	//ga harus buat var

	namasaya := "naufal andya"
	//variabel ga boleh keyword golang

	fmt.Println(namasaya)

	//value bakal berubah kalo pake =, kalo pake := & var bakal dianggap buat variabel baru
	namasaya = "wakwaw"
	fmt.Println(namasaya)

	//multiple variabel

	var (
		first_name  = "naufal"
		middle_name = "andya"
		last_name   = "faiz"
	)

	fmt.Println(first_name, middle_name, last_name)

	const kelasdia string = "3KA258" //kalo ga dipake ga masalah (tapi ada war), tapi harus ada valuenya, (immutalble)

	fmt.Println(kelasdia)
}

var _ = typeConversion

func typeConversion() {
	var nilai32 int32 = 55555
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32) //overflow bakal jadi minus balik lagi kebelakang berdasarkan kelebihannya

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var nama byte = "andya"[4]
	var e string = string(nama)

	fmt.Println(nama, e)
}

var _ = typeDeclaration

func typeDeclaration() {
	type nomorHP string //alias tipe data string

	var nomorandya nomorHP = "6289653406790"
	fmt.Println(nomorandya)

	fmt.Println(nomorHP("62895321373849"))
}

var _ = mathOperation

func mathOperation() {
	const a int = 2
	const z int = 5
	fmt.Println(a + z)

	const b uint64 = 98
	const e uint64 = 77

	fmt.Println(b + e)

	var y int32 = 7
	var o rune = 88
	result := y + o
	fmt.Println(result) //alias sama typenya sama boleh

	//gaboleh karena typenya beda walau sama2 integer
	//var u int32 = a + b - y * a

	//increment
	//tambah := tambah + b ga boleh

	var tambah uint64 = 20 //harus deklarasiin dulu
	tambah = tambah + b
	fmt.Println(tambah)

	tambah++ //bisa operator lain, liat dokumentasi
	fmt.Println(tambah)

}

var _ = typeArray

func typeArray() {
	var names [5]string

	names[0] = "andya"
	//names[7] = "fauzan" //gabisa
	names[4] = "ojan"

	fmt.Println(names)

	var datas = [3]int{
		5, 7, 8,
	}
	fmt.Println(datas)

	var strings = [7]string{
		"2", "8",
	}

	strings[5] = "7"
	fmt.Println(strings)

	var numbers = [7]uint32{
		2, 8,
	}

	numbers[5] = 12
	numbers[0] = 5 //bakal berubah

	fmt.Println(numbers)
	fmt.Println(len(numbers))

	//gatauu pasti

	var angkaAngka = [...]uint32{ //array cuma bisa dikosongin bukan diubah
		2, 8, 5, 8, 9, 1, 4,
	} //langsung kalo ...

	fmt.Println(angkaAngka)
}

var _ = typeSlice

func typeSlice() {

	// Array yang belum diinisialisasi, s1 akan berisi nilai default (kosong)
	var s1 []int
	fmt.Println(s1) // Output: []  // Slice kosong

	s2 := []int{1, 2, 3, 4}
	fmt.Println(s2) // Output: [1 2 3 4]  // Slice dengan elemen awal

	// Menambahkan elemen ke slice
	s := []int{1, 6, 64}
	s3 := append(s, 5, 7, 9, 12, 3)
	fmt.Println(s3) // Output: [1 6 64 5 7 9 12 3]  // Slice setelah elemen-elemen baru ditambahkan

	// Mengambil sub-slice
	s4 := []int{1, 6, 9, 10, 15, 20}
	subSlice0 := s4[0]
	subSlice1 := s4[1]
	subSlice2 := s4[2]
	subSlice3 := s4[3]
	subSlice4 := s4[0:1]
	subSlice5 := s4[1:2]
	subSlice6 := s4[0:3]
	subSlice7 := s4[1:4]

	fmt.Println("Ambil slice")
	fmt.Println(subSlice0) // Output: 1
	fmt.Println(subSlice1) // Output: 6
	fmt.Println(subSlice2) // Output: 9
	fmt.Println(subSlice3) // Output: 10
	fmt.Println(subSlice4) // Output: [1]  // Sub-slice dari elemen ke-0 hingga sebelum ke-1
	fmt.Println(subSlice5) // Output: [6]  // Sub-slice dari elemen ke-1 hingga sebelum ke-2
	fmt.Println(subSlice6) // Output: [1 6 9]  // Sub-slice dari elemen ke-0 hingga sebelum ke-3
	fmt.Println(subSlice7) // Output: [6 9 10]  // Sub-slice dari elemen ke-1 hingga sebelum ke-4

	// Mengambil panjang dan kapasitas slice
	s10 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(len(s10)) // Output: 7  // Panjang slice
	fmt.Println(cap(s10)) // Output: 7  // Kapasitas slice

	// Slice dari slice
	s11 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	s12 := s11[3:8]
	s13 := s12[2:4]
	s14 := s13[0:1]
	fmt.Println(s11) // Output: [1 2 3 4 5 6 7 8 9 10 11]
	fmt.Println(s12) // Output: [4 5 6 7 8]  // Sub-slice dari elemen ke-3 hingga sebelum ke-8
	fmt.Println(s13) // Output: [6 7]  // Sub-slice dari elemen ke-2 hingga sebelum ke-4 dari s12
	fmt.Println(s14) // Output: [6]  // Sub-slice dari elemen ke-0 hingga sebelum ke-1 dari s13

	// Slice kosong
	var myslice []int
	if myslice == nil {
		fmt.Println("NO VALUE IN MY SLICE 😭") // Output: NO VALUE IN MY SLICE 😭
	}

	// Menyalin slice
	s15 := []int{1, 2, 3, 19}
	s16 := make([]int, len(s15))
	fmt.Println(s16) // Output: [0 0 0 0]  // Slice dengan elemen default
	copy(s16, s15)
	fmt.Println(s16) // Output: [1 2 3 19]  // s15 disalin ke s16

	s17 := [10]int{5, 6, 7, 8, 9, 10, 11, 12}
	s17[9] = 15
	s18 := s17[9:10]
	s19 := s17[3:7]
	fmt.Println(s18) // Output: [15]  // Elemen ke-9 dari s17
	fmt.Println(s19) // Output: [8 9 10 11]  // Sub-slice dari elemen ke-3 hingga sebelum ke-7

	// Membuat slice dengan make
	s20 := make([]int, 2, 5)
	fmt.Println(s20, len(s20), cap(s20)) // Output: [0 0] 2 5  // Slice dengan panjang 2 dan kapasitas 5
	s20 = append(s20, 1, 2, 3, 4, 5, 10, 15, 20, 40)
	fmt.Println(s20, len(s20), cap(s20)) // Output: [0 0 1 2 3 4 5 10 15 20 40] 11 20  // Setelah penambahan elemen
	s21 := s20[:4]
	s22 := s20[2:]
	s23 := s20[:]
	fmt.Println(s21) // Output: [0 0 1 2]  // Sub-slice dari elemen ke-0 hingga sebelum ke-4
	fmt.Println(s22) // Output: [1 2 3 4 5 10 15 20 40]  // Sub-slice dari elemen ke-2 hingga akhir
	fmt.Println(s23) // Output: [0 0 1 2 3 4 5 10 15 20 40]  // Semua elemen slice

	// Array
	var bulan = [12]string{"january", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember"}
	myslice1 := bulan[4:10]
	myslice2 := bulan[5:7]
	myslice3 := bulan[:5]
	myslice4 := bulan[3:]
	myslice5 := bulan[:]

	fmt.Println("Array bulan:", bulan)               // Output: Array bulan: [january februari maret april mei juni juli agustus september oktober november desember]
	fmt.Println("myslice1 (bulan[4:10]):", myslice1) // Output: [mei juni juli agustus september oktober]
	fmt.Println("myslice2 (bulan[5:7]):", myslice2)  // Output: [juni juli]
	fmt.Println("myslice3 (bulan[:5]):", myslice3)   // Output: [january februari maret april mei]
	fmt.Println("myslice4 (bulan[3:]):", myslice4)   // Output: [april mei juni juli agustus september oktober november desember]
	fmt.Println("myslice5 (bulan[:]):", myslice5)    // Output: [january februari maret april mei juni juli agustus september oktober november desember]

	// Slice
	var bulan1 = []string{"january", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember"}
	bulan4 := []string{"january", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember"}

	fmt.Println("Slice bulan1:", bulan1) // Output: Slice bulan1: [january februari maret april mei juni juli agustus september oktober november desember]
	fmt.Println("Slice bulan4:", bulan4) // Output: Slice bulan4: [january februari maret april mei juni juli agustus september oktober november desember]

	// Array dengan jumlah elemen tetap
	var bulan2 = [...]string{"january", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember"}
	bulan3 := [12]string{"january", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember"}

	fmt.Println("Array bulan2:", bulan2) // Output: Array bulan2: [january februari maret april mei juni juli agustus september oktober november desember]
	fmt.Println("Array bulan3:", bulan3) // Output: Array bulan3: [january februari maret april mei juni juli agustus september oktober november desember]

	// Latihan
	fmt.Println("latihan")

	var myslice7 = []int{1, 2}                   // Slice kosong
	var myslice8 = []int{1, 2, 3, 4, 5, 6, 7, 7} // Slice dengan elemen awal

	myslice7 = append(myslice7, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(myslice7, len(myslice7), cap(myslice7)) // Output: [1 2 1 2 3 4 5 6 7 8 9] 11 12

	myslice8 = append(myslice8, 10)
	fmt.Println(myslice8, len(myslice8), cap(myslice8)) // Output: [1 2 3 4 5 6 7 7 10] 9 12

	// Menyalin slice
	var slice = []string{"h", "e", "l", "l", "o"}
	slice2 := make([]string, 6)
	copy(slice2, slice)
	fmt.Println(slice2) // Output: [h e l l o]

	// Latihan
	var slice3 = []int{3, 4, 5}
	var slice4 = []int{6, 7, 8}

	slice3 = append(slice3, slice4...)
	fmt.Println(slice3, len(slice3), cap(slice3)) // Output: [3 4 5 6 7 8] 6 6
	slice3 = append(slice3, slice4...)
	fmt.Println(slice3, len(slice3), cap(slice3)) // Output: [3 4 5 6 7 8 6 7 8] 9 12
}

var _ = typeMap

func typeMap() {
	//kaya json, key value pair

	//nested
	nama := map[string]interface{}{
		"nama": "andya",
		"ttl":  "london, 28 desember 2001",
		"detail": map[string]interface{}{
			"username": "naufalandya",
			"umur":     19,
			"akun": map[string]interface{}{
				"facebook":    "naufal andya faiz",
				"is_verified": true,
			},
		},
	}

	jsonData, err := json.MarshalIndent(nama, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(nama)
	fmt.Println(len(nama))
	// fmt.Println(jsonData.ttl)

	fmt.Println(string(jsonData))

	//biasa

	var orang map[string]string = map[string]string{}

	orang["nama"] = "Andya"
	orang["umur"] = "19" //gaboleh int karena string

	data, err := json.MarshalIndent(orang, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	buku := make(map[string]string)
	buku["author"] = "hashim"
	buku["title"] = "delulu"
	buku["wrong"] = "ups"
	buku["halaman"] = "pace"

	delete(buku, "halaman")

	fmt.Println(buku, len(buku))

	dataBuku, err := json.MarshalIndent(buku, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(dataBuku))

}

func main() {
	// place your function here to try
	typeMap()
}
