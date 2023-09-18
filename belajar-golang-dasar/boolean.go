// operasi untuk data boolean
//&& || !(kebalikan)
package main

import "fmt"

func main() {
	nilaiAkhir := 90
	absensi := 80

	lulusnilaiakhir := nilaiAkhir > 80
	lulusAbsensi := absensi >= 80

	lulus := lulusAbsensi && lulusnilaiakhir
	fmt.Println(lulus)
}
