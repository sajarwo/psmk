package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"pustakasaya"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:pass@/sekolah")
	if err!=nil{
		panic(0)
	}
	defer db.Close()

	//db.Query("INSERT INTO siswa VALUES ( 2, 'Andi', 710, 80, 157 )")

	s := pustakasaya.Siswa{}
	db.QueryRow("SELECT nama, nis, berat, tinggi FROM siswa").Scan(&s.Nama, &s.Nis, &s.Berat, &s.Tinggi)
	fmt.Println(s)

}
