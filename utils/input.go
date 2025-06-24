package utils

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func InputString(scanner *bufio.Scanner, label string) string {
	for {
		fmt.Print(label)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input != "" {
			return input
		}

		fmt.Println("Input tidak boleh kosong")
	}
}

func InputInt(scanner *bufio.Scanner, label string) int {
	for {
		fmt.Print(label)
		scanner.Scan()
		input := scanner.Text()

		value, err := strconv.Atoi(input)
		if err == nil {
			return value
		}

		fmt.Println("Input tidak valid. Harap masukkan angka.")
	}
}

func InputGender(scanner *bufio.Scanner, label string) string {
	for {
		fmt.Print(label)
		scanner.Scan()
		gender := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if gender == "pria" || gender == "wanita" {
			return gender
		}

		fmt.Println("Inputan tidak valid. Jenis Kelamin hanya boleh 'pria' atau 'wanita'")
	}
}

func GoBack(scanner *bufio.Scanner) {
	fmt.Print("\nTekan Enter untuk kembali... ")
	scanner.Scan()
}
