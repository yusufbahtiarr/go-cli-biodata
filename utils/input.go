package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputString(label string) string {
	scanner := bufio.NewScanner(os.Stdin)
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

func InputInt(label string) int {
	scanner := bufio.NewScanner(os.Stdin)
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

func InputGender(label string) string {
	scanner := bufio.NewScanner(os.Stdin)
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

func GoBack() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nTekan Enter untuk kembali... ")
	scanner.Scan()
}
