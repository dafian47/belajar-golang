package main

import (
	"belajar-golang/library"
	"fmt"
)

func main() {

	library.SayHelloGuys()
	library.IntroduceSelf("Dafian")

	var b1 = library.Book{"Sejarah Surabaya", 2001}
	fmt.Println("Judul Buku", b1.Title)
	fmt.Println("Tahun Keluar", b1.Year)

}
