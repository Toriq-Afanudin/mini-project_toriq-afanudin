package dosens

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
	var tabel []tabels.Akumulasi
	db.Where("dosen = ?", column.Id).Find(&tabel)

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   tabel,
	})
}
