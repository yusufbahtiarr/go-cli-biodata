package services

import (
	"fmt"
	"go-cli-biodata/utils"
	"time"

	"github.com/jmoiron/sqlx"
)

func AddUser(db *sqlx.DB) {
	utils.Clear()
	fmt.Println("REGISTRASI USER")
	fmt.Println("-------------------")
	username := utils.InputString("Masukkan Username: ")
	password := utils.InputString("Masukkan Password: ")

	insert := `INSERT INTO users (username, password) VALUES ($1, $2)`
	_, err := db.Exec(insert, username, password)
	if err != nil {
		fmt.Printf("gagal menyimpan user: %v", err)
	}

	fmt.Printf("\nRegistrasi berhasil!\n")

	utils.GoBack()
}

func LoginUser(db *sqlx.DB, mode *int, id_user *int) {
	utils.Clear()
	fmt.Println("LOGIN USER")
	fmt.Println("-------------------")
	username := utils.InputString("Masukkan Username: ")
	password := utils.InputString("Masukkan Password: ")

	var storedPassword string
	var user_id int
	query := `SELECT id, password FROM users WHERE username = $1`
	err := db.QueryRow(query, username).Scan(&user_id, &storedPassword)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			fmt.Println("Username tidak ditemukan!")
		} else {
			fmt.Printf("Gagal memeriksa user: %v\n", err)
		}
		utils.GoBack()
		return
	}

	if password == storedPassword {
		fmt.Printf("\nLogin berhasil!\n")
		time.Sleep(time.Second * 2)
		*mode = 2
		*id_user = user_id
	} else {
		fmt.Println("Password salah!")
	}

	utils.GoBack()
}
