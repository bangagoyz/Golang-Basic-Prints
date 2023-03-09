package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	var base8 int64 = 25
	var base16 string = "f"
	var base162 string = "f13"
	var angka int = 21
	var j bool = true
	ruski := "Я"
	oktal := strconv.FormatInt(base8, 8)
	hexToDes, rusak := strconv.ParseInt(base16, 16, 64)
	if rusak != nil {
		panic(rusak)
	}
	hexToDes2, err := strconv.ParseInt(base162, 16, 64)
	if err != nil {
		panic(err)
	}
	ubah, ukuran := utf8.DecodeRuneInString(ruski)
	var pecah float64 = 132.456000

	fmt.Println(angka)
	fmt.Printf("%T\n", angka)
	fmt.Println("%")
	fmt.Println(j)
	fmt.Println("\u042F")
	fmt.Println(angka)
	fmt.Println("bilangan base 8 adalah :" + oktal)
	fmt.Println("Nilai desimal dari ", base16, "adalah ", hexToDes)
	fmt.Println("Nilai desimal dari ", base162, "adalah ", hexToDes2)
	fmt.Println("Kode dari simbol : ", ruski, "adalah ", ubah, ukuran)
	fmt.Printf("%f\n", pecah)
	fmt.Printf("%e", pecah)
}
