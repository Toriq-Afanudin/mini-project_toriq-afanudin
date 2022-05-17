package mahasiswas

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mini.com/tabels"
)

func MahasiswaAkumulasi(c *gin.Context) {
	//KONEKSI KE DATABASE
	db := c.MustGet("db").(*gorm.DB)

	var column tabels.Mahasiswa
	db.Where("nim = ?", c.Param("nim")).Find(&column)
	var tabel []tabels.Akumulasi
	db.Where("nama = ?", column.Nama).Find(&tabel)

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   tabel,
	})
}
