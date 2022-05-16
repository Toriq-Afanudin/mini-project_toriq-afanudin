package tabels

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "root:Ayuw@hyuni1@(localhost)/sistem_presensi?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("gagal koneksi database")
	}

	// db.AutoMigrate(&Daftar_mahasiswa{})

	return db
}
