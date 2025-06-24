package main

import (
	"fmt"
	"go-cli-biodata/db"
	"go-cli-biodata/menu"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	db, err := db.DBConn()
	if err != nil {
		fmt.Printf("Gagal terhubung ke database\n")
	}
	defer db.Close()
	menu.Profile(db)
}
