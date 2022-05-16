package controls

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mini.com/tabels"
)

func DosenAkumulasi(c *gin.Context) {
	//KONEKSI KE DATABASE
	db := c.MustGet("db").(*gorm.DB)

	var column tabels.Dosen
	db.Where("nip = ?", c.Param("nip")).Find(&column)
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
	db.Where("dosen = ?", column.Gelar).Where("matakuliah = ?", m.Matakuliah).Find(&column2)
	if m.Matakuliah != column2.Matakuliah {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "saudara tidak mengampu matakuliah ini",
		})
		return
	}
	var tabel []tabels.Akumulasi
	db.Where("matakuliah = ?", column2.Matakuliah).Find(&tabel)

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   tabel,
	})
}
