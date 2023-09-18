//break -> menghentikan seluruh perulang
//continue -> menghentikan perulangan dan langsung melanjutkan perulangan selanjutnta

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("perulang ke", i)
	}

	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			continue //yg dibawah akan di skip, lanjut ke pperulang selanjutnya
		}
		fmt.Println("perulang ke", j)

	}
}
