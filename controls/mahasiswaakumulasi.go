package controls

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
	type matakuliah struct {
		Matakuliah string `json:"matakuliah"`
	}
	var m matakuliah
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "input tidak dalam bentuk json",
		})
		return
	}
	var column2 tabels.Krs
	db.Where("nama = ?", column.Nama).Where("matakuliah = ?", m.Matakuliah).Find(&column2)
	if m.Matakuliah != column2.Matakuliah {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "anda tidak mengambil matakuliah ini",
		})
		return
	}
	var tabel []tabels.Akumulasi
	db.Where("nama = ?", column.Nama).Find(&tabel)

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   tabel,
	})
}
