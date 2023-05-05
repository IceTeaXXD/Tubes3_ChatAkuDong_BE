package Algo

import (
    "strconv"
	"strings"
    "unicode"
    "math"
)

func applyOperation(a, b float64, op rune) float64 {
	/*
		I.S. Menerima 2 buah bilangan real dan 1 buah operator
		F.S. Mengembalikan hasil operasi dari kedua bilangan tersebut
	*/
    switch op {
    case '+':
        return a + b
    case '-':
        return a - b
    case '*':
        return a * b
    case '/':
        return a / b
    case '^':
        return math.Pow(a, b)
    }
    return 0
}

func Calculate(expression string) (float64, error) {
	// Inisialisai variabel
	var numStack []float64
	var opStack []rune

	// Inisialisasi map untuk menentukan prioritas operator (precedence)
	precedence := map[rune]int{
		'(': 0,
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
		'^': 3,
	}

	// Loop sepajang ekspresi
	for len(expression) > 0 {
		// Ambil token pertama dari ekspresi dan hapus token tersebut dari ekspresi
		token := expression[0]
		expression = expression[1:]

		// Percabangan untuk menentukan token
		// Operasi di dalam kurung akan dijalankan terlebih dahulu
		if token == '(' {
			// Menambahkan '(' ke stack
			opStack = append(opStack, rune(token))

		} else if token == ')' {
			for opStack[len(opStack)-1] != '(' {
				// Mengambil 2 buah angka dan 1 buah operator dari stack
				num2 := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				num1 := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				op := opStack[len(opStack)-1]
				opStack = opStack[:len(opStack)-1]
				numStack = append(numStack, applyOperation(num1, num2, op))
			}

			// Menghapus '(' dari stack
			opStack = opStack[:len(opStack)-1]

		} else if strings.ContainsRune("+-*/^", rune(token)) { // Jika token merupakan operator
			// Operasi akan dijalankan jika operator pada stack memiliki precedence yang lebih besar atau sama dengan operator token
			for len(opStack) > 0 && precedence[opStack[len(opStack)-1]] >= precedence[rune(token)] {
				// Mengambil 2 buah angka dan 1 buah operator dari stack
				num2 := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				num1 := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				op := opStack[len(opStack)-1]
				opStack = opStack[:len(opStack)-1]

				// Menghitung hasil operasi dan memasukkan hasil operasi ke stack angka
				numStack = append(numStack, applyOperation(num1, num2, op))
			}
			opStack = append(opStack, rune(token)) // Memasukkan operator token ke stack operator

		} else if unicode.IsDigit(rune(token)) || token == '.' { // Jika token merupakan angka akan di push ke stack angka
			numStr := string(token)
            for len(expression) > 0 && (unicode.IsDigit(rune(expression[0])) || expression[0] == '.') {
                numStr += string(expression[0])
                expression = expression[1:]
            }
            num, err := strconv.ParseFloat(numStr, 64)
            if err != nil {
                return 0, err
            }
            numStack = append(numStack, num)
		}
	}
	
	// Loop untuk menghitung hasil operasi yang tersisa di stack
	for len(opStack) > 0 {
		// Mengambil 2 buah angka dan 1 buah operator dari stack
		num2 := numStack[len(numStack)-1] 
		numStack = numStack[:len(numStack)-1]
		num1 := numStack[len(numStack)-1]
		numStack = numStack[:len(numStack)-1]
		op := opStack[len(opStack)-1]
		opStack = opStack[:len(opStack)-1]

		// Menghitung hasil operasi dan memasukkan hasil operasi ke stack angka
		numStack = append(numStack, applyOperation(num1, num2, op))
	}

	return numStack[0], nil
}