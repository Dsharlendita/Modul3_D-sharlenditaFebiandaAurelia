package main

import "fmt"

func main() {
	var ca1, cb1, r1 int
	var ca2, cb2, r2 int
	var a, b int

	fmt.Scan(&ca1, &cb1, &r1)

	fmt.Scan(&ca2, &cb2, &r2)

	fmt.Scan(&a, &b)

	hasil := CheckPosisiTitik(ca1, cb1, r1, ca2, cb2, r2, a, b)
	fmt.Println(hasil)
}

func CheckPosisiTitik(ca1, cb1, r1, ca2, cb2, r2, a, b int) string {
	d1 := (a-ca1)*(a-ca1) + (b-cb1)*(b-cb1)
	d2 := (a-ca2)*(a-ca2) + (b-cb2)*(b-cb2)

	kelompok1 := d1 <= r1*r1
	kelompok2 := d2 <= r2*r2

	if kelompok1 && kelompok2 {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if kelompok1 {
		return "Titik di dalam lingkaran 1"
	} else if kelompok2 {
		return "Titik di dalam lingkaran 2"
	} else {
		return "Titik di luar lingkaran 1 dan 2"
	}
}
