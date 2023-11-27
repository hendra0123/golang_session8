package main

import (
	"calculator_project/calculator"
	"calculator_project/ui"
	"fmt"
)

func main() {
	for {
		ui.ShowMenu()

		choice := ui.GetUserChoice()

		if choice == 0 {
			fmt.Println("Terima kasih! Program berhenti.")
			break
		}

		num1, num2 := ui.GetOperands(choice) // Tambahkan argumen choice di sini

		var result float64
		var err error

		switch choice {
		case 1:
			result = calculator.Add(num1, num2)
		case 2:
			result = calculator.Subtract(num1, num2)
		case 3:
			result = calculator.Multiply(num1, num2)
		case 4:
			result, err = calculator.Divide(num1, num2)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case 5:
			result, err = calculator.Power(num1, num2)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case 6:
			result, err = calculator.SquareRoot(num1)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case 7:
			result = calculator.Sin(num1)
		case 8:
			result = calculator.Cos(num1)
		case 9:
			result = calculator.Tan(num1)
		default:
			fmt.Println("Pilihan tidak valid. Coba lagi.")
			continue
		}

		ui.DisplayResult(result)
	}
}
