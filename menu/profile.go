package menu

import (
	"bufio"
	"fmt"
	"go-cli-biodata/services"
	"go-cli-biodata/utils"
	"os"

	"github.com/jmoiron/sqlx"
)

func Profile(db *sqlx.DB) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		utils.Clear()
		fmt.Println("BIODATA")
		fmt.Println("-------------------")
		fmt.Println("1. Data Biodata")
		fmt.Println("2. Tambah Biodata")
		fmt.Printf("0. Keluar\n\n")
		a := utils.InputString(scanner, "Pilih Menu : ")
		switch a {
		case "1":
			services.GetAllProfile(db, scanner)
		case "2":
			services.AddProfile(db, scanner)
		case "0":
			os.Exit(0)
		default:
			fmt.Println("Input tidak valid.")
			return
		}
	}

}
