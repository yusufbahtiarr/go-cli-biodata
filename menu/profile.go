package menu

import (
	"bufio"
	"fmt"
	"go-cli-biodata/services"
	"go-cli-biodata/utils"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
)

var mode int = 1
var id_user int

func Profile(db *sqlx.DB) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		utils.Clear()
		fmt.Println("BIODATA")
		fmt.Println("-------------------")
		switch mode {
		case 1:
			MenuMode(db, scanner, &mode, &id_user)
		case 2:
			MenuMode(db, scanner, &mode, &id_user)
		default:
			fmt.Println("Mode tidak valid.")
			time.Sleep(time.Second)
		}
	}
}

func MenuMode(db *sqlx.DB, scanner *bufio.Scanner, mode *int, id_user *int) {
	if *mode == 1 {
		fmt.Println("1. Register")
		fmt.Println("2. Login")
	} else {
		fmt.Println("1. Data Biodata")
		fmt.Println("2. Tambah Biodata")
	}
	fmt.Println("0. Keluar")
	fmt.Println("")
	choice := utils.InputString("Pilih Menu : ")
	switch choice {
	case "1":
		if *mode == 1 {
			services.AddUser(db)
		} else {
			services.GetAllProfile(db)
		}
	case "2":
		if *mode == 1 {
			services.LoginUser(db, mode, id_user)
		} else {
			services.AddProfile(db, *id_user)
		}
	case "0":
		fmt.Printf("\nProgram diakhiri..\n")
		time.Sleep(time.Second)
		os.Exit(0)
	default:
		fmt.Println("Input tidak valid.")
		time.Sleep(time.Second)
	}
}
