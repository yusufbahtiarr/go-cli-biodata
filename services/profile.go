package services

import (
	"bufio"
	"fmt"
	"go-cli-biodata/models"
	"go-cli-biodata/utils"

	"github.com/jmoiron/sqlx"
)

func GetAllProfile(db *sqlx.DB, scanner *bufio.Scanner) {
	utils.Clear()
	rows, err := db.Queryx("SELECT id, name, age, gender, created_at from profiles")
	if err != nil {
		fmt.Printf("Gagal menjalankan query.")
	}
	defer rows.Close()
	fmt.Println("DATA BIODATA")
	fmt.Println("-------------------")
	fmt.Println("ID - NAMA - UMUR - JENIS KELAMIN")
	for rows.Next() {
		var p models.Profile
		err := rows.StructScan(&p)
		if err != nil {
			fmt.Printf("Gagal membaca data.")
			continue
		}
		fmt.Printf("%d - %s - %d - %s\n", p.Id, p.Name, p.Age, p.Gender)
	}
	utils.GoBack(scanner)
}

func AddProfile(db *sqlx.DB, scanner *bufio.Scanner) {
	utils.Clear()
	fmt.Println("TAMBAH BIODATA")
	fmt.Println("-------------------")
	name := utils.InputString(scanner, "Masukkan Nama: ")
	age := utils.InputInt(scanner, "Masukkan Umur: ")
	gender := utils.InputGender(scanner, "Masukkan Jenis kelamin (pria/wanita): ")

	insert := `INSERT INTO profiles (name, age, gender) VALUES ($1, $2, $3)`
	_, err := db.Exec(insert, name, age, gender)
	if err != nil {
		fmt.Printf("gagal menyimpan biodata: %v", err)
	}

	DisplayAddProfile(name, age, gender)
	utils.GoBack(scanner)
}

func DisplayAddProfile(name string, age int, gender string) {
	fmt.Print("\033[2J\033[H")
	fmt.Println("BIODATA")
	fmt.Println("-------------------")
	fmt.Printf("Nama : %s\n", name)
	fmt.Printf("Umur : %d\n", age)
	fmt.Printf("Jenis Kelamin : %s\n", gender)
	fmt.Printf("\nBiodata berhasil ditambahkan\n")

}
