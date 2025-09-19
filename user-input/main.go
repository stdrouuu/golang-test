package main

import "fmt"

func main() {
	var name string
	fmt.Print("Masukkan namamu: ")
	fmt.Scanln(&name)

	var age int
	fmt.Print("Masukkan Umurmu: ")
	fmt.Scanln(&age)

	var fiveYearsFromNow int = age + 5

	fmt.Println("Hallo, " + name)
	fmt.Println("Umurmu sekarang adalah " + fmt.Sprint(age) + " tahun")
	fmt.Println("Umurmu 5 tahun ke depan adalah " + fmt.Sprint(fiveYearsFromNow) + " tahun")
}
