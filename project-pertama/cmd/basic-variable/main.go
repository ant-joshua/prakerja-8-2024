package main

import (
	"fmt"
	"slices"
)

func main() {
	// int32 = 4 bytes
	// int64 = 8 bytes

	var nilai int32 // private variable
	var name = "Joshua"
	// var jumlah int64
	// fixed size = array
	// var hobbies = [5]string{"Membaca", "Menulis", "Menggambar", "Mengolah data", "Mengolah data"}
	// dynamic size = slice
	var hobbies2 = []string{"Membaca", "Menulis", "Menggambar", "Mengolah data", "Mengolah data"}

	hobbies2 = append(hobbies2, "Mengolah data 2")

	hobbies2 = slices.Delete(hobbies2, 4, 1)

	nilai = 10
	// jumlah = 90 + int64(nilai)

	fmt.Printf("Nama saya %s, %T\n", name, "salam kenal")

	fmt.Println("Nama saya Joshua =" + name)

	fmt.Printf("Nilai saya adalah = %T", nilai)

}

// public function
func Tambah() {

}

// private function
func tambah() {

}
