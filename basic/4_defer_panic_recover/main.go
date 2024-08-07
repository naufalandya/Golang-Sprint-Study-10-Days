package main

import (
	"fmt"
)

func logging() {
	fmt.Println("selesai memanggil function")
}

// defer artinya nanti di akhir bakal di eksekusi kalo functionnya selesai
func runApplication() {
	defer logging()
	fmt.Println("Run application")
}

//panic memberhentikan program, tapi defer bakal tetap dieksekusi

func endApp() {
	message := recover()
	fmt.Println("run app gagal", message)
	fmt.Println("end app")
}

func runApp(error bool) {
	defer endApp() //bakal kepanggil

	if error {
		panic("ups error")
	}

	fmt.Println("run app")

}

//recover, untuk menangkap data panic, proses panic akan terhenti sehingga program akan teteap berjalan
//recover taro di defer

func main() {
	fmt.Println("Hello error handling")
	runApplication()
	runApp(true)
}
