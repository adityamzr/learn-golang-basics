package main

import "fmt"

func main() {
// SOAL 1
// Sebuah komputer memiliki 4 lampu indikator, masing-masing berwarna merah, hijau, kuning, dan biru. Lampu-lampu tersebut dapat menyala atau mati. Diketahui bahwa:

// - Jika lampu merah menyala, maka lampu hijau dan biru pasti mati.
// - Jika lampu kuning menyala, maka lampu merah pasti mati.
// - Jika lampu hijau dan biru menyala, maka lampu kuning pasti mati.
// Berapa banyak kombinasi keadaan lampu yang mungkin terjadi?


// JAWAB
// result := 0

// for r := 0; r <= 1; r++ {
// 	for g := 0; g <= 1; g++ {
// 		for y := 0; y <= 1; y++ {
// 			for b := 0; b <= 1; b++ {
// 				// Aturan 1: Jika R menyala, G dan B harus mati
// 				if r == 1 && (g == 1 || b == 1) {
// 					continue
// 				}
// 				// Aturan 2: Jika Y menyala, R harus mati
// 				if y == 1 && r == 1 {
// 					continue
// 				}
// 				// Aturan 3: Jika G dan B menyala, Y harus mati
// 				if g == 1 && b == 1 && y == 1 {
// 					continue
// 				}

// 				result++
// 				fmt.Printf("Valid: R=%d G=%d Y=%d B=%d\n", r, g, y, b)
// 			}
// 		}
// 	}
// }

// fmt.Println("Total kombinasi valid:", result)


// SOAL 3
// Sebuah brankas memiliki kode kombinasi 4 digit. Diketahui bahwa:
// - Digit pertama tidak boleh 0.
// - Digit ketiga adalah dua kali lipat digit pertama.
// - Jumlah semua digit adalah 10.

// Berapa banyak kemungkinan kombinasi kode brankas yang ada?

// JAWAB
// result := 0

// for A := 1; A <= 9; A++ {
// 		C := 2 * A
// 		if C > 9 { continue }

// 		for B := 0; B <= 9; B++ {
// 				for D := 0; D <= 9; D++ {
// 						if A+B+C+D == 10 {
// 								result++
// 								fmt.Println(A, B, C, D) 
// 						}
// 				}
// 		}
// }
// fmt.Println("Total kombinasi:", result)



	// BUBBLESORT
	// angka := []int{2, 3, 1, 4, 5}

	// fmt.Println("Sebelum diurutkan:", angka)
	// bubbleSort(angka)
	// fmt.Println("Setelah diurutkan:", angka)

	// FAKTORIAL
	// input := 5
	
	// fmt.Println("Masukan: ", input)
	// fmt.Println("Keluaran: ", faktorial(input))

	// BINARY SEARCH
	// data := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	// target := 12

	// index := binarySearch(data, target)

	// if index != -1 {
	// 	fmt.Printf("Target ada di indeks ke %d", index)
	// } else {
	// 	fmt.Printf("Output: %d", index)
	// }
}

// func bubbleSort(arr []int) {
// 	n := len(arr)
// 	for i := 0; i < n-1; i++ {
// 		swapped := false
// 		for j := 0; j < n-i-1; j++ {
// 			if arr[j] > arr[j+1] {
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 				swapped = true
// 			}
// 		}
// 		if !swapped {
// 			break
// 		}
// 	}
// }

// func faktorial(n int) int {
// 	if(n == 0 || n == 1){
// 		return 1
// 	}

// 	return n * faktorial(n-1)
// }

// func binarySearch(arr []int, target int) int {
// 	left := 0
// 	right := len(arr) - 1

// 	for left <= right {
// 		mid := left + (right-left)/2

// 		if arr[mid] == target {
// 			return mid
// 		} else if arr[mid] < target {
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}

// 	return -1
// }