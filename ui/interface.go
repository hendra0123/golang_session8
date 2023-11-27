package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ShowMenu menampilkan menu pilihan operasi kalkulator.
func ShowMenu() {
	fmt.Println("")
	fmt.Println("Kalkulator Sederhana")
	fmt.Println("Pilih operasi:")
	fmt.Println("1. Penjumlahan (+)")
	fmt.Println("2. Pengurangan (-)")
	fmt.Println("3. Perkalian (*)")
	fmt.Println("4. Pembagian (/)")
	fmt.Println("5. Pemangkatan (^)")
	fmt.Println("6. Akar Kuadrat (√)")
	fmt.Println("7. Sinus (sin)")
	fmt.Println("8. Kosinus (cos)")
	fmt.Println("9. Tangen (tan)")
	fmt.Println("0. Keluar")
	fmt.Println("")
	fmt.Print("Masukkan pilihan Anda: ")
}

// GetUserChoice membaca pilihan operasi dari pengguna dan mengembalikannya.
func GetUserChoice() int {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		input := scanner.Text()

		choice, err := strconv.Atoi(input)
		if err == nil && choice >= 0 && choice <= 9 {
			return choice
		}
		fmt.Println("")
		fmt.Println("Masukan pilihan valid (0-6). Coba lagi.")
		fmt.Print("Masukkan pilihan Anda: ")
	}
}

// GetOperands menerima input dari pengguna dan mengembalikan dua nilai float64.
func GetOperands(choice int) (float64, float64) {
	scanner := bufio.NewScanner(os.Stdin)

	var num1, num2 float64
	var err error

	fmt.Print("Masukkan angka pertama: ")
	for {
		scanner.Scan()
		input := scanner.Text()

		num1, err = strconv.ParseFloat(input, 64)
		if err == nil {
			break
		}

		fmt.Println("Masukan harus berupa angka. Coba lagi.")
		fmt.Print("Masukkan angka pertama: ")
	}

	// Hanya minta input untuk angka kedua jika operasi memerlukan dua angka.
	if isOperationRequiringTwoOperands(choice) {
		fmt.Print("Masukkan angka kedua: ")
		for {
			scanner.Scan()
			input := scanner.Text()

			num2, err = strconv.ParseFloat(input, 64)
			if err == nil {
				break
			}

			fmt.Println("Masukan harus berupa angka. Coba lagi.")
			fmt.Print("Masukkan angka kedua: ")
		}
	}

	return num1, num2
}

// GetUserInput menerima input dari pengguna dan mengembalikan angka dan operasi yang dipilih.
func GetUserInput() (float64, float64, string) {
	ShowMenu()
	choice := GetUserChoice()

	if choice == 0 {
		fmt.Println("")
		fmt.Println("Terima kasih! Program berhenti.")
		os.Exit(0)
	}

	num1, num2 := GetOperands(choice)

	return num1, num2, strconv.Itoa(choice)
}

// DisplayResult function displays the result to the user.
func DisplayResult(result float64) {
	fmt.Printf("Result: %f\n", result)
}

// isOperationRequiringTwoOperands mengembalikan true jika operasi memerlukan dua angka, dan false jika tidak.
func isOperationRequiringTwoOperands(choice int) bool {
	switch choice {
	case 1, 2, 3, 4, 5:
		return true
	default:
		return false
	}
}
